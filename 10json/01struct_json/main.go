package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p1 := Person{"zhoulin", 70}
	p_json, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("error", err.Error)
		return
	}
	fmt.Println(string(p_json))

	str := `{"name":"liwen","age":15}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2) //c传指针为了在中内部修改p2
	fmt.Println(p2)
	fmt.Printf("%v\n", p2)
}
