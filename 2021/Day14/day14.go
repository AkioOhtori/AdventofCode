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

	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	var scanhelper int = 0
	var template string
	var instructions [][]string

	for scanner.Scan() {
		temp_str := scanner.Text()
		if scanhelper == 0 {
			template = temp_str
		} else if scanhelper == 1 {

		} else {
			instructions = append(instructions, strings.Split(temp_str, " -> "))
		}
		scanhelper++
	}

	for s := 0; s < 10; s++ {
		new_template := ""
		for i := 1; i < len(template); i++ {
			pair := (template[i-1 : i+1])
			for _, x := range instructions {
				if x[0] == pair {
					new_template += pair[:1] + x[1]
				}
			}
		}
		new_template += template[len(template)-1:]
		template = new_template
	}

	//finally, find answers
	var a = make(map[string]int)
	for _, i := range template {
		a[string(i)] += 1
	}
	fmt.Println(a)
	most := 0
	least := 9999999
	for _, i := range a {
		fmt.Println(i)
		if i > most {
			most = i
		}
		if i < least {
			least = i
		}
	}
	fmt.Printf("The answer to Part 1 is: %v\n", most-least)

}
