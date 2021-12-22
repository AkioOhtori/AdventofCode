package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input

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

func prettyPrintMatrixLetters(m [][]int) {
	fmt.Println()
	for _, y := range m {
		for _, x := range y {
			if x == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func expand(image [][]int, i int) [][]int {
	var new_image [][]int
	var n int = 0

	new_image = append(new_image, make([]int, len(image[0])+6))
	new_image = append(new_image, make([]int, len(image[0])+6))
	new_image = append(new_image, make([]int, len(image[0])+6))
	for i := range image {
		var new_row []int
		new_row = append(new_row, n, n, n)
		new_row = append(new_row, image[i]...)
		new_row = append(new_row, n, n, n)
		new_image = append(new_image, new_row)
	}
	new_image = append(new_image, make([]int, len(image[0])+6))
	new_image = append(new_image, make([]int, len(image[0])+6))
	new_image = append(new_image, make([]int, len(image[0])+6))
	// prettyPrintMatrix2D(new_image)
	if n == 1 {
		for x := 0; x < 3; x++ {
			for z := 0; z < len(new_image[0]); z++ {
				new_image[x][z] = 1
				new_image[len(new_image)-x-1][z] = 1
			}
		}
	}
	return new_image
}

func fullExpand(image [][]int, e int) [][]int {
	var new_image [][]int
	for y := 0; y < e; y++ {
		new_image = append(new_image, make([]int, len(image[0])+e*6))
		new_image = append(new_image, make([]int, len(image[0])+e*6))
		new_image = append(new_image, make([]int, len(image[0])+e*6))
	}
	for i := range image {
		var new_row []int
		for y := 0; y < e; y++ {
			new_row = append(new_row, 0, 0, 0)
		}
		new_row = append(new_row, image[i]...)
		for y := 0; y < e; y++ {
			new_row = append(new_row, 0, 0, 0)
		}
		new_image = append(new_image, new_row)
	}
	for y := 0; y < e; y++ {
		new_image = append(new_image, make([]int, len(image[0])+e*6))
		new_image = append(new_image, make([]int, len(image[0])+e*6))
		new_image = append(new_image, make([]int, len(image[0])+e*6))
	}

	return new_image
}

func translate(image [][]int, yy int, xx int, algorithm []int) int {
	var b string
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			b += strconv.Itoa(image[y+yy][x+xx])
		}
	}
	poop, _ := strconv.ParseInt(b, 2, 64)
	// fmt.Println(xx, yy, b, poop, algorithm[poop])
	out := algorithm[poop]
	return out
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

	//load in only input line and convert it from hex to binary
	var helper int = 0
	var algorithm []int
	var image [][]int
	for scanner.Scan() {
		temp_str := scanner.Text()
		if helper == 0 {
			temp_sli := strings.Split(temp_str, "")
			for _, x := range temp_sli {
				if x == "." {
					algorithm = append(algorithm, 0)
				} else {
					algorithm = append(algorithm, 1)
				}
			}
		} else if helper == 1 {
			// continue
		} else {

			temp_sli := strings.Split(temp_str, "")
			// temp_row := make([]int, len(temp_sli))
			var temp_row []int
			for _, x := range temp_sli {
				if x == "." {
					temp_row = append(temp_row, 0)
				} else {
					temp_row = append(temp_row, 1)
				}

			}
			image = append(image, temp_row)

		}
		helper++
	}

	// Expand by 2 in EVERY direction
	// evaluate, I guess

	image = fullExpand(image, 50)
	prettyPrintMatrixLetters(image)

	for i := 0; i < 50; i++ {
		// temp := expand(image, i)
		// image = temp
		// copy(image, temp)

		var new [][]int
		for i := 0; i < len(image); i++ {
			new = append(new, make([]int, len(image[i])))
		}

		for y := 1; y < len(new)-1; y++ {
			for x := 1; x < len(new)-1; x++ {
				new[y][x] = translate(image, y, x, algorithm)
			}
		}
		// prettyPrintMatrixLetters(new)
		image = new
	}
	var answer_pt1 int
	for y := 0; y <= len(image)-70; y++ { //this is a bodge to deal with trash on the edges
		for _, x := range image[y] {
			answer_pt1 += x
		}
	}
	prettyPrintMatrixLetters(image)
	fmt.Println(answer_pt1)

}
