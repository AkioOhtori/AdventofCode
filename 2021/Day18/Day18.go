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

func checkNew(o []string, z int, a int) ([]string, int) {
	var n int = 1
	b, _ := strconv.Atoi(o[z])
	c, err := strconv.Atoi(o[z+1])
	if err != nil {
		b = b*10 + c
		n = 2
		fmt.Printf("Found a double newb at z=%v and made it %v", z, b)
	}
	var new []string
	if (a + b) > 9 {
		temp := strconv.Itoa((a + b))
		new = []string{temp[:1], temp[1:]}
	} else {
		new = []string{strconv.Itoa((a + b))}
	}
	return new, n

}

func explode(o []string, open int) []string {
	x, y, end := findNumbers(o, open)
	fmt.Println("Exploding ", x, y, " at position ", open)
	fmt.Println(o)

	var left []string
	var right []string

	for z := open; z >= 0; z-- { //NOT TESTED
		_, err := strconv.Atoi(o[z])
		if isError(err) {
			continue
		} else {
			var new []string
			// if (l + x) > 9 {
			// 	temp := strconv.Itoa((l + x))
			// 	new = []string{temp[:1], temp[1:]}
			// } else {
			// 	new = []string{strconv.Itoa((l + x))}
			// }
			var n int
			new, n = checkNew(o, z, x)
			left = append(left, o[:z]...)
			left = append(left, new...)
			left = append(left, o[z+n:open]...)
			fmt.Println("left = ", left)
			break
		}
	}
	if len(left) == 0 { //TESTED OK
		//no number found so...
		left = append(left, o[:open]...)
	}

	//RIGHT!!
	for z := end; z < len(o); z++ { //TESTED OK
		_, err := strconv.Atoi(o[z])
		if isError(err) {
			continue
		} else {

			var new []string
			// if (r + y) > 9 {
			// 	temp := strconv.Itoa((r + y))
			// 	new = []string{temp[:1], temp[1:]}
			// } else {
			// 	new = []string{strconv.Itoa((r + y))}
			// }
			var n int
			new, n = checkNew(o, z, y)
			right = append(right, "0")
			// right = append(right, ",")
			right = append(right, o[end+1:z]...)
			right = append(right, new...)
			right = append(right, o[z+n:]...)
			fmt.Printf("right = %v\n", right)
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
	fmt.Println("These string aren't too big to fail!  SPLIT!", o[i:i+2])
	s := o[i] + o[i+1]
	num, _ := strconv.Atoi(s)
	var num_l []string
	var num_r []string
	if num >= 20 {
		t := strconv.Itoa(int(math.Floor(float64(num) / 2.0)))
		num_l = append(num_l, t[:1])
		num_l = append(num_l, t[1:])
		t = strconv.Itoa(int(math.Ceil(float64(num) / 2.0)))
		num_r = append(num_r, t[:1])
		num_r = append(num_r, t[1:])

	} else {
		t := strconv.Itoa(int(math.Floor(float64(num) / 2.0)))
		num_l = append(num_l, t)
		t = strconv.Itoa(int(math.Ceil(float64(num) / 2.0)))
		num_r = append(num_r, t)

	}
	var new []string
	new = append(new, o[:i]...)
	// new = append(new, []string{"[", strconv.Itoa(num_l), ",", strconv.Itoa(num_r), "]"}...)
	new = append(new, "[")
	new = append(new, num_l...)
	new = append(new, ",")
	new = append(new, num_r...)
	new = append(new, "]")
	new = append(new, o[i+2:]...)
	fmt.Println("Finished splitting and returned: ", new)
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
		var commas int
		var resolved bool = false
	outout:
		for boobs := 0; boobs < 2000; boobs++ {
			fmt.Println("hello")
		out:
			for i, e := range combined {
				// fmt.Println(i, "\t", e)

				switch e {
				case "[":
					opens++
					if opens >= 5 {
						fmt.Println("Exploding at position ", i)
						temp := explode(combined, i)
						// fmt.Println(temp)
						combined = temp
						copy(combined, temp)
						resolved = true
						break out
					}

				case ",":
					commas++
				case "]":
					opens--
				case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
					// fmt.Printf("We're at i=%v which is %v and checking for split\n", i, e)
					_, err := strconv.Atoi(combined[i+1])
					if isError(err) {
						continue
					} else {
						fmt.Println("Doing a split()!", e, combined)
						temp := split(combined, i)
						combined = temp
						copy(combined, temp)
						resolved = true
						break out
					}
				default: //number greater than 10... I hope!?
					//split

				} //switch
			} //inner for loop
			fmt.Println("Broke out and got ", combined)
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
	}
}
