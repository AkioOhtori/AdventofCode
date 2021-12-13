package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var path = "input.txt" //path to problem input
const PART int = 2
const ORIGIN int = 0
const DESTIN int = 1

var answer [][]string

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func checkLower(p string) bool {
	for _, x := range p {
		if unicode.IsLower(x) {
			return true
		} else {
			return false
		}
	}
	return false //default
}

func checkFor(p []string, new string) bool {
	for _, y := range p {
		if y == new {
			// fmt.Printf("Found %v in p = %v\n", new, p)
			return true
		}
	}
	// fmt.Printf("Did not find %v in p = %v\n", new, p)
	return false
}

func traverseTree(tree [][]string, p []string, begin string) {
	p = append(p, begin)

	for _, branch := range tree {
		if branch[ORIGIN] == begin {
			// fmt.Printf("Found a path from %v to %v in %v\n", begin, branch[DESTIN], branch)
			if branch[DESTIN] == "end" {
				// p = append(p, branch[DESTIN])
				answer = append(answer, p)
				continue
			} else if checkLower(branch[DESTIN]) {
				if checkFor(p, branch[DESTIN]) {
					continue
				}
			}
			// p = append(p, branch[DESTIN])
			// fmt.Printf("Started at %v and going to %v\n", begin, branch[DESTIN])
			new := make([]string, len(p))
			copy(new, p)
			traverseTree(tree, new, branch[DESTIN])
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
	var input [][]string
	var starts [][]string

	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	for scanner.Scan() {
		input_str := strings.Split(scanner.Text(), "-")
		// fmt.Println(input_str)
		if input_str[ORIGIN] == "start" {
			starts = append(starts, input_str)
		} else if input_str[DESTIN] == "start" {
			starts = append(starts, []string{"start", input_str[ORIGIN]})
		} else {
			input = append(input, input_str)
			//if not an endpoint, flip it and add it to the valid paths
			if input_str[DESTIN] != "end" {
				input = append(input, []string{input_str[1], input_str[0]})
			}
		}
	}
	var p []string
	for _, start := range starts {
		p = append(p, "start")
		traverseTree(input, p, start[DESTIN])
	}
	// fmt.Println(answer)
	fmt.Println(len(answer))
}
