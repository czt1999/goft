package main

import "fmt"

func main() {
	fmt.Println("=========== closure function: adder ============")
	var a = genAdder(0)
	fmt.Printf("adder a(): %d \n", a()) // 1
	fmt.Printf("adder a(): %d \n", a()) // 2
	fmt.Printf("adder a(): %d \n", a()) // 3
	var b = genAdder(100)
	fmt.Printf("adder b(): %d \n", b()) // 101
	fmt.Printf("adder b(): %d \n", b()) // 102
	fmt.Printf("adder b(): %d \n", b()) // 103
	fmt.Printf("adder a(): %d \n", a()) // 4

	fmt.Println("========= function factory: multiplier ==========")
	var multiplyThree = multiplier(3)
	var multiplyFour = multiplier(4)
	fmt.Printf("multiplyThree(4): %d \n", multiplyThree(4)) // 12
	fmt.Printf("multiplyFour(4) : %d \n", multiplyFour(4))  // 16
}

//
// 产生闭包，变量 i 会从栈逃逸到堆上
//
func genAdder(a int) func() int {
	var i = a
	return func() int {
		i++
		return i
	}
}

//
// 函数工厂
//
func multiplier(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}
