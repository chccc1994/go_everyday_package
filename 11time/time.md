## TIME
[1.Go语言基础之time包](https://www.liwenzhou.com/posts/Go/go_time/)
>time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

time.Time类型表示时间。我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。示例代码如下：
```go
func timeDemo() {
    now := time.Now() //获取当前时间
    fmt.Printf("current time:%v\n", now)

    year := now.Year()     //年
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
    second := now.Second() //秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```
+ 时间戳
```go
time.Unix(sec, nsec int64) //通过 Unix 时间戳生成 time.Time 实例；
time.Time.Unix() //得到 Unix 时间戳；
time.Time.UnixNano()// 得到 Unix 时间戳的纳秒表示；
```
