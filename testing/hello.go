package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	b := a
	b["one"] = 5
	fmt.Println(a)
	fmt.Println(b)

}
