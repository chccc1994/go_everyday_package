package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age`
}

func main() {
	p := person{
		Name: "wangsan",
		Age:  89,
	}
	//pJson, err := json.Marshal(p)
	pJson, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%v\n", string(pJson))
}
