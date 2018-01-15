package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, arg:= range os.Args[1:] {
		data, err:= ioutil.ReadFile(arg)
		if err != nil {
			continue
		}
		for _, line:= range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n:= range counts {
		fmt.Printf("%s\t%d\n", line, n)
	}
}