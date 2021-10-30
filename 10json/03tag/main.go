package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city,omitempty"`
}

func main() {
	amy := person{
		Name:    "liliu",
		Country: "China",
		City:    "Beijing",
	}
	data, _ := json.MarshalIndent(struct {
		person
		Gender string `gender`
	}{
		person: amy,
		Gender: "女",
	}, "", "	")
	fmt.Println(string(data))

}
