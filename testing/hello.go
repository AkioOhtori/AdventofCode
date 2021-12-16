package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "CB"
	b, _ := strconv.ParseUint(a, 16, 32)
	fmt.Println(b)
	bin := fmt.Sprintf("%04b", b)
	fmt.Println(bin)
	// bin += fmt.Sprintf("%04b", b)
	// fmt.Println(bin)
}
