package main

// print内置函数以特定于实现的方式格式化其参数，并将结果写入standard error。
// 打印对于引导和调试非常有用；它不能保证使用这种语言。

// println内置函数以特定于实现的方式格式化其参数，并将结果写入标准错误。
// 始终在参数之间添加空格，并添加换行符。Println对于引导和调试非常有用；它不能保证使用这种语言。
func main() {
	print("I'm is man", "\n")
	map1 := map[int]int{5: 6}
	print(map1, "\n")

	println(map1[5])
}
