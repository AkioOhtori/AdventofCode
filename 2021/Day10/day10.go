package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var path = "input.txt" //path to problem input

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
	//closing characters and their corresponding openers
	var match = map[string]string{")": "(", "]": "[", ">": "<", "}": "{"}
	//Store the points values for each part
	var points_pt1 = map[string]int{")": 3, "]": 57, ">": 25137, "}": 1197}
	var points_pt2 = map[string]int{"(": 1, "[": 2, "<": 4, "{": 3}
	var scores_pt2 []int

	//Go through the lines and convert them to slices of strings
	for scanner.Scan() {
		input_str := strings.Split(scanner.Text(), "")
		var opens []string
		var bad bool = false
		// var closes []string
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
					bad = true
					break out
					// otherwise we're OK and just need to remove the corresponding opener
				} else {
					opens = append(opens[1:]) //remove the first element
				}
			}
		}

		//If we completed with no errors, we need to calculate the score for that line
		var score_pt2 int = 0
		if !bad { //then we're done but incomplete
			bad = false //reset the "bad" flag (Probably a better way to do this)
			//There is no reason to calculate the corresponding closers, so we don't
			for _, i := range opens {
				score_pt2 = score_pt2*5 + points_pt2[i]
			}
			scores_pt2 = append(scores_pt2, score_pt2)
			score_pt2 = 0
		}
	}

	// Calculate Part 1 Scores
	var score_pt1 int = 0
	for _, i := range illegal {
		score_pt1 += points_pt1[i]
	}
	fmt.Printf("The Answer to Part 1 is %v\n", score_pt1)

	//Calculate Part 2 Scores
	sort.Ints(scores_pt2)                                 //sort the scores
	score_pt2 := scores_pt2[len(scores_pt2)/2]            // find the middle
	fmt.Printf("The Answer to Part 2 is %v\n", score_pt2) //print
}
