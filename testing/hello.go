package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} //len 9, mid = "5"
	fmt.Println(a[:5])
	fmt.Println(a[5:])

}
