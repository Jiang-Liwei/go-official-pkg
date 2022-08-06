package main

import "fmt"

// recover 内建函数允许程序管理恐慌过程中的Go程。
// 在已推迟函数（而不是任何被它调用的函数）中，执行 recover 调用会通过恢复正常的执行 并取回传至 panic 调用的错误值来停止该恐慌过程序列。
// 若 recover 在已推迟函数之外被调用， 它将不会停止恐慌过程序列。在此情况下，或当该Go程不在恐慌过程中时，
// 或提供给 panic 的实参为 nil 时，recover 就会返回 nil。因此 recover 的返回值就报告了该Go程是否 在恐慌过程中。

func badCall() {
	panic("bad end")
}

func badCall2() {
	panic("ssss")
}

func main() {
	var a int = 99
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(a)
			fmt.Println(r)
		}
		if r == nil {
			fmt.Println("panic了一个空指针")
		}
	}()
	defer func(a *int) {
		*a = 88
	}(&a)
	badCall()
	badCall2()
	//panic(nil)
	fmt.Println("正常执行")
	panic("这是第二个错误")
}
