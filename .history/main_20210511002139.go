package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", []byte("Go语言")) // Go语言
	fmt.Printf("%q\n", "Golang")       // "Go语言"
	fmt.Printf("%x\n", "study Go")     // 676f6c616e67
	fmt.Printf("%X\n", "study Go")     // 676F6C616E67
}

// func showMemoryAddress(x *int) { // *int参数类型位指向整数的指针，而不是整数
// 	fmt.Println(*x)
// 	//fmt.Println(*x) //本身就是指针，打印地址不需要  &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
// 	return
// }
