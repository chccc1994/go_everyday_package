package main

import (
	"fmt"
	"os"
)

func main() {
	// 为了减少代码的篇幅，基本所有的错误在这篇文字里面我都丢弃
	wd, _ := os.Getwd()
	fmt.Println("获取当前工作目录的根路径:", wd)

	os.Chdir("/home/chccc/workspace/go/src")
	w, _ := os.Getwd()
	fmt.Println("获取x修改后的当前工作目录的根路径:", w)

	os.Chdir(wd) // 将工作目录改回原来的目录

	os.Chmod(wd+"/api.ini", 0777) // 修改文件的权限

	var err error
	_, err = os.Stat(wd + "/api.ini")
	if err != nil {
		fmt.Println("os.Stat():", err.Error())
		return
	}

	// 文件的权限修改，删除，移动等不做演示，这都是基本命令，不浪费太多篇幅
	err = os.Remove(wd + "/api.ini")
	if err != nil {
		fmt.Println("file remove...\n")
		panic(err)
	}
	// 创建文件
	file, _ := os.Create(wd + "/api.ini")
	fmt.Printf("%v\n", file.Name())

}

// golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:如果返回的错误为nil,
// 说明文件或文件夹存在如果返回的错误类型使用os.IsNotExist()判断为true,
// 说明文件或文件夹不存在如果返回的错误为其它类型,则不确定是否在存在
