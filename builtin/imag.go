package main

import "fmt"

// imag 内建函数返回复数 c 的虚部。
// 其返回值为对应于 c 类型的浮点数。
func main() {
	var floatOne float32
	var floatTwo float32
	floatOne = 222.22
	floatTwo = 666.22
	fmt.Println(floatOne)
	fmt.Println(floatTwo)

	complex1 := complex(floatOne, floatTwo)

	fmt.Println(complex1)

	result := imag(complex1)
	fmt.Println(result)
}
