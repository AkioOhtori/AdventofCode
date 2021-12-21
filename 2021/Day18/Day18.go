package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var path = "sample2.txt" //path to problem input
const PART int = 1

// Function to handle errors
func isError(err error) bool {
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	return (err != nil)
}

func findNumbers(input []string, open int) (int, int, int) {
	var comma int = 0
	var end_r int = 0
	// for i, c := range input {
	for i := open; i < len(input); i++ {
		// fmt.Printf("%v\t%v\t%T\n", i, input[i], input[i])
		if input[i] == (",") {
			comma = i
		} else if input[i] == "]" {
			end_r = i
			break
		}
	}
	x_str := input[open+1 : comma]
	y_str := input[comma+1 : end_r]

	// fmt.Println("Found X and Y!", x_str, y_str)

	var xx string
	var yy string
	for _, c := range x_str {
		xx += c
	}
	for _, c := range y_str {
		yy += c
	}
	x_out, _ := strconv.Atoi(xx)
	y_out, _ := strconv.Atoi(yy)

	return x_out, y_out, end_r
}

func explode(o []string, open int) []string {
	x, y, end := findNumbers(o, open)
	fmt.Println("Exploding ", x, y, " at position ", open)
	fmt.Println("\t\t\t", o)

	var left []string
	var right []string

	for z := open; z >= 0; z-- { //NOT TESTED
		l, err := strconv.Atoi(o[z])
		if isError(err) {
			continue
		} else {
			left = append(left, o[:z]...)
			left = append(left, strconv.Itoa((l + x)))
			left = append(left, o[z+1:open]...)
			break
		}
	}
	if len(left) == 0 { //TESTED OK
		//no number found so...
		left = append(left, o[:open]...)
	}

	//RIGHT!!
	for z := end; z < len(o); z++ { //TESTED OK
		r, err := strconv.Atoi(o[z])
		if isError(err) {
			continue
		} else {
			right = append(right, "0")
			right = append(right, o[end+1:z]...)
			right = append(right, strconv.Itoa((r + y)))
			right = append(right, o[z+1:]...)
			break
		}
	}
	if len(right) == 0 { //NOT TESTED
		right = append([]string{"0"}, o[open+5:]...)
	}
	// fmt.Println("Outputting: ", append(left, right...))
	return (append(left, right...))

}

func split(o []string, i int) []string {
	// fmt.Println("These string aren't too big to fail!  SPLIT!", o[i:i+2])
	num, _ := strconv.Atoi(o[i])
	num_l := strconv.Itoa(int(math.Floor(float64(num) / 2.0)))
	num_r := strconv.Itoa(int(math.Ceil(float64(num) / 2.0)))
	var new []string
	new = append(new, o[:i]...)
	new = append(new, "[", num_l, ",", num_r, "]")
	// new = append(new, num_l)
	// new = append(new, ",")
	// new = append(new, num_r)
	// new = append(new, "]")
	new = append(new, o[i+1:]...)
	fmt.Println("Finished splitting:\t", new)
	return new
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

	var old []string
	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	for scanner.Scan() {
		var new []string
		var combined []string
		temp_str := scanner.Text()
		new = strings.Split(temp_str, "")
		fmt.Println("New line is ", new)
		if len(old) == 0 {
			old = new
			copy(old, new)
			continue
		} else {
			// copy(combined, old)
			combined = []string{"["}
			combined = append(combined, old...)
			combined = append(combined, ",")
			combined = append(combined, new...)
			combined = append(combined, "]")
			fmt.Println("New starting point is: ", combined)
			// break
		}

		//reduce

		/*
			To reduce a snailfish number, you must repeatedly do the first action in this list that applies to the snailfish number:

			If any pair is nested inside four pairs, the leftmost such pair explodes.
			If any regular number is 10 or greater, the leftmost such regular number splits.
		*/

		var opens int
		// var commas int
		var resolved bool = false
	outout:
		for { //} boobs := 0; boobs < 20000; boobs++ {
			fmt.Println("hello")
		out:
			for i, e := range combined {
				// fmt.Println(i, "\t", e)

				switch e {
				case "[":
					opens++
					if opens >= 5 {
						// fmt.Println("Exploding at position ", i)
						temp := explode(combined, i)
						// fmt.Println(temp)
						combined = temp
						copy(combined, temp)
						resolved = true
						break out
					}

				case ",":
					// commas++
				case "]":
					opens--
				case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
					// fmt.Printf("We're at i=%v which is %v and checking for split\n", i, e)
				default: //number greater than 10... I hope!?
					fmt.Println("Doing a split!", e, "\t", combined)
					temp := split(combined, i)
					combined = temp
					copy(combined, temp)
					resolved = true
					break out
				} //switch
			} //inner for loop
			fmt.Println("Broke out and got \t", combined)
			fmt.Println()
			opens = 0
			if resolved {
				resolved = false
			} else {
				fmt.Println("Finished?")
				break outout
			}

		} //outer for loop
		old = combined
		copy(old, combined)
		fmt.Printf("Fuuuuuuuuuuuuuuuully reduced this line and moving to the next!!\n")
		break
	}
}
