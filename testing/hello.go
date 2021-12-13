package main

import "fmt"

func main() {
	x := []string{"a", "B", "cd", "AB"}
	y := make([]string, len(x))
	fmt.Println(x)
	fmt.Println(y)
	copy(y, x)
	fmt.Println(x)
	fmt.Println(y)

}
