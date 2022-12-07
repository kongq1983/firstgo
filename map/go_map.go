package _map

import "fmt"

func Demo()  {

	demo0()
	demo1()

}

func demo0()  {
	mymap := make(map[string]int)  // 创建一个从string到int的

	mymap["one"]=1
	mymap["two"]=2

	for name, age := range mymap {
		fmt.Println(name, age)
	}

}

func demo1()  {

	mymap := map[string]int{
		"admin": 18,
		"guest": 28,
		"user": 20,
	}

	for name, age := range mymap {
		fmt.Println(name, age)
	}

}