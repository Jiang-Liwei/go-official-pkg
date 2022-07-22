package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Lstat("./tar.go")
	/**
	 * FileInfoHeader 中返回的结构体详情
	 *  type Header struct {
	 *		Name       string    // name of header file entry // 记录头域的文件名
	 *		Mode       int64     // permission and mode bits // 权限和模式位
	 *		Uid        int       // user id of owner // 所有者的用户ID
	 *      Gid        int       // group id of owner // 所有者的组ID
	 *		Size       int64     // length in bytes // 字节数（长度）
	 *		ModTime    time.Time // modified time // 修改时间
	 *		Typeflag   byte      // type of header entry // 记录头的类型
	 *		Linkname   string    // target name of link // 链接的目标名
	 *		Uname      string    // user name of owner // 所有者的用户名
	 *		Gname      string    // group name of owner // 所有者的组名
	 *		Devmajor   int64     // major number of character or block device // 字符设备或块设备的major number
	 *		Devminor   int64     // minor number of character or block device // 字符设备或块设备的minor number
	 *		AccessTime time.Time // access time // 访问时间
	 *		ChangeTime time.Time // status change time // 状态改变时间
	 *		Xattrs     map[string]string
	 *	}
	 */

	/**
	 * FileInfoHeader返回一个根据fi填写了部分字段的Header。
	 * 如果fi描述一个符号链接，FileInfoHeader函数将link参数作为链接目标。
	 * 如果fi描述一个目录，会在名字后面添加斜杠。
	 * 因为os.FileInfo接口的Name方法只返回它描述的文件的无路径名， 有可能需要将返回值的Name字段修改为文件的完整路径名。
	 */

	header, err := tar.FileInfoHeader(file, "sss")
	if err != nil {
		fmt.Println(err)
		return
	}

	//FileInfo返回该Header对应的文件信息。（os.FileInfo类型）
	fileInfo := header.FileInfo()
	if fileInfo == nil {
		return
	}
	fmt.Println("FileInfo 的结果：", fileInfo.Name())

	// NewReader创建一个从buf读取的Reader。
	var buf bytes.Buffer
	reader := tar.NewReader(&buf)

	next, err := reader.Next()
	fmt.Println("Next 的结果：", next)

	tarCreate, err := os.Create("./one.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(tar *os.File) {
		_ = tarCreate.Close()
	}(tarCreate)

	tw := tar.NewWriter(tarCreate)
	defer func(tw *tar.Writer) {
		_ = tw.Close()
	}(tw)

	// Flush结束当前文件的写入。（可选的）
	defer func(tw *tar.Writer) {
		err = tw.Flush()
	}(tw)
	// Write向tar档案文件的当前记录中写入数据。
	// 如果写入的数据总数超出上一次调用WriteHeader的参数hdr.Size字节， 返回ErrWriteTooLong错误。
	write, err := tw.Write([]byte("sss"))
	if err != nil {
		return
	}
	fmt.Println(write)
	// WriteHeader写入hdr并准备接受文件内容。
	// 如果不是第一次调用本方法，会调用Flush。
	// 在Close之后调用本方法会返回ErrWriteAfterClose
	_ = tw.WriteHeader(header)
}
