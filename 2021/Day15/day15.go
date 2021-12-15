package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
const PART int = 1

var endx int = 0
var endy int = 0
var answers = []int{math.MaxInt64}
var step_danger = [200]int{}

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func prettyPrintMatrix2D(m [][]int) {
	fmt.Println()
	for _, x := range m {
		fmt.Println(x)
	}
	fmt.Println()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func exploreCave(cave [][]int, y int, x int, danger int, step int) {
	// if x +1 >= len(cave[y])
	var right int = 9
	var down int = 9
	step++

	if danger >= answers[0] {
		return
	} else if danger > step_danger[step]+50 {
		return
	} else if danger < step_danger[step] {
		step_danger[step] = danger
	}

	if x == endx && y == endy {
		answers = append(answers, danger) //DOUBLE CHECK THIS
		sort.Ints(answers)
		fmt.Printf("%v\t", danger)
		return
	} else if x == endx {
		right = 999
		down = cave[y+1][x]
	} else if y == endy {
		down = 999
		right = cave[y][x+1]
	} else {
		down = cave[y+1][x]
		right = cave[y][x+1]
	}

	// fmt.Println(y, x, danger, down, right)

	if abs(right-down) < 3 {
		//explore both
		exploreCave(cave, y+1, x, danger+down, step)
		exploreCave(cave, y, x+1, danger+right, step)
	} else if right < down {
		//go right
		exploreCave(cave, y, x+1, danger+right, step)
	} else {
		//go down
		exploreCave(cave, y+1, x, danger+down, step)
	}

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

	var input [][]int

	//Go through the instructions and convert them to slices of slices, [y],[x]
	for scanner.Scan() {
		temp_str := strings.Split(scanner.Text(), "")
		temp_int := []int{}
		for _, x := range temp_str {
			i, _ := strconv.Atoi(x)
			temp_int = append(temp_int, i)
		}
		input = append(input, temp_int)
	}
	// prettyPrintMatrix2D(input)
	endx = len(input[0]) - 1
	endy = len(input) - 1
	// fmt.Println(endx, endy)

	for x := 0; x < len(step_danger); x++ {
		step_danger[x] = math.MaxInt64 - 1000
	}

	exploreCave(input, 0, 0, 0, -1)
	sort.Ints(answers)
	fmt.Println(answers[0])
}
