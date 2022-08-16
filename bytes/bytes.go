package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// MinRead 是 Buffer.ReadFrom 传递给 Read 调用的最小切片大小。
	// 只要 Buffer 至少有 MinRead 字节超出了保存 r 内容所需的字节数，ReadFrom 就不会增长底层缓冲区。
	minread := bytes.MinRead
	fmt.Println(minread)

	//如果无法分配内存以将数据存储在缓冲区中，则将 ErrTooLarge 传递给恐慌。
	ErrTooLarge := bytes.ErrTooLarge
	println(ErrTooLarge)

	// Compare 返回一个整数，按字典顺序比较两个字节切片。
	// 如果 a == b，结果将为 0，如果 a < b，则为 -1，如果 a > b，则为 +1。一个 nil 参数相当于一个空切片。
	a := []byte{1, 2, 3, 4, 100}
	b := []byte{1, 2, 3, 4, 5}
	c := []byte{1, 2, 3, 4, 5, 10, 55}
	if bytes.Compare(a, b) < 0 {
		println("a < b")
	}
	if bytes.Compare(a, b) <= 0 {
		println("a <= b")
	}
	if bytes.Compare(a, b) > 0 {
		println("a > b")
	}
	if bytes.Compare(a, b) >= 0 {
		println("a >= b")
	}
	fmt.Println(bytes.Compare(a, c))

	// Prefer Equal to Compare for equality comparisons.
	if bytes.Equal(a, b) {
		println("a = b")
	}
	if !bytes.Equal(a, b) {
		println("a != b")
	}

	// 包含报告子切片是否在 b 内。
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))

	// ContainsAny 报告 chars 中是否有任何 UTF-8 编码的代码点在 b 内。
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	fmt.Println(bytes.ContainsAny([]byte(""), ""))

	// ContainsRune 报告符文是否包含在 UTF-8 编码的字节片 b 中。
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'f'))
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '!'))
	fmt.Println(bytes.ContainsRune([]byte(""), '@'))

	// func Count(s, sep []byte) int
	// Count 计算 s 中 sep 的非重叠实例的数量。
	// 如果 sep 是一个空切片，则 Count 返回 1 + s 中 UTF-8 编码的代码点数。
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune

	// func Cut(s, sep []byte) (before, after []byte, found bool)
	// 在 sep 的第一个实例周围切割切片 s，返回 sep 之前和之后的文本。
	// 找到的结果报告 sep 是否出现在 s 中。如果 sep 没有出现在 s 中，cut 返回 s, nil, false。
	// Cut 返回原始切片 s 的切片，而不是副本。

	show := func(s, sep string) {
		before, after, found := bytes.Cut([]byte(s), []byte(sep))
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")

	// func Equal(a, b []byte) bool
	// Equal 报告 a 和 b 是否具有相同的长度并包含相同的字节。一个 nil 参数相当于一个空切片。
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))

	// func EqualFold(s, t []byte) bool
	// EqualFold 报告被解释为 UTF-8 字符串的 s 和 t 在简单的 Unicode 大小写折叠下是否相等，这是一种更一般的不区分大小写形式。
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))

	// func Fields(s []byte) [][]byte
	// 字段将s解释为UTF-8编码的代码点序列。
	// 它围绕由unicode定义的一个或多个连续空格字符的每个实例分割切片。
	// IsSpace，返回s的子切片，如果s仅包含空白，则返回空切片。
	fmt.Printf("Fields are: %q\n", bytes.Fields([]byte("  foo bar  baz   ")))

	// func FieldsFunc(s []byte, f func(rune) bool) [][]byte
	// FieldsFunc将s解释为UTF-8编码的代码点序列。它在满足f（c）的每个代码点c处分割切片s，并返回s的子切片。
	// 如果s中的所有代码点满足f（c），或len（s）==0，则返回空切片。
	// FieldsFunc 不保证它调用 f(c) 的顺序，并假设 f 对于给定的 c 总是返回相同的值。
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q \n", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))

	// func HasPrefix(s, prefix []byte) bool
	// HasPrefix 测试byte 切片 s 是否以前缀开头。
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))

	// func HasSuffix(s, suffix []byte) bool
	// HasSuffix 测试byte 切片 s 是否以 suffix 结尾。
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))

	// func Index(s, sep []byte) int
	// Index 返回 s 中第一个 sep 实例的索引，如果 s 中不存在 sep，则返回 -1。
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))

	// func IndexAny(s []byte, chars string) int
	// IndexAny 将 s 解释为 UTF-8 编码的 Unicode 代码点序列。
	// 它返回 chars 中任何 Unicode 代码点在 s 中第一次出现的字节索引。
	// 如果 chars 为空或没有共同的代码点，则返回 -1。
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aieouy"))
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))

	// func IndexByte(b []byte, c byte) int
	// IndexByte 返回 b 中 c 的第一个实例的索引，如果 b 中不存在 c，则返回 -1。
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k')))
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('g')))

	// func IndexFunc(s []byte, f func(r rune) bool) int
	// IndexFunc 将 s 解释为 UTF-8 编码的代码点序列。
	// 它返回满足 f(c) 的第一个 Unicode 代码点的 s 中的字节索引，如果不满足，则返回 -1
	f2 := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f2))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f2))

	// func IndexRune(s []byte, r rune) int
	// IndexRune 将 s 解释为 UTF-8 编码的代码点序列。
	// 它返回给定符文在 s 中第一次出现的字节索引。
	// 如果 s 中不存在 rune，则返回 -1。
	// 如果 r 是 utf8.RuneError，它返回任何无效 UTF-8 字节序列的第一个实例。
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))

	// func Join(s [][]byte, sep []byte) []byte
	// Join 连接 s 的元素以创建一个新的字节切片。
	// 分隔符 sep 放置在结果切片中的元素之间。
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	sByte := bytes.Join(s, []byte(", "))
	fmt.Println(sByte)
	fmt.Printf("%s \n", sByte)

	// func LastIndex(s, sep []byte) int
	// LastIndex 返回 s 中 sep 的最后一个实例的索引，如果 s 中不存在 sep，则返回 -1。
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
	// func LastIndexAny(s []byte, chars string) int
	// LastIndexAny 将 s 解释为 UTF-8 编码的 Unicode 代码点序列。
	// 它返回 chars 中任何 Unicode 代码点在 s 中最后一次出现的字节索引。
	// 如果 chars 为空或没有共同的代码点，则返回 -1。
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "MüQp"))
	fmt.Println(bytes.LastIndexAny([]byte("go 地鼠"), "地大"))
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "z,!."))

	// func LastIndexByte(s []byte, c byte) int
	// LastIndexByte 返回 s 中 c 的最后一个实例的索引，如果 c 不存在于 s 中，则返回 -1。
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('g')))
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('r')))
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('z')))

	// func LastIndexFunc(s []byte, f func(r rune) bool) int
	// LastIndexFunc 将 s 解释为 UTF-8 编码的代码点序列。
	// 它返回满足 f(c) 的最后一个 Unicode 代码点的 s 中的字节索引，如果不满足，则返回 -1。
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber))

	// func Map(mapping func(r rune) rune, s []byte) []byte
	// Map 返回字节切片 s 的副本，其中所有字符都根据映射函数进行了修改。
	// 如果映射返回负值，则从字节片中删除该字符而不进行替换。
	// s 和输出中的字符被解释为 UTF-8 编码的代码点。

	// func Repeat(b []byte, count int) []byte
	// 重复返回一个由 b 的计数副本组成的新字节片。
	// 如果 count 为负数或 (len(b) * count) 的结果溢出，它会panics。
	fmt.Printf("ba%s\n", bytes.Repeat([]byte("na"), 2))
	// func Replace(s, old, new []byte, n int) []byte
	// Replace 返回切片 s 的副本，其中前 n 个不重叠的 old 实例替换为 new。
	// 如果 old 为空，则匹配切片的开头和每个 UTF-8 序列之后，为 k-rune 切片产生多达 k+1 个替换。
	// 如果 n < 0，则替换次数没有限制。
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))

	// func ReplaceAll(s, old, new []byte) []byte
	// ReplaceAll 返回切片 s 的副本，其中所有不重叠的 old 实例都被 new 替换。
	// 如果 old 为空，它会在切片的开头和每个 UTF-8 序列之后匹配，产生最多 k+1 个替换 k-rune 切片。
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))

	// func Runes(s []byte) []rune
	// Runes 将 s 解释为 UTF-8 编码的代码点序列。
	// 它返回一个与 s 等效的符文（Unicode 代码点）切片。
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}

	// func Split(s, sep []byte) [][]byte
	// 将切片 s 拆分为由 sep 分隔的所有子切片，并返回这些分隔符之间的子切片的切片。
	// 如果 sep 为空，Split 在每个 UTF-8 序列之后拆分。它相当于SplitN，计数为-1。

	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))

	// func SplitAfter(s, sep []byte) [][]byte
	// SplitAfter 在 sep 的每个实例之后将 s 切片为所有子切片，并返回这些子切片的一个切片。
	// 如果 sep 为空，SplitAfter 在每个 UTF-8 序列之后拆分。它等效于计数为 -1 的 SplitAfterN。
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
	// func SplitAfterN(s, sep []byte, n int) [][]byte
	// SplitAfterN 在 sep 的每个实例之后将 s 切片为子切片，并返回这些子切片的切片。
	// 如果 sep 为空，SplitAfterN 在每个 UTF-8 序列之后拆分。计数确定要返回的子切片数：
	// n > 0：最多 n 个子切片；最后一个子切片将是未拆分的剩余部分。 n == 0：结果为零（零个子切片） n < 0：所有子切片
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))

	// func SplitN(s, sep []byte, n int) [][]byte
	// SplitN 将 s 切片为由 sep 分隔的子切片，并返回这些分隔符之间的子切片的切片。
	// 如果 sep 为空，SplitN 在每个 UTF-8 序列之后拆分。计数确定要返回的子切片数：
	// n > 0：最多 n 个子切片；最后一个子切片将是未拆分的剩余部分。 n == 0：结果为零（零个子切片） n < 0：所有子切片
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	// func ToLower(s []byte) []byte
	// ToLower 返回字节切片 s 的副本，其中所有 Unicode 字母都映射到它们的小写字母。
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher")))
	// func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
	// ToLowerSpecial 将 s 视为 UTF-8 编码的字节并返回一个副本，其中所有 Unicode 字母都映射到它们的小写，优先考虑特殊的大小写规则。
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	totitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToLower : " + string(totitle))
	// func ToTitle(s []byte) []byte
	// ToTitle 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中所有 Unicode 字母都映射到它们的标题大小写。
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))
	// func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
	// ToTitleSpecial 将 s 视为 UTF-8 编码的字节并返回一个副本，
	// 其中所有 Unicode 字母都映射到它们的标题大小写，优先考虑特殊的大小写规则。
	str = []byte("ahoj vývojári golang")
	totitle = bytes.ToTitleSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToTitle : " + string(totitle))

	// func ToValidUTF8(s, replacement []byte) []byte
	// ToValidUTF8 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中代表无效 UTF-8 的每个字节运行被替换中的字节替换，这可能是空的。

	// func Trim(s []byte, cutset string) []byte
	// Trim 通过切掉 cutset 中包含的所有前导和尾随 UTF-8 编码的代码点来返回 s 的子切片。
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))

	// func TrimFunc(s []byte, f func(r rune) bool) []byte
	// TrimFunc 通过切掉所有满足 f(c) 的前导和尾随 UTF-8 编码的代码点 c 来返回 s 的子切片。
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))

	// func TrimLeft(s []byte, cutset string) []byte
	// TrimLeft 通过切掉 cutset 中包含的所有前导 UTF-8 编码代码点来返回 s 的子切片。
	fmt.Print(string(bytes.TrimLeft([]byte("453gopher8257"), "0123456789")))
	// func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
	// TrimLeftFunc 将 s 视为 UTF-8 编码字节，并通过切掉所有满足 f(c) 的前导 UTF-8 编码代码点 c 来返回 s 的子切片。
	var bTwo = []byte("Goodbye,, world!")
	bTwo = bytes.TrimPrefix(bTwo, []byte("Goodbye,"))
	bTwo = bytes.TrimPrefix(bTwo, []byte("See ya,"))
	fmt.Printf("Hello%s", bTwo)

	// func TrimRight(s []byte, cutset string) []byte
	// TrimRight 通过切掉包含在 cutset 中的所有尾随 UTF-8 编码的代码点来返回 s 的子切片。
	fmt.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789")))

	// func TrimRightFunc(s []byte, f func(r rune) bool) []byte
	// TrimRightFunc 通过切掉所有满足 f(c) 的尾随 UTF-8 编码代码点 c 来返回 s 的子切片。
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))

	// func TrimSpace(s []byte) []byte
	// TrimSpace returns a subslice of s by slicing off all leading and trailing white space, as defined by Unicode.
	fmt.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))

	// func TrimSuffix(s, suffix []byte) []byte
	// TrimSuffix 返回没有提供的尾随后缀字符串的 s。如果 s 不以 suffix 结尾，则 s 原样返回。
	var bThere = []byte("Hello, goodbye, etc!")
	bThere = bytes.TrimSuffix(bThere, []byte("goodbye, etc!"))
	bThere = bytes.TrimSuffix(bThere, []byte("gopher"))
	bThere = append(bThere, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	os.Stdout.Write(bThere)
}
