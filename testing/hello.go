package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for x := range a {
		fmt.Println(x)
	}
	for x, _ := range a {
		fmt.Println(x)
	}
}
