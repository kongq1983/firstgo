package print

import "fmt"

func Demo()  {
	sprintf()
	printf()
}

func sprintf() {

	name := "admin"
	age :=20

	message := fmt.Sprintf("sprintf1 name=%s age=%d",name,age)
	fmt.Println(message)
	message1 := fmt.Sprintf("sprintf2 name=%s age=%v",name,age)
	fmt.Println(message1)
	message2 := fmt.Sprintf("sprintf3 name=%v age=%v",name,age)
	fmt.Println(message2)

}

func printf() {

	name := "admin"
	age :=28

	fmt.Printf("printf1 name=%s age=%d \n",name,age)
	fmt.Printf("printf1 name=%v age=%v \n",name,age)

	fmt.Printf("printf1 name.type=%T age.type=%T \n",name,age)

	//x/X ：十六进制编码（小写/大写，以字节为元素进行编码，而不是字符）
	//如果使用了 " " 旗标，则在每个元素之间添加空格。
	//如果使用了 "#" 旗标，则在十六进制格式之前添加 0x 前缀。
	// 16禁制输出
	fmt.Printf("%x\n", age) //
	fmt.Printf("%X\n", age) //
	fmt.Printf("%#X\n", age) //


	//要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &age) //  带0x前缀的十六进制地址值
	fmt.Printf("%#p\n", &age) // 不带0x前缀的十六进制地址值

}