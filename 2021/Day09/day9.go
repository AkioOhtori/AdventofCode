package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

//Recurisive function to, essentially, check for 9s and add non-nines to the basin
//Returns the coordinants of all parts of the basin explored so far (b)
func exploreBasin(i [][]int, x int, y int, b [][2]int) [][2]int {
	val := i[y][x]
	if val == 9 {
		return b
	}

	//Check one to the left, right, up, and down to see if they're 9
	//If they're not, add them to the list (if they aren't already there) and explore
	xr := x + 1
	if xr < len(i[y]) { //make sure it is in bounds, if not, skip
		if i[y][xr] != 9 {
			if checkAnswer(b, [2]int{y, xr}) {
				b = append(b, [2]int{y, xr})
				b = exploreBasin(i, xr, y, b)
			}
		}
	}
	yd := y + 1
	if yd < len(i) {
		if i[yd][x] != 9 {
			if checkAnswer(b, [2]int{yd, x}) {
				b = append(b, [2]int{yd, x})
				b = exploreBasin(i, x, yd, b)
			}
		}
	}
	xl := x - 1
	if xl >= 0 {
		if i[y][xl] != 9 {
			if checkAnswer(b, [2]int{y, xl}) {
				b = append(b, [2]int{y, xl})
				b = exploreBasin(i, xl, y, b)
			}
		}
	}
	yu := y - 1
	if yu >= 0 {
		if i[yu][x] != 9 {
			if checkAnswer(b, [2]int{yu, x}) {
				b = append(b, [2]int{yu, x})
				b = exploreBasin(i, x, yu, b)
			}
		}
	}
	return b //send back the expored basin so far
}

//Checks new basin coord to make sure we don't alreay have them
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
	var basin [][2]int
	var basinSizes []int

	//Go through the instructions and convert them to slices of ints
	for scanner.Scan() {
		var temp []int
		input_str := strings.Split(scanner.Text(), "")
		for _, x := range input_str {
			y, _ := strconv.Atoi(x)
			temp = append(temp, y)
		}
		input_int = append(input_int, temp)
	}

	//Time to explore! (Nested for loops to example entire x by y matrix)
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

					//Then deal with Part 2 stuff
					basin = append(basin, [2]int{y, x})          //add center of basin to basin
					basin = exploreBasin(input_int, x, y, basin) //explore the basin
					basinSizes = append(basinSizes, len(basin))  //add the size of the basin to the answers list
					basin = [][2]int{}                           //clear the basin for next checks
				}
			} //end ifs
		} //end x for loop
	} // end y for loop

	//Part 1 Answer (could have done this inline but heh)
	var answer int = 0
	for _, a := range ans_pt1 {
		answer += a + 1
	}
	fmt.Printf("The answer to Part 1 is %v danger units\n", answer)

	//Part 2 Answer Calculation
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes))) //Sort largest to smol
	ans_part2 := basinSizes[0] * basinSizes[1] * basinSizes[2]

	fmt.Printf("The answer to Part 2 is %v, apparently\n", ans_part2)

}
