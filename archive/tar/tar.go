package main

import (
	"archive/tar"
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
	fmt.Println(fileInfo.Name())
}
