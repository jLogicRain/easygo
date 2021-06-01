package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%b", 10.20) // 5742089524897382p-49
	fmt.Printf("%e", 10.20) // 1.020000e+01
	fmt.Printf("%E", 10.20) // 1.020000E+01
	fmt.Printf("%f", 10.20) // 10.200000
	fmt.Printf("%F", 10.20) // 10.200000
	fmt.Printf("%g", 10.20) // 10.2
	fmt.Printf("%G", 10.20) // 10.2
}

// func showMemoryAddress(x *int) { // *int参数类型位指向整数的指针，而不是整数
// 	fmt.Println(*x)
// 	//fmt.Println(*x) //本身就是指针，打印地址不需要  &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
// 	return
// }
