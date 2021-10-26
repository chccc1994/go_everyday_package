package main

import (
	"fmt"

	"github.com/unknwon/com"
)

func main() {
	var f = "12.3"
	str := com.StrTo(f).MustFloat64()
	fmt.Printf("TYPE = %T ,VALUE=%f\n", str, str)
	var d = "10"
	fmt.Printf("type= %T,value =%d\n", com.StrTo(d).MustInt(), com.StrTo(d).MustInt())
}
