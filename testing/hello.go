package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "ade"
	s2 := "de"
	fmt.Println(strings.ReplaceAll(s1, s2, ""))
}
