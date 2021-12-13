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

func incrementOctopuses(octo [][]int) (tens [][2]int) {
	// var tens [][2]int
	for y := 0; y < len(octo); y++ {
		for x := 0; x < len(octo[y]); x++ {
			octo[y][x]++
			if octo[y][x] > 9 {
				tens = append(tens, [2]int{y, x})
			}
		}
	}
	return
}

func incrementAdjacent(o [][]int, n [][2]int, y int, x int) [][2]int {
	length := len(o)
	var incrementMap [][2]int

	//Generates a list of valid coordinates for increment
	for yy := -1; yy <= 1; yy++ {
		for xx := -1; xx <= 1; xx++ {
			xxx := xx + x
			yyy := yy + y
			if xx == 0 && yy == 0 {
				continue
			} else if xxx == length || yyy == length {
				continue
			} else if xxx == -1 || yyy == -1 {
				continue
			} else {
				incrementMap = append(incrementMap, [2]int{yyy, xxx})
			}
		}
	}
	//Iterates over the previously created coordinates and increments
	for _, yx := range incrementMap {
		o[yx[0]][yx[1]]++
		if o[yx[0]][yx[1]] == 10 {
			if !(checkNew(n, [2]int{yx[0], yx[1]})) {
				n = append(n, [2]int{yx[0], yx[1]})
			}
		}
	}
	return n //Return the new "tens" matrix
}

// Function to better print matrices
func prettyPrintSlice(m [][]int) {
	fmt.Printf("\n")
	for _, y := range m {
		fmt.Println(y)
	}
	fmt.Printf("\n")
}

//Checks new coord to make sure we don't already have them
func checkNew(a [][2]int, new [2]int) bool {
	for _, i := range a {
		if i == new {
			return true
		}
	}
	return false
}

func main() {
	//Open input file
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	//Load all lines in file into a slice of slices
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_int [][]int

	//Go through the instructions and convert them to slices of ints, [y],[x]
	for scanner.Scan() {
		var temp []int
		input_str := strings.Split(scanner.Text(), "")
		for _, x := range input_str {
			y, _ := strconv.Atoi(x)
			temp = append(temp, y)
		}
		input_int = append(input_int, temp)
	}

	/*
		From the prompt, every day 3 things happen:
		1) Every octopus gets a +1
		2) Anything with an energy of 9 flashes and adds 1 to adjacent
		3) Anything that flashed goes to 0
	*/
	var answer_pt1 int = 0

	for iterations := 0; iterations < 10000; iterations++ {

		// Step 1 - Increment all and record any flashes
		tens := incrementOctopuses(input_int)

		// Step 2 - Increment everything around the flashes AND process new ones
		for i := 0; i < len(tens); i++ {
			tens = incrementAdjacent(input_int, tens, tens[i][0], tens[i][1])
		}
		answer_pt1 += len(tens)
		// Step 3 - Zero out anything that flashed, regardless of current value
		for _, flash := range tens {
			input_int[flash[0]][flash[1]] = 0
		}
		if len(tens) == 100 { //all octopi flashed
			prettyPrintSlice(input_int)
			fmt.Printf("The answer to Part 2 is step %v\n", iterations+1)
			break
		}
		if iterations == 99 {
			fmt.Printf("The answer to part 1, how many flashes after 100 cycles, is %v\n", answer_pt1)
		}

	}

}
