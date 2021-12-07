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

var path = "input.txt" //path to problem input
const PART int = 2
const X1 int = 0
const Y1 int = 1
const X2 int = 2
const Y2 int = 3

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

func prettyPrintMatrix2D(arr [][]int) {
	y_print := make([]int, len(arr[0]))
	for x := 0; x < len(arr[0]); x++ {
		y_print[x] = x
	}
	fmt.Printf("X\t%v\n", y_print)
	for i, x := range arr {
		fmt.Printf("%v\t%v\n", i, x)
	}
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
	var dir [500][]int
	max_x := 0
	max_y := 0
	x := 0

	//Go through the instructions and make them usable
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		directions := append(strings.Split(line[0], ","), strings.Split(line[2], ",")...)
		dir[x] = convTxtSlicetoInt(directions)

		//If it goes the wrong direction, flip it
		//Doing Y first means X1 will ALWAYS be greater than Y1 and we need that
		if dir[x][Y1] > dir[x][Y2] {
			temp := []int{dir[x][X2], dir[x][Y2], dir[x][X1], dir[x][Y1]}
			copy(dir[x], temp)
		}
		if dir[x][X1] > dir[x][X2] {
			temp := []int{dir[x][X2], dir[x][Y2], dir[x][X1], dir[x][Y1]}
			copy(dir[x], temp)
		}

		//used for determining the size of the matrix
		if dir[x][X1] > max_x {
			max_x = dir[x][X1]
		}
		if dir[x][X2] > max_x {
			max_x = dir[x][X2]
		}
		if dir[x][Y1] > max_y {
			max_y = dir[x][Y1]
		}
		if dir[x][Y2] > max_y {
			max_y = dir[x][Y2]
		}
		x++
		//this doesn't actually do anything?

	}

	fmt.Printf("max x = %v and max y = %v\n", max_x, max_y)

	overlap := 0 //final answer

	//create empty field of correct size
	fields := make([][]int, max_y+1)
	for i := 0; i < len(fields); i++ {
		fields[i] = make([]int, max_x+1)
	}
	fmt.Printf("Fields size is x = %v and y = %v\n", len(fields[0]), len(fields))

	// Go through the instructions and mask off what they say
	for _, xyxy := range dir {
		if len(xyxy) == 0 {
			break
		}
		// Linear X Progression
		if xyxy[Y1] == xyxy[Y2] {
			for x := xyxy[X1]; x <= xyxy[X2]; x++ {
				fields[xyxy[Y1]][x]++
				if fields[xyxy[Y1]][x] == 2 {
					overlap++
				}
			}
			// Linear Y progression
		} else if xyxy[X1] == xyxy[X2] {
			for y := xyxy[Y1]; y <= xyxy[Y2]; y++ {
				fields[y][xyxy[X1]]++
				if fields[y][xyxy[X1]] == 2 {
					overlap++
				}
			}
			//Part 2 ONLY diag progression
		} else if PART == 2 {
			s := 1 //sign for decreasing Y
			if xyxy[Y1] > xyxy[Y2] {
				s = -1
			}
			// fmt.Printf("Doing a diag for %v with s=%v\n", xyxy, s)
			for x := 0; x+xyxy[X1] <= xyxy[X2]; x++ {
				xx := xyxy[X1] + x
				yy := (x * s)
				yy = xyxy[Y1] + yy

				fields[yy][xx]++
				if fields[yy][xx] == 2 {
					overlap++
				}
			}

		}
	}
	if path == "sample.txt" {
		prettyPrintMatrix2D(fields)
	}

	fmt.Printf("I think we have it and the overlap is %v!\n", overlap)
}
