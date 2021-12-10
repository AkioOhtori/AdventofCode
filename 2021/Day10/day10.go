package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	//Open input file
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	//Load all lines in file into a slice of slices
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var illegal []string
	var match = map[string]string{")": "(", "]": "[", ">": "<", "}": "{"}
	var scores = map[string]int{")": 3, "]": 57, ">": 25137, "}": 1197}

	//Go through the lines and convert them to slices of strings
	for scanner.Scan() {
		input_str := strings.Split(scanner.Text(), "")
		var opens []string
	out:
		//evaluate each character in the string
		for _, x := range input_str {
			switch x {
			//if it is an opener, add it to the lis
			case "(", "[", "<", "{":
				opens = append([]string{x}, opens...)
				//if it is a closer, check if it is correct
			case ")", "]", ">", "}":
				//if the closer is NOT the opposite of the most recent open, it is bad
				if opens[0] != match[x] {
					illegal = append(illegal, x)
					break out
					// otherwise we're OK and just need to remove the corresponding opener
				} else {
					opens = append(opens[1:]) //remove the first element
				}
			default:
				fmt.Println("fuck")
			}
		}
	}
	var score_pt1 int = 0
	for _, i := range illegal {
		score_pt1 += scores[i]
	}
	fmt.Printf("The Answer to Part 1 is %v\n", score_pt1)
}
