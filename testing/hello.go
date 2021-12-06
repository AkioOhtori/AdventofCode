package main

import (
	"fmt"
)

func main() {
	var m1 = [2]int{1, 2}
	var m2 = [2]int{3, 4}
	// var mm [][]int
	mm := make([][]int, 0, 100)

	mm[0] = m1[:] // append(mm, m1)
	mm[1] = m2[:] // append(mm, m2)
	fmt.Print(mm)
}
