package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var h int = 0 //horizontal
	var d int = 0 //depth
	var a int = 0 //aim
	/* PART ONE CODE
	for scanner.Scan() {
		txt := (scanner.Text())
		dir := strings.Fields(txt)
		val, _ := strconv.Atoi(dir[1])

		if dir[0] == "up" {
			d -= val
		} else if dir[0] == "down" {
			d += val
		} else if dir[0] == "forward" {
			h += val
		} else {
			h -= val
		}
	}
	*/ //END PART 1 CODE

	// PART 2 CODE
	for scanner.Scan() {
		txt := (scanner.Text())
		dir := strings.Fields(txt)
		val, _ := strconv.Atoi(dir[1])

		if dir[0] == "up" {
			a -= val //up X decreases your aim by X units
		} else if dir[0] == "down" {
			a += val //down X increases your aim by X units
		} else if dir[0] == "forward" {
			h += val       //increases your horizontal position by X units
			d += (val * a) //increases your depth by your aim multiplied by X
		} else {
			fmt.Println("Poop")
		}
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the depth is %v and the horizontal %v.\n", d, h)
	fmt.Printf("The answer is %v!", h*d)

}
