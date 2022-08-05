package main

import "fmt"

type NewOne struct {
	int2    int
	string2 string
	bool2   bool
	array2  [10]int
	slice2  []int
	map2    map[int]int
	chan2   chan int
}

// 新的内置函数分配内存。
// 第一个参数是一个类型，而不是一个值，返回的值是一个指向该类型新分配的零值的指针。
func main() {
	slice1 := new([]int)
	fmt.Println("slice1", slice1)

	map1 := new(map[int]int)
	fmt.Println("map1", map1)

	int1 := new(int)
	fmt.Println("int1", int1)

	string1 := new(string)
	fmt.Println("string1", string1)
	fmt.Println("string1", &string1)
	bool1 := new(bool)
	fmt.Println("bool1", bool1)

	array1 := new([10]int)
	array1[0] = 55
	fmt.Println("array1", array1)

	chan1 := new(chan int)
	fmt.Println("chan1", chan1)

	newOne := new(NewOne)
	fmt.Println(newOne)
	newOne.array2[3] = 9
	newOne.slice2 = []int{2, 4, 5}
	newOne.map2 = map[int]int{1: 2}
	newOne.chan2 = make(chan int, 2)
	fmt.Println(newOne)
}

// 总结new只有数组和结构体能使用
