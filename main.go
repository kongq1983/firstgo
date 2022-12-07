package main

import (
	"fmt"
	"kq.com/firstgo/emp"
	"kq.com/firstgo/http"
	"kq.com/firstgo/httplib"
	"kq.com/firstgo/json"
	_map "kq.com/firstgo/map"
	"kq.com/firstgo/print"
	"kq.com/firstgo/random"
	"kq.com/firstgo/redis"
	"kq.com/firstgo/reflect"
	mystring "kq.com/firstgo/string"
	"math/rand"
	"time"
)

func main1() {

	var e1 = emp.Employee{Id:1,Username: "admin",Name: "管理员",Age: 20}
	var e2 = emp.Employee{Id:2,Username: "guest",Name: "匿名用户",Age: 20}

	fmt.Println("e1",e1.ToString())
	fmt.Println("e2",e2.ToString())

	var equal = e1.IsEqualAge(&e2)
	fmt.Println("e1.age==e2.age",equal)

	var equal1 = e1.IsEqualAge1(e2)
	fmt.Println("e1.age==e2.age",equal1)


}

/**
	秒杀的例子
 */
func main2()  {

	//getSetValueDemo()

	//incrByDemo()

	var key = "product:1"
	redis.Set(key,"20")


	for i:=0;i<30;i++{
		go decrByDemo(i)
	}

	//go decrByDemo(1)


	time.Sleep(time.Second*10)
}

/**
 	val to upper
 */
func main5()  {

	//setToUpper()

	//redis.IterAllKeys()
	//print("\n=============================\n ")
	//redis.IterSetKeys()
	//print("\n=============================\n ")
	//redis.IterHashKeys()
	//print("\n=============================\n ")
	redis.IterZSetKeys();
}

func mainff()  {
	redis.HSetModel()
	redis.HGetToModel()
}

/**
	锁例子
 */
func main6()  {

	for i:=0;i<30;i++{
		go distributeLock(i)
	}

	time.Sleep(time.Second*10)

}

func mainfa()  {

	name,_ := random.Hello("admin")

	fmt.Printf("name=%s \n",name)

	index := rand.Intn(10)

	fmt.Printf("index=%d \n",index)

}

func mainpp()  {
	print.Demo()

	reflect.Demo()
}

func mainee()  {

	_map.Demo()

	fmt.Println(mystring.ToUpper("abcdefg"))

}


func getSetValueDemo()  {
	//myredis
	var key = "name"
	var key1 = "name1"
	var val = "admin"
	redis.Set(key,val)

	redis.SetNX(key1,val,60 * time.Second)

	var loadValue = redis.Get(key);
	fmt.Printf("key=%s loadValue=%s \n",key,loadValue)

	var loadValue1 = redis.Get(key1);
	fmt.Printf("key1=%s loadValue1=%s \n",key1,loadValue1)

	var loadValue2 = redis.Get(key);
	fmt.Printf("key=%s loadValue2=%s \n",key,loadValue2)
}

func incrByDemo()  {

	var key = "number1"



	var num, err = redis.IncrByLua(key)

	if err != nil {
		fmt.Printf(" IncrByLua key=%s fail ! err=%s \n", key, err)
	}

	fmt.Printf("key=%s val=%d \n",key,num)

}

/**
	商品销售
 */
func decrByDemo(threadIndex int) {

	var key = "product:1"

	var num, err = redis.DecrProductByLua(key)

	if err != nil {
		fmt.Printf(" DecrProductByLua key=%s fail ! err=%s \n", key, err)
	}

	fmt.Printf("threadIndex=%d key=%s val=%d \n",threadIndex,key,num)

}


/**
  分布式锁
 */
func distributeLock(threadIndex int) {

	var key = "distribute:lock:1"

	var num, err = redis.DistributeLock(key,threadIndex)

	if err != nil {
		fmt.Printf(" DistributeLock key=%s fail ! err=%s \n", key, err)
	}

	if num > 0 {
		fmt.Printf("success!  threadIndex=%d key=%s val=%d \n",threadIndex,key,num)
	} else {
		fmt.Printf("failure!  threadIndex=%d key=%s val=%d \n",threadIndex,key,num)
	}


}

func setToUpper()  {

	var key = "name2"
	var val = "admin"

	var str, err = redis.SetToUpper(key,val)

	if err != nil {
		fmt.Printf(" setToUpper key=%s fail ! err=%s \n", key, err)
	}

	fmt.Printf("key=%s val=%s \n",key,str)

}


func Sum(a *[5]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func multiParam(args ...string) {
	//接受的参数放在args数组中
	for _, e := range args {
		fmt.Println(e)
	}
}


func multiParam1(name string, args ...int) {
	fmt.Println(name)
	//接受的参数放在args数组中
	for _, e := range args {
		fmt.Println(e)
	}
}

func maine()  {

	// 数组中，...符号指定的长度等于数组中元素的数量
	array := [...]float64{7.0, 8.5, 9.1,5.4,20}
	x := Sum(&array)  // Note the explicit address-of operator

	fmt.Println(x)

	names := []string{"jerry", "herry"}
	multiParam(names...)
	//multiParam(names...)

	multiParam1("go", 1,2,3,4,5)

	nums := [] int{1,2,3}
	multiParam1("golang", nums...)

	arr := [...]int{1, 2, 3,4}
	//multiParam1("golang1", arr...)

	fmt.Println(arr)

}

func mainHttp()  {

	http.Start()

}

//%v,原样输出
//%T，打印类型
//%t,bool类型
//%s，字符串
//%f，浮点
//%d，10进制的整数
//%b，2进制的整数
//%o，8进制
//%x，%X，16进制
//%x：0-9，a-f
//%X：0-9，A-F
//%c，打印字符
//%p，打印地址

// json
func mainjson()  {
	json.Start()
}


func maineaaa()  {
	// test http_util.go

	common := http.Common{}

	ips := []string{"192.168.1.1","192.168.2.1"}

	contain :=common.Contains("192.168.1.1",ips)

	fmt.Printf("contain=%t",contain)

	aPath := "D:\\log\\go\\a.txt"
	bPath := "D:\\log\\go\\b.txt"

	exists := common.FileExists(aPath)

	fmt.Printf("D:\\log\\go\\a.txt  exists=%t \n",exists)

	common.WriteFile(bPath,"this is a.txt content")
	common.WriteAppendFile(aPath,"this is a.txt append content\n")

	//List<String> command = new ArrayList<String>();
	//command.add( "cmd.exe" );
	//command.add( "/c" );
	//command.add( "ipconfig -all" );

	//"ping", "www.baidu.com"

	//cmds := []string{"cmd.exe","/c","java","-","version"}
	//cmds := []string{"ping","www.baidu.com"}
	cmds := []string{"java","-version"}
	//cmds := []string{"java -version"} // 这个会报错
	//cmds := []string{"ipconfig","-all"}

	str, status := common.Exec(cmds,1000)

	fmt.Printf("str=%s ,status=%d \n",str,status)


}


// httplib
func main(){

	//httplib.GetBaidu()
	//httplib.UserAddByJson()
	httplib.UserAddByParam()

}