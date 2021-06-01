package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>我正在学Golang。</p>")
}

type UserInfo struct {
	Name  string
	Doing string
	Book  string
}

func main() {
	fmt.Printf("%b\n", 99)     // 101
	fmt.Printf("%c\n", 0x534e) // 中
	fmt.Printf("%d\n", 0x534e) // 18
	fmt.Printf("%o\n", 10)     // 12
	fmt.Printf("%q\n", "爱我中华") // '中'
	fmt.Printf("%q\n", 0x534e) // '中'
	fmt.Printf("%x\n", 123456) // d
	fmt.Printf("%X\n", 123456) // D
	fmt.Printf("%U", 0x534e)   // U+4E2D
	//var ip *int;
	// var c *int

	// var b *int

	// fmt.Println(c)
	// *c = *b
	// fmt.Println(*c)
	//fmt.Println(k) // 获取变量在计算机内存中的地址，可在变量名前面加上&字符。
	// &k 引用的是变量k的值，值所在的内存地址
	//showMemoryAddress(*k) //返回的地址是相同的

	//setup := testPointer()
	//fmt.Println("setup: ", setup)
	//fmt.Println("*setup: ", *setup)

	//http.HandleFunc("/", handlerFunc)
	//http.ListenAndServe(":8001", nil)
}

// func showMemoryAddress(x *int) { // *int参数类型位指向整数的指针，而不是整数
// 	fmt.Println(*x)
// 	//fmt.Println(*x) //本身就是指针，打印地址不需要  &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
// 	return
// }
