package main

import "os"

func main() {
	wd, _ := os.Getwd()

	file, _ := os.Open(wd + "/main.go") // 以只读的方式打开文件

	// 获取文件的信息
	fInfo, _ := file.Stat()
	println("是否是一个目录:", fInfo.IsDir())
	println("文件的修改时间:", fInfo.ModTime().String())
	println("文件的名字:", fInfo.Name())
	println("文件的大小:", fInfo.Size())
	println("文件的权限:", fInfo.Mode().String())

	file.Close() // 不要忘记关闭文件
}
