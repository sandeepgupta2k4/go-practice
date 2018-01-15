package main

import (
	"fmt"
)
// Test does
func Test() int {
	return 10
}

// Fibo returns nth fibonacci number
func Fibo(n int) int{
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(y)
		x, y = y, x+y
	}
	return y
}

// Fact returns factorial
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n - 1)
}