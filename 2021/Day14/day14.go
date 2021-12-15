package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var path = "input.txt" //path to problem input
const PART int = 2

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
		scanhelper++ //ugh could probably do this another way
	}

	var a = make(map[string]int)
	var pairs = make(map[string]int)
	var children = make(map[string][]string)

	for _, i := range template {
		a[string(i)] += 1
	}
	if PART == 1 {
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
	} else {
		//make map of pairs
		for i := 1; i < len(template); i++ {
			pair := (template[i-1 : i+1])
			pairs[pair] += 1
		}
		fmt.Println(pairs)
		fmt.Println(a)
		fmt.Println()
		for p, new := range ins_map {
			//p = pair
			//new = inserted letter
			//want = resulting pairs (ex: AB -> C; therefor AB = AC CB)
			children[p] = []string{p[:1] + new, new + p[1:]}
		}
		// process maps of pairs
		for s := 0; s < 40; s++ {
			pairs_new := make(map[string]int)
			for p, num := range pairs {
				//need: for every element in pairs
				//0) score resulting letter
				//1) determine resulting pairs
				//2) add to those resulting pairs
				a[ins_map[p]] += num
				pairs_new[(children[p][0])] += num
				pairs_new[(children[p][1])] += num
			}
			//save new set for the next round
			pairs = pairs_new
		}
	}

	//finally, find answers
	most := 0
	least := 999999999999999999
	for _, i := range a {
		if i > most {
			most = i
		}
		if i < least {
			least = i
		}
	}
	p := message.NewPrinter(language.English)
	p.Printf("B=%d\tC=%d\nH=%d\tN=%d\n", a["B"], a["C"], a["H"], a["N"])
	fmt.Printf("The answer to Part %v is: %v\n", PART, most-least)

}
