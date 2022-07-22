package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	createTar()
	unTar()
}

// tar 压缩
func createTar() {

	// 创建一个 tar 文件
	f, err := os.Create("./output.tar")
	if err != nil {
		panic(err)
		return
	}

	// 使用 defer
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// NewWriter 创建一个写入w的*Writer。
	tw := tar.NewWriter(f)
	defer func(tw *tar.Writer) {
		_ = tw.Close()
	}(tw)
	fileInfo, err := os.Stat("./main.exe") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Println(err)
	}
	err = tw.WriteHeader(hdr) //写入头文件信息
	if err != nil {
		fmt.Println(err)
	}
	f1, err := os.Open("./main.exe")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将main.exe文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

// tar 解压
func unTar() {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileInfo := hdr.FileInfo()
		fmt.Println(fileInfo.Name())
		f, err := os.Create("123" + fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}
