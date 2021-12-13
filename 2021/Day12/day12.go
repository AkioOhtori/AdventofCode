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
var lowers []string

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

func checkForDouble(p []string) bool {
	//This function checks to see if ANY double small cave exists
	for _, l := range lowers {
		var i int = 0
		for _, x := range p {
			if l == x {
				i++
				if i > 1 {
					return true
				}
			}
		}
	}
	return false
}

// This is the main recursive function that travels the tree starting at START+1
func traverseTree(tree [][]string, p []string, origin string) {

	p = append(p, origin) //Record where we came from

	//look through the tree to see if we can find a matching path
	for _, branch := range tree {
		if branch[ORIGIN] == origin { //match found!
			if branch[DESTIN] == "end" { //if it is an end, record the answer
				answer = append(answer, p)
				continue
			} else if checkLower(branch[DESTIN]) { //check if lower case
				if checkFor(p, branch[DESTIN]) { //check if already exists in answer
					if PART == 1 { //part 1 doesn't allow lowercase duplicates
						continue
					} else if checkForDouble(p) { //part 2 allows exactly 1 lower case duplicate
						//since we can only get here by having at least 1 of DESTIN
						//if ANY double also exists, we basically treat it like Part 1
						continue
					}
				}
			}

			//we're going another level deep, but want to keep this p separate, so make a new one
			new := make([]string, len(p))
			copy(new, p)

			//Going down another fork!
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
	var input [][]string  //input minus starting conditions
	var starts [][]string //starting conditions

	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	for scanner.Scan() {
		input_str := strings.Split(scanner.Text(), "-")

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
		//For part 2 we need to know all the "small" caves (lower case)
		if PART == 2 {
			if checkLower(input_str[0]) {
				lowers = append(lowers, input_str[0])
			}
			if checkLower(input_str[1]) {
				lowers = append(lowers, input_str[1])
			}
		}
	}

	//traverse the tree for all starting conditions
	for _, start := range starts {
		traverseTree(input, []string{}, start[DESTIN])
	}

	fmt.Printf("The answer to Part %v is %v \n", PART, len(answer))
}
