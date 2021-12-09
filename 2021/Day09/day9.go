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

func fuckingFuck(i [][]int, x int, y int, a [][2]int) [][2]int {
	val := i[y][x]
	// fmt.Printf("Checking %v,%v = %v\n", x, y, i[y][x])
	if val == 9 {
		return a
	}

	//Do xr (x + 1) first
	xr := x + 1
	if xr < len(i[y]) { //make sure it is in bounds, if not, skip
		if i[y][xr] != 9 {
			// fmt.Printf("I'm on xr and looking at %v,%v = %v\n", y, xr, i[y][xr])
			if checkAnswer(a, [2]int{y, xr}) {
				a = append(a, [2]int{y, xr})
				a = fuckingFuck(i, xr, y, a)
				// fmt.Printf("Adding %v,%v = %v\n", xr, y, i[y][x])
			}
		}
	}
	yd := y + 1
	if yd < len(i) {
		if i[yd][x] != 9 {
			// fmt.Printf("I'm on yd and looking at %v,%v = %v\n", yd, x, i[yd][x])
			if checkAnswer(a, [2]int{yd, x}) {
				a = append(a, [2]int{yd, x})
				a = fuckingFuck(i, x, yd, a)
				// fmt.Printf("Adding %v,%v = %v\n", x, yd, i[y][x])
			}
		}
	}
	xl := x - 1
	if xl >= 0 {
		if i[y][xl] != 9 {
			// fmt.Printf("I'm on xl and looking at %v,%v = %v\n", y, xl, i[y][xl])
			if checkAnswer(a, [2]int{y, xl}) {
				a = append(a, [2]int{y, xl})
				a = fuckingFuck(i, xl, y, a)
				// fmt.Printf("Adding %v,%v = %v\n", xl, y, i[y][x])
			}
		}
	}
	yu := y - 1
	if yu >= 0 {
		if i[yu][x] != 9 {
			// fmt.Printf("I'm on yu and looking at %v,%v = %v\n", yu, x, i[yu][x])
			if checkAnswer(a, [2]int{yu, x}) {
				a = append(a, [2]int{yu, x})
				a = fuckingFuck(i, x, yu, a)
				// fmt.Printf("Adding %v,%v = %v\n", x, yu, i[y][x])
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
	// fmt.Printf("Adding %v\n", new)
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
out:
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
					basin = append(basin, [2]int{y, x})
					// fmt.Printf("Found and added y=%v,x=%v = %v\n", y, x, input_int[y][x])
					basin = fuckingFuck(input_int, x, y, basin)
					basinSizes = append(basinSizes, len(basin))
					basin = [][2]int{}
					if false {
						break out
					}
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
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	ans_part2 := basinSizes[0] * basinSizes[1] * basinSizes[2]

	fmt.Printf("The answer to Part 1 is %v danger units\n", answer)
	fmt.Printf("The answer to Part 2 is %v, apparently\n", ans_part2)

}
