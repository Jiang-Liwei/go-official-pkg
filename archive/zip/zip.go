package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	zipFile, _ := os.Create("./one.zip")
	// NewReader返回一个从r读取数据的*Reader，r被假设其大小为size字节。
	w := zip.NewWriter(zipFile)

	// RegisterCompressor使用指定的方法ID注册一个Compressor类型函数。
	// 常用的方法Store和Deflate是内建的。
	// 这里内建了一个 Deflate
	w.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestCompression)
	})

	// RegisterDecompressor使用指定的方法ID注册一个Decompressor类型函数。

	// zip.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
	//	 return flate.NewReader(r)
	// })

	zipGo, _ := os.Lstat("./zip.go")
	// FileInfoHeader返回一个根据fi填写了部分字段的Header。
	// 因为os.FileInfo接口的Name方法只返回它描述的文件的无路径名， 有可能需要将返回值的Name字段修改为文件的完整路径名。

	zipFileHeader, _ := zip.FileInfoHeader(zipGo)
	// FileInfo返回一个根据h的信息生成的os.FileInfo。
	zipFileInfo := zipFileHeader.FileInfo()
	fmt.Println(zipFileInfo.Name())
	// func (*FileHeader) ModTime 改为 FileHeader.Modified
	//fmt.Println(zipFileHeader.Modified) 打印最近修改时间

	// Mode返回h的权限和模式位。
	fileModel := zipFileHeader.Mode()
	fmt.Println(fileModel)

	// 将ModifiedTime和ModifiedDate字段设置为给定的UTC时间。（精度2s）
	// 已经弃用
	// zipFileHeader.SetModTime(time.Time{})
	// 给与777权限
	zipFileHeader.SetMode(os.FileMode(0777))

	// OpenReader会打开name指定的zip文件并返回一个*ReadCloser。
	zipReader, _ := zip.OpenReader("one.zip")
	fmt.Println(zipReader)

	// 使用给出的文件名添加一个文件进zip文件。 本方法返回一个io.Writer接口（用于写入新添加文件的内容）。
	// 文件名必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔。
	// 新增文件的内容必须在下一次调用CreateHeader、Create或Close方法之前全部写入。
	_, _ = w.Create("./zip.go")
	// 使用给出的*FileHeader来作为文件的元数据添加一个文件进zip文件。
	// 本方法返回一个io.Writer接口（用于写入新添加文件的内容）。
	// 新增文件的内容必须在下一次调用CreateHeader、Create或Close方法之前全部写入。
	_, _ = w.CreateHeader(zipFileHeader)

	// Close方法通过写入中央目录关闭该*Writer。 本方法不会也没办法关闭下层的io.Writer接口。
	_ = zipReader.Close()
	// Close关闭zip文件，使它不能用于I/O。
	_ = w.Close()
	//Flush将所有缓冲数据刷新到底层写入程序。通常不需要调用Flush；呼叫Close就足够了。
	_ = w.Flush()

}

/**
 * DataOffset返回文件的可能存在的压缩数据相对于zip文件起始的偏移量。
 * 大多数调用者应使用Open代替，该方法会主动解压缩数据并验证校验和。
 * func (f *File) DataOffset() (offset int64, err error)
 */
