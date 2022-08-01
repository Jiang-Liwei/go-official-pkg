package main

import "fmt"

// builtin 包为Go的预声明标识符提供了文档.
// 此处列出的条目其实并不在 buildin 包中，对它们的描述只是为了让 godoc 给该语言的特殊标识符提供文档。
// 既此处为内置函数
func main() {
	// append 内建函数将元素追加到切片的末尾。
	// 若它有足够的容量，其目标就会重新切片以容纳新的元素。
	// 否则，就会分配一个新的基本数组。
	// append 返回更新后的切片。
	// 因此必须存储追加后的结果，通常为包含该切片自身的变量：
	var array = []int{1, 2, 3, 4, 5, 6}
	var slice []int = array[2:]
	// 这里不知道是不是新版的原因，查相关帖子的时候切片好像只能slice[0:1]这样取值
	// 就是 Go1.17 会正式支持切片（Slice）转换到数据（Array）（不知道是不是这个原因）
	a := slice[2]
	fmt.Println(a)
	slice[2] = 99
	fmt.Println(slice)
	fmt.Println(array)

	// 这里追加了一个7由于切片发生了扩容分配了一个新的数组给切片所以这一次数组的值无变化
	fmt.Printf("%p \n", &slice[0])
	slice = append(slice, 7)
	slice[1] = 99
	fmt.Printf("%p", &slice[0])
	fmt.Println(slice)
	slice = append(slice, 8)
	// 这里可以发现第一次append时slice[0]的地址发生了改变，第二次却没有
	// 是因为扩容规则为不足1024 直接容量翻倍，否则 则扩容1/4
	fmt.Printf("%p", &slice[0])
	fmt.Println(slice)
	fmt.Println(array)
}
