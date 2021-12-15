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
	var ins_map = make(map[string]string)

	for scanner.Scan() {
		temp_str := scanner.Text()
		if scanhelper == 0 {
			template = temp_str
		} else if scanhelper == 1 {

		} else {
			tmp := strings.Split(temp_str, " -> ")
			ins_map[tmp[0]] = tmp[1]
		}
		scanhelper++
	}

	var a = make(map[string]int)
	for _, i := range template {
		a[string(i)] += 1
	}

	for s := 0; s < 10; s++ {
		new_template := ""
		for i := 1; i < len(template); i++ {
			pair := (template[i-1 : i+1])

			new_template += pair[:1] + ins_map[pair]
			a[ins_map[pair]] += 1

		}
		new_template += template[len(template)-1:]
		template = new_template
	}

	//finally, find answers
	most := 0
	least := 9999999
	for _, i := range a {
		if i > most {
			most = i
		}
		if i < least {
			least = i
		}
	}
	fmt.Printf("The answer to Part 1 is: %v\n", most-least)

}
