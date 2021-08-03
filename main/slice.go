package main

import "fmt"

func main() {
	// 数组（固定长度）
	var arr = [4]int{1, 2, 3, 4}
	fmt.Printf("arr: %v \n", arr)

	// 切片
	var s1 = arr[0:]
	var s2 = arr[2:4]
	fmt.Printf("s1 : %v \n", s1)
	fmt.Printf("s2 : %v \n", s2)

	// s1与s2共享同一底层数组
	s1[2] = -3
	fmt.Printf(">>> set s1[2] to -3 ...\n")
	fmt.Printf("arr: %v \n", arr)
	fmt.Printf("s1 : %v \n", s1)
	fmt.Printf("s2 : %v \n", s2)

	// 触发s1扩容，拷贝返回新的底层数组
	s1 = append(s1, 5)
	fmt.Printf(">>> append 5 to s1, which makes s1 grow ...\n")
	fmt.Printf("slice 1: %v \n", s1)

	// s1已不再指向原先的数组
	s1[3] = -4
	fmt.Printf(">>> set s1[3] to -4 ...\n")
	fmt.Printf("arr: %v \n", arr)
	fmt.Printf("s1 : %v \n", s1)
	fmt.Printf("s2 : %v \n", s2)
}
