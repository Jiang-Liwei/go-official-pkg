package main

import "fmt"

func main() {
	var floatOne float64
	var floatTwo float64
	floatOne = 222.22
	floatTwo = 666.22
	result := complex(floatOne, floatTwo)
	var ComplexType complex64 = complex64(result)
	// ComplexType 仅用于文档目的。它是任一复杂类型的替代：complex64 或 complex128。
	fmt.Println("ComplexType", ComplexType)
	// FloatType 仅用于文档目的。它是浮点类型的替代：float32 或 float64。
	fmt.Println("FloatType", floatOne)
	// IntegerType 仅用于文档目的。它是任何整数类型的替代：int、uint、int8 等。
	var intOne uint = 999
	fmt.Println("IntegerType", intOne)
	// Type 此处的类型仅用于文档目的。它是任何 Go 类型的替代，但代表任何给定函数调用的相同类型。
	type Type map[string]string
	// Type1 仅用于文档目的。它是任何 Go 类型的替代，但代表任何给定函数调用的相同类型。
	type Type1 int
	// any 是 interface{} 的别名，在所有方面都等效于 interface{}。
	type any = interface{}
	// bool 是一组布尔值，true 和 false。
	type bool1 bool
	// byte 是 uint8 的别名，在所有方面都等同于 uint8。按照惯例，它用于区分字节值和 8 位无符号整数值。
	type byte1 byte
	// comparable是由所有可比较类型（布尔、数字、字符串、指针、通道、可比较类型数组、字段均为可比较类型的结构）实现的接口。
	// 可比较接口只能用作类型参数约束，而不能用作变量类型。
	type comparable1 interface{ comparable }
	// complex128 是具有 float64 实部和虚部的所有复数的集合。
	type complex128One complex128
	// complex64 是具有 float32 实部和虚部的所有复数的集合。
	type complex64One complex64
	// error 内置接口类型是表示错误情况的常规接口，nil 值表示没有错误。
	type error1 interface {
		Error() string
	}
	// float32 是所有 IEEE-754 32 位浮点数的集合。
	type float32One float32
	// float64 是所有 IEEE-754 64 位浮点数的集合。
	type float64One float64
	// int 是一个有符号整数类型，大小至少为 32 位。然而，它是一种独特的类型，而不是 int32 的别名
	type int1 int
	// int16 是所有有符号 16 位整数的集合。范围：-32768 到 32767。
	type int16One int16
	// int32 是所有带符号的 32 位整数的集合。范围：-2147483648 到 2147483647。
	type int32One int32
	// int64 是所有有符号 64 位整数的集合。范围：-9223372036854775808 到 9223372036854775807。
	type int64One int64
	// int8 是所有有符号 8 位整数的集合。范围：-128 到 127。
	type int8One int8
	// rune 是 int32 的别名，在所有方面都等效于 int32。按照惯例，它用于区分字符值和整数值。
	type runeOne = int32
	// string 是所有 8 位字节字符串的集合，通常但不一定代表 UTF-8 编码的文本。
	// 字符串可以为空，但不能为零。字符串类型的值是不可变的。
	type string1 string
	// uint 是至少 32 位大小的无符号整数类型。然而，它是一种独特的类型，而不是 uint32 的别名。
	type uint1 uint
	// uint16 是所有无符号 16 位整数的集合。范围：0 到 65535。
	type uint16One uint16
	// uint32 是所有无符号 32 位整数的集合。范围：0 到 4294967295。
	type uint32One uint32
	// uint64 是所有无符号 64 位整数的集合。范围：0 到 18446744073709551615。
	type uint64One uint64
	// uint8 是所有无符号 8 位整数的集合。范围：0 到 255。
	type uint8One uint8
	// uintptr 是一个整数类型，它大到足以容纳任何指针的位模式。
	type uintptr1 uintptr
}
