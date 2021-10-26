### [strconv 包](https://studygolang.com/pkgdoc)
>该包主要实现基本数据类型与其字符串表示的转换。常用函数为Atoi()、Itia()、parse系列、format系列、append系列。

>大概可以得出答案：基本的数据类型指的是：布尔类型、数值型（整型、浮点型）
其他数据类型转换为字符串的函数多以：Format 为关键字
字符串转换为其他数据类型的函数多以：Parse 为关键字

```go 
func Atoi(s string) (int, error) //Atoi是ParseInt(s, 10, 0)的简写。
func CanBackquote(s string) bool
func FormatBool(b bool) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
// 函数将浮点数表示为字符串并返回。
// bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
// fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
// prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
func IsGraphic(r rune) bool
func IsPrint(r rune) bool
func Itoa(i int) string //Itoa是FormatInt(i, 10) 的简写。
func ParseBool(str string) (bool, error)
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
```

>格式化打印占位符：
    %v:  原样输出
    %T: 打印类型
    %t: bool类型
    %s: 字符串
    %f: 浮点
    %d: 10进制的整数
    %b: 2进制的整数
    %o: 8进制
    %x: %X，16进制
        %x：0-9，a-f
        %X：0-9，A-F
    %c: 打印字符
    %p: 打印地址

