package main

import (
	"fmt"
	"strconv"
)

// 转 int
// 将整数转换为字符串形式。base 表示转换进制，取值在 2 到 36 之间。
// 结果中大于 10 的数字用小写字母 a - z 表示。
func StringToInt(value string) (result int) {
	result, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return

}

// int 转 string
func IntToString(value int) (result string) {
	result = strconv.Itoa(value)
	return
}

// 字符串转Bool
// 将字符串转换为布尔值
// 它接受真值：1, t, T, TRUE, true, True
// 它接受假值：0, f, F, FALSE, false, False
func StringToBool(value string) (result bool) {
	result, _ = strconv.ParseBool(value)
	return
}

// bool转字符串
func BoolToString(value bool) (result string) {
	result = strconv.FormatBool(value)
	return
}

// string to float
func StringToFloat(value string) (result float64) {
	result, _ = strconv.ParseFloat(value, 32)
	return
}

// float to string
func FloatToString(value float64) (result string) {
	result = strconv.FormatFloat(value, 'E', -1, 32)
	return
}

func main() {
	//1. to int
	fmt.Printf("type:%T ,value: %d\n", StringToInt("250"), StringToInt("250"))
	// 2 int to string
	fmt.Printf("type:%T ,value: %s\n", IntToString(250), IntToString(250))
	//3. string to bool
	fmt.Printf("type:%T ,value: %t\n", StringToBool("abs"), StringToBool("TRUE"))
	// 4. bool to string
	fmt.Printf("type:%T ,value: %s\n", BoolToString(true), BoolToString(true))

	// 5. string to float
	fmt.Printf("type:%T ,value: %f\n", StringToFloat("12.3"), StringToFloat("12.3"))
	//6. float to string
	fmt.Printf("type:%T ,value: %s\n", FloatToString(12.3), FloatToString(12.3))
}

// https://www.cnblogs.com/golove/p/3262925.html
