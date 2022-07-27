package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {

}

func zipFunc() {
	zipFile, _ := os.Create("./one.zip")

	// 关闭zip文件，使其无法作用 I/O
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {
			fmt.Println("文件关闭失败咯")
		}
	}(zipFile)

	zipWriter := zip.NewWriter(zipFile)

	// 关闭zip文件的写入
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			fmt.Println("文件写入关闭关闭失败咯")
		}
	}(zipWriter)
	// remove the trailing path separator if path is a directory
	srcPath := strings.TrimSuffix("./use.go", string(os.PathSeparator))

	// visit all the files or directories in the tree
	err := filepath.Walk(srcPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(srcPath), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}

		// create writer for the file header and save content of the file
		headerWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(headerWriter, f)
		return err
	})

	if err != nil {
		fmt.Println("失败咯")
	}

}

func unzip() {

}
