package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name" label:"Name is: "`
	Age    int    `json:"age" label:"Age is: "`
	Gender string `json:"gender,omitempty" label:"Gender is: "`
}

func main() {
	var p = Person{
		Name: "Tom",
		Age:  20,
	}
	bytes, _ := json.Marshal(p)
	jsonStr := string(bytes)
	fmt.Println(jsonStr)
}
