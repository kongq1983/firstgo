package httplib

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
	"strconv"
)

type User struct {

	Id      int64    `json:"id"`
	Status      int64    `json:"status"`
	Hobby      string   `json:"hobby"`
	Username      string   `json:"username"`
	Name      string   `json:"name"`

}


type Response struct {
	Code      int32   `json:"code"`
	Msg      string   `json:"msg"`
	Data       any   `json:"data"`   // 注意any  服务器返回这个值,不确定具体的类型
	SysCode      string    `json:"sysCode"`
}


func GetBaidu(){

	str, err := httplib.Get("http://www.baidu.com/").String()
	if err != nil {
		// error
	}
	fmt.Println(str)

}

func UserAddByJson()  {

	var user User
	user.Id = 1
	user.Name = "管理员"
	user.Username="admin"
	user.Hobby = "football"
	user.Status = 1

	req := httplib.Post("http://localhost:10000/user/add")
	req.JSONBody(user) // 会设置application/json

	str,err := req.String()
	if err != nil {
		// error
		log.Println(err.Error())
	}

	fmt.Println(str)

}


func UserAddByParam()  {

	var user User
	user.Id = 1
	user.Name = "管理员"
	user.Username="admin"
	user.Hobby = "football"
	user.Status = 1

	req := httplib.Post("http://localhost:10000/user/add1")
	req.Param("id",strconv.FormatInt(int64(user.Id), 10))
	req.Param("name",user.Name)
	req.Param("username",user.Username)
	req.Param("hobby",user.Hobby)
	req.Param("status",strconv.FormatInt(int64(user.Status), 10))

	req.Debug(true)

	var res Response
	req.ToJSON(&res)  // 注意这个 服务器返回值 赋值到res


	str,err := req.String()
	if err != nil {
		// error
		log.Println(err.Error())
	}

	fmt.Println(str)
	fmt.Printf("res.code= %d  \n",res.Code)
	fmt.Printf("res.Data= %v  \n",res.Data) // 服务器如果返回int 这里用string接收,则会接收不到,所有用any
	fmt.Printf("res.Msg= %s  \n",res.Msg)
	fmt.Printf("res.SysCode= %s  \n",res.SysCode)

}
