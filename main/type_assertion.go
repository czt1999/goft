package main

import "fmt"

func main() {
	var i interface{} = 1
	// 断言i的类型为int并取其值
	t1, ok := i.(int)
	fmt.Printf("%d %t\n", t1, ok) // 1 true
	// 断言i的类型为string并取其值
	t2, ok := i.(string)
	fmt.Printf("%s %t\n", t2, ok) //  false
}
