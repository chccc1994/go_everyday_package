package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//Unix 获取指定日期的时间戳
	dt, _ := time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51")
	fmt.Println("dt:", dt.Unix())

}

// current time:2021-10-30 14:42:54.452747131 +0800 CST m=+0.000028988
// 2021-10-30 14:42:54
