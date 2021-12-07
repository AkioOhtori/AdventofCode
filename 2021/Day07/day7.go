package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
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

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_str []string
	var input_flt []float64

	//Go through the instructions and make them usable
	for scanner.Scan() {
		input_str = strings.Split(scanner.Text(), ",")
	}
	for _, x := range input_str {
		tmp, _ := strconv.Atoi(x)
		input_flt = append(input_flt, float64(tmp))
	}

	// Add the fuel required to get to the median
	var fuel float64 = 0
	median, _ := stats.Median(input_flt)
	mean, _ := stats.Mean(input_flt)
	avg := math.Floor(mean)
	if PART == 1 {
		for _, x := range input_flt {
			fuel += math.Abs(median - x)
		}
	} else {
		for _, x := range input_flt {
			y := math.Abs(avg - x)
			fuel += (((y * y) + y) / 2) //triangle formula
		}
	}
	fmt.Printf("The data was %v long with a median of %v average of %v and total fuel used was %v\n", len(input_flt), median, mean, int(fuel))
}
