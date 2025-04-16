package main

import "C"

//export Add
func Add(a, b int) int {
	return a + b
}

//export Subtract
func Subtract(a, b int) int {
	return a - b
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

//export Divide
func Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {}
