package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
const PART int = 1

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	//Open input file
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_str []string
	var input_int []int
	//Go through the instructions and make them usable
	for scanner.Scan() {
		input_str = strings.Split(scanner.Text(), ",")
	}

	for _, x := range input_str {
		tmp, _ := strconv.Atoi(x)
		input_int = append(input_int, tmp)
	}

	var fishies [7]int
	var fry [9]int

	for _, fish := range input_int {
		fishies[fish]++
	}

	for day := 0; day < 80; day++ {
		cycle := day % 7
		puberty := day % 9
		fishies[cycle] += fry[puberty]
		fry[puberty] = fishies[cycle]
		// fmt.Printf("%v\t%v\t%v\t%v\n", day, fishies, fry, puberty)
	}
	var total int = 0
	for _, fish := range fishies {
		total += fish
	}
	for _, fish := range fry {
		total += fish
	}

	fmt.Println(total)
}
