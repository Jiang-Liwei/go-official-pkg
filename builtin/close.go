package main

import "fmt"

// close内置函数关闭通道，该通道必须是双向的或仅发送。
// 它应该只由发送方执行，而不是由接收方执行，并具有在接收到最后一个发送值后关闭通道的效果。
// 在从闭合通道c接收到最后一个值后，从c接收的任何值都将成功，不会阻塞，并返回通道元素的零值。
// 格式：x, ok := <-c. 对于已关闭且空的通道，也会将 ok 设置为 false。
func main() {
	var chan1 chan string
	var chan2 chan string
	var chanClose chan string

	chan1 = make(chan string, 5)
	chan2 = make(chan string, 5)
	chanClose = make(chan string, 5)

	defer close(chan1)
	// 无法关闭通道，通道为空
	close(chanClose)

	// 触发死锁无法传值
	chan1 <- "chan1"
	chan2 <- "chan2 1"
	chan2 <- "chan2 2"
	chan2 <- "chan2 3"
	// 关闭后无法操作
	// chanClose <- "chanClose"

	close(chan2)

	chan1String, ok := <-chan1
	fmt.Println(chan1String, ok)

	oneChan2String := <-chan2
	fmt.Println(oneChan2String)

	// 依旧能获取到chan内的值，但是只能是第一个
	chan2String, ok := <-chan2
	fmt.Println(chan2String, ok)

	// 关闭后依然可以取出
	thereChan2String := <-chan2
	fmt.Println(thereChan2String)

	// 取空后为false
	fourChan2String, ok := <-chan2
	fmt.Println(fourChan2String, ok)

	// 这里为空直接false
	chanCloseString, ok := <-chanClose
	fmt.Println(chanCloseString, ok)
}
