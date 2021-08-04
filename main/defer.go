package main

import "fmt"

func main() {
	// 将rec()放在此处无效

	setData(10)

	fmt.Println("main exixts")
}

func setData(x int) {
	defer rec()
	// 数组越界则触发panic
	var arr [10]int
	arr[x] = 1
}

func rec() {
	if err := recover(); err != nil {
		fmt.Println("=========== recover ============")
		fmt.Println(err)
		fmt.Println("================================")
	}
}
