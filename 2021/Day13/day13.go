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

func prettyPrintMatrix2D(m [][]int) {
	fmt.Println()
	for _, x := range m {
		fmt.Println(x)
	}
	fmt.Println()
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
	var fold_instructions [][]string
	// var paper [][]int
	var max_x int
	var max_y int

	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	for scanner.Scan() {
		temp_str := scanner.Text()
		if len(temp_str) > 0 {
			if temp_str[:1] != "f" {
				input_str := strings.Split(temp_str, ",")
				// fmt.Println(input_str)
				var temp_int []int
				for _, x := range input_str {
					y, _ := strconv.Atoi(x)
					temp_int = append(temp_int, y)
				}
				input_int = append(input_int, temp_int)

				if temp_int[0] > max_x {
					max_x = temp_int[0]
				}
				if temp_int[1] > max_y {
					max_y = temp_int[1]
				}

			} else {
				t := strings.Split(temp_str, "=")
				fold_instructions = append(fold_instructions, []string{t[0][11:], t[1]})
			}
		}

	}
	//create empty field of correct size
	paper := make([][]int, max_y+1)
	for i := 0; i < len(paper); i++ {
		paper[i] = make([]int, max_x+1)
	}
	for _, dot := range input_int {
		paper[dot[1]][dot[0]] = 1
	}
	count := 0
	for _, y := range paper {
		for _, x := range y {
			count += x
		}
	}

	fmt.Printf("The matrix is x = %v, y = %v with %v dots\n", max_x+1, max_y+1, count)
	// prettyPrintMatrix2D(paper)

	for i := 0; i < 1; i++ {
		fmt.Println(fold_instructions[i])
		if fold_instructions[i][0] == "y" {
			along, _ := strconv.Atoi(fold_instructions[i][1])
			new := make([][]int, along)

			for y := 0; y < along; y++ {
				new[y] = make([]int, len(paper[y]))
				copy(new[y], paper[y])

				for x := 0; x < len(paper[y]); x++ {
					new[y][x] = (new[y][x] | paper[len(paper)-1-y][x])
				}
			}
			paper = new

		} else {
			along, _ := strconv.Atoi(fold_instructions[i][1])
			new := make([][]int, len(paper))
			for y := 0; y < len(paper); y++ {
				new[y] = make([]int, along)
				copy(new[y], paper[y][:along])

				for x := 0; x < along; x++ {
					new[y][x] = (new[y][x] | paper[y][len(paper[y])-1-x])
					count += new[y][x]
				}
				// break
			}

			paper = new
		}

	}
	fmt.Printf("The new, folded matrix is x=%v, y=%v\n", len(paper[0]), len(paper))
	var answer_pt1 int = 0
	// prettyPrintMatrix2D(paper)
	//Calculate the answer
	for _, y := range paper {
		for _, x := range y {
			answer_pt1 += x
		}
	}
	fmt.Println(answer_pt1)

} //EOF
