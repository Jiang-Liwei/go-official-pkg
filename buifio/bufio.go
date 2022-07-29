package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := []byte{'A', 65, 'b'}
	fmt.Println(a)

	// ScanBytes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个字节作为一个token返回。
	advance, token, _ := bufio.ScanBytes(a, true)
	fmt.Println(advance, token)

	// ScanLines是用于Scanner类型的分割函数（符合SplitFunc）， 本函数会将每一行文本去掉末尾的换行标记作为一个token返回。
	// 返回的行可以是空字符串。换行标记为一个可选的回车后跟一个必选的换行符。
	// 最后一行即使没有换行符也会作为一个token返回。
	advanceLines, tokenLines, _ := bufio.ScanLines(a, true)
	fmt.Println(advanceLines, tokenLines)

	// ScanRunes是用于Scanner类型的分割函数（符合SplitFunc）， 本函数会将每个utf-8编码的unicode码值作为一个token返回。
	// 本函数返回的rune序列和range一个字符串的输出rune序列相同。 错误的utf-8编码会翻译为U+FFFD = "\xef\xbf\xbd"， 但只会消耗一个字节。
	// 调用者无法区分正确编码的rune和错误编码的rune。
	advanceRunes, tokenRunes, _ := bufio.ScanRunes(a, true)
	fmt.Println(advanceRunes, tokenRunes)

	// ScanWords是用于Scanner类型的分割函数（符合SplitFunc）， 本函数会将每一行文本去掉末尾的换行标记作为一个token返回。
	// 返回的行可以是空字符串。换行标记为一个可选的回车后跟一个必选的换行符。 最后一行即使没有换行符也会作为一个token返回。
	advanceWords, tokenWords, _ := bufio.ScanWords(a, true)
	fmt.Println(advanceWords, tokenWords)
	file, _ := os.OpenFile("./bufio.go", os.O_RDWR, 0)

	// NewReadWriter分配新的ReadWriter来进行r和w的调度。
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)
	readWriter := bufio.NewReadWriter(reader, writer)
	fmt.Println(readWriter)

	// NewReaderSize返回了一个新的读取器，这个读取器的缓存大小至少大于制定的大小。
	// 如果io.Reader参数已经是一个有足够大缓存的读取器，它就会返回这个Reader了。
	b := bufio.NewReaderSize(file, 55)
	fmt.Println(b)

	// Buffered返回当前缓存的可读字节数。
	byteNum := b.Buffered()
	fmt.Println(byteNum)

	// Discard跳过接下来的n个字节，返回丢弃的字节数。
	// 如果Discard跳过的字节数少于n个，它也会返回一个错误。
	// 如果0<=n<=b.Buffered（），则可以保证在不读取底层io.Reader的情况下成功放弃。
	discarded, err := reader.Discard(50000)
	fmt.Println(discarded, err)

	// Peek返回没有读取的下n个字节。在下个读取的调用前，字节是不可见的。
	// 如果Peek返回的字节数少于n， 它一定会解释为什么读取的字节数段了。
	// 如果n比b的缓冲大小更大，返回的错误是ErrBufferFull。
	bytePeek, err := b.Peek(1)
	fmt.Println(bytePeek, err)

	// Read读取数据到p。 返回读取到p的字节数。 底层读取最多只会调用一次Read，因此n会小于len(p)。
	// 在EOF之后，调用这个函数返回的会是0和io.Eof。
	p := []byte{'A'}
	n, err := b.Read(p)
	fmt.Println(n, err)

	// ReadByte读取和回复一个单字节。
	// 如果没有字节可以读取，返回一个error。
	c, err := b.ReadByte()
	fmt.Println(c, err)

	// ReadBytes读取输入到第一次终止符发生的时候，返回的slice包含从当前到终止符的内容（包括终止符）。
	// 如果ReadBytes在遇到终止符之前就捕获到一个错误，它就会返回遇到错误之前已经读取的数据，和这个捕获 到的错误（经常是 io.EOF）。
	// 当返回的数据没有以终止符结束的时候，ReadBytes返回err != nil。 对于简单的使用，或许 Scanner 更方便。
	delim := byte('a')
	line, err := b.ReadBytes(delim)
	fmt.Println(line, err)

	// ReadLine是一个底层的原始读取命令。许多调用者或许会使用ReadBytes('\n')或者ReadString('\n')来代替这个方法。
	// ReadLine尝试返回单个行，不包括行尾的最后一个分隔符。如果一个行大于缓存，调用的时候返回了ifPrefix， 就会返回行的头部。行剩余的部分就会在下次调用的时候返回。
	// 当调用行的剩余的部分的时候，isPrefix将会设为false， 返回的缓存只能在下次调用ReadLine的时候看到。
	// ReadLine会返回了一个非空行，或者返回一个error， 但是不会两者都返回。
	// ReadLine返回的文本不会包含行结尾（"\r\n"或者"\n"）。
	// 如果输入没有最终的行结尾的时候，不会返回 任何迹象或者错误。
	// 在 ReadLine 之后调用 UnreadByte 将总是放回读取的最后一个字节 （可能是属于该行末的字符），即便该字节并非 ReadLine 返回的行的一部分。
	line, isPrefix, err := b.ReadLine()
	fmt.Println(line, isPrefix, err)

	// ReadRune读取单个的UTF-8编码的Unicode字节，并且返回rune和它的字节大小。
	// 如果编码的rune是可见的，它消耗一个字节并且返回1字节的unicode.ReplacementChar (U+FFFD)。
	r, size, err := b.ReadRune()
	fmt.Println(r, size, err)

	// ReadSlice从输入中读取，直到遇到第一个终止符为止，返回一个指向缓存中字节的slice。
	// 在下次调用的时候这些字节就是已经被读取了。
	// 如果ReadSlice在找到终止符之前遇到了error， 它就会返回缓存中所有的数据和错误本身（经常是 io.EOF）。
	// 如果在终止符之前缓存已经被充满了，ReadSlice会返回ErrBufferFull错误。
	// 由于ReadSlice返回的数据会被下次的I/O操作重写，因此许多的客户端会选择使用ReadBytes或者ReadString代替。
	// 当且仅当数据没有以终止符结束的时候，ReadSlice返回err != nil
	line, err = b.ReadSlice(delim)
	fmt.Println(line, err)

	// ReadString读取输入到第一次终止符发生的时候，返回的string包含从当前到终止符的内容（包括终止符）。
	// 如果ReadString在遇到终止符之前就捕获到一个错误，它就会返回遇到错误之前已经读取的数据，和这个捕获 到的错误（经常是 io.EOF）。
	// 当返回的数据没有以终止符结束的时候，ReadString返回err != nil。 对于简单的使用，或许 Scanner 更方便。
	lineString, err := b.ReadString(delim)
	fmt.Println(lineString, err)

	// Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据。
	defer b.Reset(file)

	// Size返回底层缓冲区的大小（字节）。
	bSize := b.Size()
	fmt.Println(bSize)

	// UnreadByte将最后的字节标志为未读。只有最后的字节才可以被标志为未读。
	err = b.UnreadByte()
	fmt.Println(err)

	// UnreadRune将最后一个rune设置为未读。如果最新的在buffer上的操作不是ReadRune，则UnreadRune 就返回一个error。
	//（在这个角度上看，这个函数比UnreadByte更严格，UnreadByte会将最后一个读取 的byte设置为未读。）
	err = b.UnreadRune()
	fmt.Println(err)

	// WriteTo实现了io.WriterTo。
	nn, err := b.WriteTo(file)
	fmt.Println(nn, err)

	// NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner)

	// Bytes方法返回最近一次Scan调用生成的token。
	// 底层数组指向的数据可能会被下一次Scan的调用重写。
	scannerBytes := scanner.Bytes()
	fmt.Println(scannerBytes)

	// Err返回Scanner遇到的第一个非EOF的错误。
	err = scanner.Err()
	fmt.Println(err)

	// Split设置该Scanner的分割函数。本方法必须在Scan之前调用。
	scanner.Split(bufio.ScanWords)

	// Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得）， 并让Scanner的扫描位置移动到下一个token。
	// 当扫描因为抵达输入流结尾或者遇到错误而停止时， 本方法会返回false。在Scan方法返回false后， Err方法将返回扫描时遇到的任何错误；除非是io.EOF，此时Err会返回nil。
	// 若 split 函数返回了 100 个空标记而没有推进输入，那么它就会派错（panic）。这是 scanner 的一个常见错误。
	boolScan := scanner.Scan()
	fmt.Println(boolScan)

	// Bytes方法返回最近一次Scan调用生成的token， 会申请创建一个字符串保存token并返回该字符串。
	scannerString := scanner.Text()
	fmt.Println(scannerString)

	// NewWriter返回一个新的，有默认尺寸缓存的Writer。
	writer = bufio.NewWriter(file)
	fmt.Println(writer)

	// NewWriterSize返回一个新的Writer，它的缓存一定大于指定的size参数。
	// 如果io.Writer参数已经是足够大的有缓存的Writer了，函数就会返回它底层的Writer。
	writerSize := bufio.NewWriterSize(writer, 10)
	fmt.Println(writerSize)

	// Available返回buffer中有多少的字节数未使用。
	available := writerSize.Available()
	fmt.Println(available)

	// AvailableBuffer 返回具有 b.Available() 容量的空缓冲区。
	// 此缓冲区旨在附加到并传递给立即后续的 Write 调用。
	// 缓冲区仅在 writerSize 上的下一次写操作之前有效。
	availableBuffer := writerSize.AvailableBuffer()
	fmt.Println(availableBuffer)

	// Buffered返回已经写入到当前缓存的字节数。
	buffered := writerSize.Buffered()
	fmt.Println(buffered)

	// Flush将缓存上的所有数据写入到底层的io.Writer中。
	err = writerSize.Flush()
	fmt.Println(err)

	// ReadFrom实现了io.ReaderFrom。
	readFrom, err := writerSize.ReadFrom(file)
	fmt.Println(readFrom, err)

	// 重置丢弃任何未刷新的缓冲数据，清除所有错误，并重置 b 以将其输出写入 w。
	// 对 Writer 的零值调用 Reset 会将内部缓冲区初始化为默认大小。
	writerSize.Reset(file)

	// Size 返回底层缓冲区的大小（以字节为单位）。
	sizeTwo := writerSize.Size()
	fmt.Println(sizeTwo)

	// Write 将 a 的内容写入缓冲区。它返回写入的字节数。如果 aa < len(a)，它还会返回一个错误，解释为什么写入很短。
	aa, err := writerSize.Write(a)
	fmt.Println(aa, err)

	// WriteByte 写入单个字节。
	err = writerSize.WriteByte(delim)
	fmt.Println(err)

	// WriteRune 写入单个 Unicode 代码点，返回写入的字节数和任何错误。
	runeSize, err := writerSize.WriteRune(2)
	fmt.Println(runeSize, err)

	// WriteString 写入一个字符串。它返回写入的字节数。
	// 如果计数小于 len(s)，它还会返回一个错误，解释为什么写入很短。
	isInt, err := writerSize.WriteString("sss")
	fmt.Println(isInt, err)

}
