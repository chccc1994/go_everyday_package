package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file, err := os.Open(wd + "/hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, 4) // 文件内容不多，我们一次性读4个字节，多读几次，不一次性读完
	var str string
	for {
		n, err := file.Read(b)
		if err != nil {
			if err == io.EOF { // EOF表示文件读取完毕
				break // 退出
			}
		}
		str += string(b[:n]) // 保存文件内容
	}
	println(str) // 打印文件

	file.Close() // 不要忘记关闭文件
}
