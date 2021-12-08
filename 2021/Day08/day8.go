package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var path = "input.txt" //path to problem input
const PART int = 1
const PIPE int = 10

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

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_str [][]string
	var answer int = 0

	//Go through the instructions and make them usable
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		input_str = append(input_str, tmp)
	}
	// fmt.Println(input_str[1][10])

	for _, line := range input_str {
		for i := PIPE + 1; i < len(line); i++ {

			switch len(line[i]) {
			case 2, 4, 3, 7:
				// fmt.Println(line[i])
				answer++
			default:
				continue
			}
		}
	}
	fmt.Println(answer)
}
