package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Based on the puzzle input, seems like our best bet is to do some matrix math?
First up though: Parsing the input - Medium
Second: iterating though to create bit mask
Third: don't fuck it up
*/

var path = "sample.txt" //path to problem input
const PART int = 1

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func convTxtSlicetoInt(arr []string) []int {
	var output []int
	for _, txt := range arr {
		val, _ := strconv.Atoi(txt)
		output = append(output, val)
	}
	return output
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

	var fishies [18]int
	var buffer [8]int

	for _, fish := range input_int {
		fishies[fish]++
		// switch fish {
		// case 0:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 1:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 2:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 3:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 4:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 5:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 6:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 7:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// case 8:
		// 	fmt.Println(fish)
		// 	fishies[n]++
		// }
	}

	for day := 0; day < 80; day++ {
		if day < 7 {
			fishies[day+8] += fishies[day]
		} else {
			cycle := day % 6
			fishies[cycle] += fishies[cycle+6]
			fishies[cycle+6] = 0
			fishies[cycle+8] += fishies[cycle]

		}
		fmt.Printf("%v\t%v\n", day, fishies)
	}
	var total int = 0
	for _, fish := range fishies {
		total += fish
	}

	fmt.Println(fishies)

	fmt.Println(total)
}
