package main

import "fmt"

func inc(sli []int, n int) []int {
	sli = append(sli, n+1)
	return sli
}

func main() {
	x := []int{1, 2, 3}

	// for _, poop := range x {
	// 	x = inc(x, poop)
	// 	fmt.Println(x)
	// }
	moist := len(x)
	for z := 0; z < len(x); z++ {
		x = inc(x, 5)
		moist = len(x)
		fmt.Println(x)
		if moist == 10 {
			break
		}
	}
}
