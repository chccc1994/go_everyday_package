package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()

	file, _ := os.OpenFile(wd+"/hello.txt", os.O_WRONLY, 777) // 打开文件，os.O_WRONLY表示以只写的方式打开，777表示文件的权限，以最大的权限打开

	n, err := file.Write([]byte("hello world!"))
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("n:", n)
	}
	ret, err := file.WriteString("dalgurak")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("ret", ret)
	}

	// file.WriteAt 方法在指定位置写入数据，这个很少使用，不做展示，大家自行测试

	file.Close() // 不要忘记关闭文件
}
