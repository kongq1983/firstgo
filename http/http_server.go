package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	//now := time.Now().Format("20060102")

	common := Common{}

	clientIp := common.GetClientIp(r)

	str := "%s, welcome to you ! \n  now=%s \n  yesterday=%s \n  clientIp=%s \n"

	res := fmt.Sprintf(str,name,common.GetToNow(),common.GetYesterDay(),clientIp)

	fmt.Fprint(w,res)
}

func printParam(writer http.ResponseWriter, request *http.Request) {

	fmt.Printf("printParam is called! \n")

	// 检查是否POST请求
	if request.Method != "POST" {
		writer.WriteHeader(405)
		return
	}
	// 解析form
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(400)
		return
	}
	// 设置Content-Type
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	data := map[string]interface{}{
		"method":   request.Method,
		"url":      request.RequestURI,
		"header":   request.Header,
		// 包含URL里的查询参数
		"query":    request.URL.Query(),
		// 只包含body里的参数
		"postForm": request.PostForm,  // 如果是application/json 则是nil
		// 包含URL和body里的参数
		"form":     request.Form,
	}

	json.NewEncoder(writer).Encode(data)

}


func printJsonParam(writer http.ResponseWriter, request *http.Request) {

	fmt.Printf("printJsonParam is called! \n")

	// 检查是否POST请求
	if request.Method != "POST" {
		writer.WriteHeader(405)
		return
	}

	bytes,err := io.ReadAll(request.Body)

	if err != nil {
		log.Println(err.Error())
		return
	}

	// 设置Content-Type
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	fmt.Printf(" body =%s \n",string(bytes))

	writer.Write(bytes)

	//json.NewEncoder(writer).Encode(bytes)

}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",hello)
	mux.HandleFunc("/printParam",printParam)
	mux.HandleFunc("/printJsonParam",printJsonParam)

	server := &http.Server{
		Addr: ":10000",
		Handler:mux,
	}

	if err := server.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}