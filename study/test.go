package main

import "fmt"

func main() {
	fmt.Println("Hello,World!")
	fmt.Println("1+1=", 1+1)
	fmt.Printf("1+1=%d", 1+1)
	// 使用8位无符号整型
	var X uint8 = 225
	fmt.Println(X+1, X)

	// 使用16位有符号整型
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2, Y)

	// a := 20.45
	// b := 34.89

	// //两个浮点数相减
	// c := b - a

	// //显示结果
	// fmt.Printf("结果: %f", c)

	// //显示c变量的类型
	// fmt.Printf("\nc的类型是 : %T", c)

	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)
	fmt.Println(a)
	fmt.Println(b)

	//显示类型
	fmt.Printf("a的类型是 %T 以及"+"b的类型是 %T", a, b)

	//变量声明和初始化不使用表达式
	var myvariable1 int
	var myvariable2 string
	var myvariable3 float64

	//显示0值变量
	fmt.Printf("myvariable1的值是 : %d\n", myvariable1)
	fmt.Printf("myvariable2的值是 : %d\n", myvariable2)
	fmt.Printf("myvariable3的值是 : %d\n", myvariable3)
}
