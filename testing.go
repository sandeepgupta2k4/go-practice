package main

import (
	"fmt"
)

var i = sum(1, 4)

func sum(first int, second int) int {
	fmt.Println("Inside Sum")
	return first + second
}

func main() {
	fmt.Println("Inside main" + string(i))
}
