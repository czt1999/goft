package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 2
	fmt.Printf("type of a: %T, value of a: %v \n", a, a)

	t, v := reflect.TypeOf(a), reflect.ValueOf(a)

	fmt.Printf("type of TypeOf(a): %T, value of TypeOf(a): %v \n", t, t)
	fmt.Printf("type of ValueOf(a): %T, value of ValueOf(a): %v \n", v, v)
}
