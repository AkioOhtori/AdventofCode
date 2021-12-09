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

	//Load all lines in file into a slice of slices
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_int [][]int
	var answers []int

	//Go through the instructions and make them usable
	for scanner.Scan() {
		var temp []int
		input_str := strings.Split(scanner.Text(), "")
		for _, x := range input_str {
			y, _ := strconv.Atoi(x)
			temp = append(temp, y)
		}
		input_int = append(input_int, temp)
	}

	for y := 0; y < len(input_int); y++ {
		for x := 0; x < len(input_int[y]); x++ {
			//need x+1, x-1, y+1, y-1
			//setting up up, down, left, and right
			xl := x - 1
			xr := x + 1
			yu := y - 1
			yd := y + 1
			//if index is out of bounds, set it to the other one so checks still work
			if xr >= len(input_int[y]) {
				xr = xl
			} else if xl < 0 {
				xl = xr
			}
			if yu < 0 {
				yu = yd
			} else if yd >= len(input_int) {
				yd = yu
			}

			//check 'em!
			val := input_int[y][x]
			if val < input_int[yu][x] && val < input_int[yd][x] {
				if val < input_int[y][xl] && val < input_int[y][xr] {
					answers = append(answers, val) //add to answers array
				}
			} //end ifs
		} //end x for loop
	} // end y for loop

	//could have done this inline but heh
	var answer int = 0
	for _, a := range answers {
		answer += a + 1
	}
	fmt.Printf("The answer to Part 1 is %v danger units\n", answer)

}
