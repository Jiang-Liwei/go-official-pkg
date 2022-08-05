package main

import "os"

// panic内置函数停止当前goroutine的正常执行。
// 当函数F调用panic时，F的正常执行立即停止。
// 任何被F延迟执行的函数都会以通常的方式运行，然后F返回给它的调用者。
// 对调用者G来说，F的调用行为就像是对紧急状态的调用，终止G的执行并运行任何延迟函数。
// 这将继续，直到正在执行的goroutine中的所有函数都按相反顺序停止。
// 此时，程序将以非零退出代码终止。这种终止序列称为恐慌，可以由内置函数recover控制。

func main() {
	chan1 := make(chan int)
	//当主进程结束 goroutine 并没有执行
	go tst(chan1)

	// 这样的调用是会在panic后执行的
	defer tst(chan1)

	panic("报错啦")
	// 这里是为了测试没有panic时 tst 的执行问题
	<-chan1
}

func tst(chan1 chan int) {
	// 不解释直接结束

	file, err := os.Create("/opt/Go/pkg/test.go")
	println(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	chan1 <- 1
}
