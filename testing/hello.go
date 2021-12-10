package main

import (
	"fmt"
)

func main() {
	x := []string{"a", "b", "c"}
	fmt.Println(len(x))
	x = append(x[1:])
	fmt.Println(len(x))
	fmt.Println((x))
}
