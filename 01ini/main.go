package main

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

func main() {

	cfg, err := ini.Load("conf/app.ini") //	加载ini 配置文件
	if err != nil {
		log.Fatalf("配置文件读取错误 'conf/app.ini': %v", err)
	}
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值
	// https://ini.unknwon.io/docs/howto/work_with_values
	fmt.Println("", cfg.Section("").Key("RUN_MODE").MustString("debug"))
	fmt.Println("RUN MODE:", cfg.Section("").Key("RUN_MODE").String())

	fmt.Println("mysql:", cfg.Section("database").Key("TYPE").String())
	fmt.Println("mysql:", cfg.Section("database").Key("HOST").String())
	fmt.Println("mysql:", cfg.Section("database").Key("NAME").String())
}
