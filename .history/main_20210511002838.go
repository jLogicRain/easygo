package main

import (
	"fmt"
)

type UserInfo struct {
	Name  string
	Doing string
	Book  string
}

func main() {
	user := UserInfo{Name: "logic", Doing: "学习", Book: "Golang"}
	fmt.Printf("%p", &user)
}

// func showMemoryAddress(x *int) { // *int参数类型位指向整数的指针，而不是整数
// 	fmt.Println(*x)
// 	//fmt.Println(*x) //本身就是指针，打印地址不需要  &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
// 	return
// }
