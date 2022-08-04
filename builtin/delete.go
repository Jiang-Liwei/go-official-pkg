package main

import "fmt"

// delete 内置函数从映射中删除具有指定键 (m[key]) 的元素。
// 如果 m 为 nil 或没有这样的元素，则 delete 是空操作。
func main() {
	var map1 = make(map[int]int)
	for i := 1; i < 7; i++ {
		map1[i] = i
	}
	fmt.Println(map1)
	delete(map1, 3)
	fmt.Println(map1)
}
