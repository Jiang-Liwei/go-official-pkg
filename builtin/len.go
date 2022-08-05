package main

import "fmt"

// len 内置函数根据类型返回 v 的长度
// 数组：v中的元素数。
// 指向数组的指针：*v中的元素数（即使v为零）。
// 切片或map：v中的元素数；如果v为零，则len（v）为零。
// 字符串：v中的字节数。
// 通道：通道缓冲区中排队（未读）的元素数；
// 如果v为零，则len（v）为零。
// 对于某些参数，例如字符串文字或简单的数组表达式，结果可以是常量。
// 有关详细信息，请参阅 Go 语言规范的“长度和容量”部分。
func main() {
	var array1 = [9]int{1}
	var array1Pointer *[9]int
	array1Pointer = &array1
	fmt.Println("array1", len(array1))
	fmt.Println("array1Pointer", len(array1Pointer))

	slice1 := array1[2:]
	fmt.Println("slice1", len(slice1))
	map1 := make(map[string]string)
	fmt.Println("map1", len(map1))
	map1["a"] = "A"
	fmt.Println("map1", len(map1))

	// 中文三个长度
	var string1 = "付aa江"
	fmt.Println("string1", len(string1))

	chan1 := make(chan int, 5)
	fmt.Println("chan1", len(chan1))
	chan1 <- 99
	fmt.Println("chan1", len(chan1))

}
