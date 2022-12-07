package json

import (
	"encoding/json"
	"fmt"
	"kq.com/firstgo/http"
	"time"
)

// `json:"name"`
//我们再每个属性后面都加上了tag,tag其实是没有什么意义的,它的整个go语言的运行中没有什么意义,
//只是给json.Marshal看的;如果有tag再生成json的时候会按照tag里面的内容替换成大写的属性名

type MyFileInfo struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Md5       string   `json:"md5"`
	Size      int64    `json:"size"`
	Peers     []string `json:"peers"`
	TimeStamp int64    `json:"timeStamp"`
}





func jsonObjectDemo()  {

	common := http.Common{}

	var info MyFileInfo
	info.Name = "java"
	info.Md5 = common.MD5(info.Name)
	info.Path = "/tmp/book/java"
	info.Size = 1024 * 1
	info.TimeStamp = time.Now().Unix()

	info.Peers = []string{"http://localhost:10001", "http://localhost:10002"}

	jsons, errs := json.Marshal(info)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}

	fmt.Println("MyFileInfo json data :", string(jsons))

	//反序列化
	var info2 MyFileInfo

	//将Json转换成struct用json.Unmarshal就可以了,go语言对解析Json有同样良好的支持,
	//再解析的时候要注意的是一定要把struct的地址传过去,
	//如果将值传过去,那么就是数据的复制,不会影响到自己定义的struct;
	//只有将地址传过去才能完成 正常的解析
	errs = json.Unmarshal(jsons, &info2) // 第2个参数 要传地址
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("MyFileInfo info2 :", info2)


}


func jsonArrayDemo()  {

	fmt.Println("------------------------------------------")

	//var array = [...]MyFileInfo{}
	var array []MyFileInfo
	//var array = MyFileInfo{}
	common := http.Common{}

	for i:=0;i<3;i++ {

		var info MyFileInfo
		info.Name = "java"
		info.Md5 = common.MD5(info.Name)
		info.Path = "/tmp/book/java"
		info.Size = int64(i)
		info.TimeStamp = time.Now().Unix()

		info.Peers = []string{"http://localhost:10001", "http://localhost:10002"}

		array = append(array,info)

	}


	jsons, errs := json.Marshal(array)
	if errs != nil {
		fmt.Println("json array marshal error:", errs)
	}

	fmt.Println("MyFileInfo array json data :", string(jsons))

	fmt.Println("------------------------------------------")
}


func Start(){
	jsonObjectDemo()
	jsonArrayDemo()
}