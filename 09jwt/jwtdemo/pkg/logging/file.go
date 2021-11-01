package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 文件创建等
var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}
func mkDir(filePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+filePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func openLogFile(filePath string) *os.File {

	_, err := os.Stat(filePath)
	// if err != nil {

	// }
	switch {
	case os.IsNotExist(err):
		mkDir(getLogFilePath())
	case os.IsPermission(err):
		log.Fatal("Permission:%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return handle
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}
