package main

import "fmt"

// complex 内建函数将两个浮点数值构造成一个复数值。
// 其实部和虚部的大小必须相同，即 float32 或 float64（或可赋予它们的），
// 其返回值 即为对应的复数类型（complex64 对应 float32，complex128 对应 float64）。
func main() {
	var floatOne float32
	var floatTwo float32
	floatOne = 222.22
	floatTwo = 666.22
	fmt.Println(floatOne)
	fmt.Println(floatTwo)

	result := complex(floatOne, floatTwo)

	fmt.Println(result)
}
