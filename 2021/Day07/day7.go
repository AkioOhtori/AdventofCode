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
	var input_int []int
	var input_flt []float64
	var maxx int
	// var minn int = 9999

	//Go through the instructions and make them usable
	for scanner.Scan() {
		input_str = strings.Split(scanner.Text(), ",")
	}
	for _, x := range input_str {
		tmp, _ := strconv.Atoi(x)
		input_int = append(input_int, tmp)
		input_flt = append(input_flt, float64(tmp))
		if tmp > maxx {
			maxx = tmp
		}
	}
	var fuel float64 = 0
	median, _ := stats.Median(input_flt)
	fmt.Println(math.Abs(-1))
	for _, x := range input_flt {
		fuel += math.Abs(median - x)
	}
	fmt.Println(len(input_int), maxx, median, fuel)
}
