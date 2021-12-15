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
var distances [][]int

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

//Bastardized Dijkstra's Algorithm
func mapCave(cave [][]int, y int, x int, danger int) {
	//make an array of all the possible directions
	d := [4][2]int{{y + 1, x}, {y - 1, x}, {y, x + 1}, {y, x - 1}}
	//check to make sure we aren't done
	if x == endx && y == endy {
		answers = append(answers, danger)
		return
	}
	//iterate over the directions
	for _, yx := range d {
		//check direction is valid
		if yx[0] > endy || yx[1] > endx || yx[0] < 0 || yx[1] < 0 {
			continue
		} else {
			next := cave[yx[0]][yx[1]]
			//if the "distance" (danger) is lower than we've done, go for it
			if danger+next < distances[yx[0]][yx[1]] {
				distances[yx[0]][yx[1]] = danger + next
				mapCave(cave, yx[0], yx[1], danger+next)
			}
		}
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

	endx = len(input[0]) - 1 //ending x coords
	endy = len(input) - 1    //ending y coords

	//prefill the "distances" matrix with "infinity"
	for y := 0; y <= endy; y++ {
		var t []int
		for x := 0; x <= endx; x++ {
			t = append(t, math.MaxInt64)
		}
		distances = append(distances, t)
	}

	//Make it so!
	mapCave(input, 0, 0, 0)
	sort.Ints(answers)
	fmt.Printf("The Answer to Part 1 is %v \n", answers[0])
}
