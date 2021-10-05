package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

// see: https://golangbyexample.com/understanding-uintptr-golang/

func main() {
	s := &sample{a: 1, b: "hello"}

	// 1.
	// One purpose of uintptr is to be used along with unsafe.Pointer for unsafe memory access.
	// Arithmetic operations cannot be performed on unsafe.Pointer. To perform such arithmetic,
	// 1) converts unsafe.Pointer to uintptr
	// 2) performs arithmetic on uintptr
	// 3) converts back uintptr to unsafe.Pointer

	// Getting the address of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	// Typecasting it to a string pointer and printing the value of it
	fmt.Println(*(*string)(p))

	// 2.
	// Another purpose of uintptr is when you want to save the pointer address value
	// for printing it or storing it.

	// Getting the address as an uintptr
	startAddress := uintptr(unsafe.Pointer(s))
	fmt.Printf("Start Address of s: %d\n", startAddress)

	// Notes that since the address is just stored and does not reference anything,
	// the corresponding object can be garbage collected.
}
