package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(a[:1])
	fmt.Println(a[1:])

	b := "01"
	fmt.Println(b[:1])
	fmt.Println(b[1:])

	fmt.Println((math.Floor(float64(7) / 2.0)))
	fmt.Println((math.Ceil(float64(7) / 2.0)))
}
