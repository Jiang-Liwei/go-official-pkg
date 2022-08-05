package main

import "fmt"

// make内置函数分配并初始化类型为slice、map或chan的对象（仅限）。
// 与new一样，第一个参数是类型，而不是值。
// 与new不同，make的返回类型与其参数的类型相同，而不是指向它的指针。
// 结果的规格取决于类型：
// 切片：大小指定长度。切片的容量为等于其长度。第二整数参数可以被提供给指定不同的容量；
// 它必须不小于长例如，make（[]int，0，10）分配一个底层数组并返回长度为0、容量为10的切片，即由该底层阵列支持。
// map：为空map分配了足够的空间来容纳指定的元素数。可以省略尺寸分配小的起始大小。
// 通道：通道的缓冲区用指定的缓冲容量。如果为零或省略了大小，则通道为无缓冲。
func main() {
	slice1 := make([]int, 5, 10)
	fmt.Println("slice1", slice1)

	map1 := make(map[int]int, 5)
	fmt.Println("map1", map1)
	map1[1] = 3
	map1[2] = 3
	map1[3] = 3
	map1[4] = 3
	map1[5] = 3
	map1[6] = 3
	map1[7] = 3
	fmt.Println("map1", map1)

	// 无缓存通道
	chan1 := make(chan int)
	// 缓冲为10
	chan2 := make(chan string, 10)
	fmt.Println("chan1", len(chan1))
	fmt.Println("chan2", len(chan2))

}
