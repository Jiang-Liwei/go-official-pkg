package main

import "fmt"

// copy 内建函数将元素从来源切片复制到目标切片中。
//（特殊情况是，它也能将字节从字符串复制到字节切片中）。
// 来源和目标可以重叠。 copy 返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个。
func main() {
	var array1 = [9]int{1, 2, 3}

	fmt.Println(array1)
	var slice1 = array1[3:6]

	fmt.Println(slice1)
	i := 4
	for k, _ := range slice1 {
		slice1[k] = i
		i++
	}

	fmt.Println(slice1)
	fmt.Println(array1)

	var slice2 = array1[6:9]
	fmt.Println(slice2)
	j := 7
	for k, _ := range slice2 {
		slice2[k] = j
		j++
	}
	fmt.Println(slice2)
	fmt.Println(array1)

	// 从切片第一个元素开始覆盖，直到目标或来源最后一个元素耗尽
	var slice3 = []int{66, 66, 66, 66}
	copy(slice1, slice3)
	fmt.Println(array1)
	copy(slice3, slice2)
	fmt.Println(slice3)
}
