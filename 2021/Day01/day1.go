package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var path = "input1.txt" //path to problem input

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

	var txtlines []string

	// split lines into array/slice
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	defer file.Close()

	//initialize output variable
	var count int = 0

	/* Day 1 Part 1

	a, err := strconv.Atoi(txtlines[0])
	if isError(err) {
		return
	}
	//i++
	for i := 1; i < len(txtlines); i++ {
		b, err := strconv.Atoi(txtlines[i])
		if isError(err) {
			break
		}
		if b > a {
			count++
			fmt.Printf("%v is greater than %v! Count = %v\n", b, a, count)
		}
		a = b
	}
	*/

	//Day 1 Part 2

	//We start with the first three values, so need to get those queued
	a, _ := strconv.Atoi(txtlines[0])
	b, _ := strconv.Atoi(txtlines[1])
	c, _ := strconv.Atoi(txtlines[2])

	//Iterate over the remaining values
	for i := 3; i < len(txtlines); i++ {
		d, err := strconv.Atoi(txtlines[i])
		if isError(err) {
			break
		}
		windowa := a + b + c
		windowb := b + c + d

		//Check to see if the height has decreased (which is what we want)
		if windowb > windowa {
			count++
			fmt.Printf("%v is greater than %v! Count = %v\n", windowb, windowa, count)
		}

		//Slide the window for next iteration
		a = b
		b = c
		c = d
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", count)

}
