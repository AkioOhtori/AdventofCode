package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
const PART int = 2

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

	var fishies [7]int //where we store the old fish
	var fry [9]int     //where we store the new fish

	//Process the input file of already alive fishies ("old")
	for _, fish := range input_int {
		fishies[fish]++
	}

	time := 256
	if PART == 1 {
		time = 80
	}

	//Start makin' babies!
	for day := 0; day < time; day++ {
		cycle := day % 7
		puberty := day % 9
		fishies[cycle] += fry[puberty]
		fry[puberty] = fishies[cycle]
		// fmt.Printf("%v\t%v\t%v\t%v\n", day, fishies, fry, puberty)
	}

	//How many total fishies did we end up with?
	var total int = 0
	for _, fish := range fishies {
		total += fish
	}
	for _, fish := range fry {
		total += fish
	}

	//fin (lol)
	fmt.Println(total)
}
