package main

import "fmt"

// Day 1
// May 20, 2020
// Function Declaration
func main() {
	fmt.Println("Hello World")

	// Calling a function
	functionSample()

	// function returning value
	i := getInt()
	fmt.Println(i)

	// function taking input param
	s := add(3, 4)
	fmt.Println(s)

	m := multiply(4, 5)
	fmt.Println(m)

	a, b := multipleReturnValues(6, 7)
	fmt.Println(a, b)

	c := namedReturnValue(9)
	fmt.Println(c)

	conditions()
	loop()
	dataStructures()
}

func functionSample() {
	fmt.Println("Prints something")
}

func getInt() int {
	return 1
}

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func multipleReturnValues(a, b int) (int, int) {
	return a, b
}

func namedReturnValue(a int) (b int) {
	b = a
	return
}

// conditions

func gain(a int) int {
	if a == 10 {
		return 5
	} else if a == 6 {
		return 10
	} else {
		return 3
	}
}

func switchExample(a int) {
	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 5:
		fmt.Println("Medium")
	default:
		fmt.Println("big")
	}
}

func conditions() {
	fmt.Println(gain(10))
	fmt.Println(gain(6))
	fmt.Println(gain(2))
	switchExample(1)
	switchExample(2)
	switchExample(4)
	switchExample(10)
}

// loops
func forExample() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func whileForm() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func loop() {
	forExample()
	whileForm()
}

// Basic Data Structures

func dataStructures() {
	// slice
	// map
}
