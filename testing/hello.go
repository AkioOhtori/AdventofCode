package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "fold along y=13"
	y := strings.Split(a, "=")
	fmt.Println(y)
	fmt.Println(y[1])
	fmt.Println(y[0][11:])

}
