package main

import (
	"fmt"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	for _,arg:= range medals {
		fmt.Println(arg)
	}
}