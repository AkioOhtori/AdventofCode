package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "sample.txt" //path to problem input
const PART int = 1

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func fuckingFuck(i [][]int, x int, y int, a [][2]int) [][2]int {
	val := i[y][x]
	if val == 9 {
		return a
	}

	//Do xr (x + 1) first
	xr := x + 1
	if xr < len(i[y]) { //make sure it is in bounds, if not, skip
		if i[y][xr] > val && i[y][xr] != 9 {
			if checkAnswer(a, [2]int{y, xr}) {
				a = append(a, [2]int{y, xr})
				a = fuckingFuck(i, xr, y, a)
			}
		}
	}
	yd := y + 1
	if yd < len(i) {
		if i[yd][x] > val && i[yd][x] != 9 {
			if checkAnswer(a, [2]int{yd, x}) {
				a = append(a, [2]int{yd, x})
				a = fuckingFuck(i, x, yd, a)
			}
		}
	}
	xl := x - 1
	if xl >= 0 {
		if i[y][xl] > val && i[y][xl] != 9 {
			if checkAnswer(a, [2]int{y, xr}) {
				a = append(a, [2]int{y, xr})
				a = fuckingFuck(i, xl, y, a)
			}
		}
	}
	yu := y - 1
	if yu >= 0 {
		if i[yu][x] > val && i[y][xl] != 9 {
			if checkAnswer(a, [2]int{yu, x}) {
				a = append(a, [2]int{yu, x})
				a = fuckingFuck(i, x, yu, a)
			}
		}
	}

	return a
}

func checkAnswer(a [][2]int, new [2]int) bool {
	for _, i := range a {
		if i == new {
			return false
		}
	}
	return true
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
	var ans_pt1 []int
	var ans_pt2 [][2]int
	var omgwtf []int

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
					ans_pt1 = append(ans_pt1, val) //add to ans_pt1 array
					ans_pt2 = fuckingFuck(input_int, x, y, ans_pt2)
					omgwtf = append(omgwtf, len(ans_pt2))
					ans_pt2 = [][2]int{}
				}
			} //end ifs
		} //end x for loop
	} // end y for loop

	//could have done this inline but heh
	var answer int = 0
	// var answer2 int = 0
	for _, a := range ans_pt1 {
		answer += a + 1
	}

	fmt.Printf("The answer to Part 1 is %v danger units\n", answer)

	fmt.Println(len(ans_pt2))
	fmt.Println(omgwtf)

}
