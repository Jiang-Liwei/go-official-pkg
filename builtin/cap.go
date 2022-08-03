package main

import (
	"fmt"
	"strconv"
	"time"
)

// 数组：返回数组中v的个数（与len（v）相同）。
// 数组指针：返回数组中*v的个数（与len（v）相同）。
// 切片：在重新切片时，切片能够达到的最大长度；若 v 为 nil，len(v) 即为零。
// 信道：按照元素的单元，相应信道缓存的容量；若 v 为 nil，len(v) 即为零。
func main() {
	var array [9]string
	i := 1
	for k, _ := range array {
		array[k] = strconv.Itoa(i)
		i++
	}
	var aa *[9]string
	aa = &array //
	fmt.Println(cap(array))
	fmt.Println(cap(aa))

	slice := aa[2:8]
	fmt.Println(slice)
	slice[5] = "999"
	fmt.Println(array)
	// 因为是从数组的第二个切片开始 使用可重新切片的长度只有7
	fmt.Println(cap(slice))
	slice[5] = "111"
	fmt.Println(array)
	var chan1 chan string
	chan1 = make(chan string, 5)

	go test(chan1)
	time.Sleep(time.Second)
	fmt.Println(cap(chan1))
	fmt.Println(<-chan1)
}

func test(chan1 chan string) {
	chan1 <- "99"
}
