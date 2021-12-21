package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
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
	// fmt.Println("Exploding ", x, y, " at position ", open)
	// fmt.Println("\t\t\t", o)

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
	new = append(new, o[i+1:]...)
	// fmt.Println("Finished splitting:\t", new)
	return new
}

func callSplit(combined []string) (done bool, output []string) {
	done = true
doom:
	for i, e := range combined {
		switch e {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", "[", "]": //all valid
			//don't care
		default: //any two didget number
			// fmt.Println("Doing a split!", e, "\t", combined)
			temp := split(combined, i)
			output = temp
			copy(output, temp)
			done = false
			break doom
		}
	}
	return
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
	var reduced []string
	var input [][]string
	var answer_pt2 int
	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	for scanner.Scan() {
		var new []string

		temp_str := scanner.Text()
		new = strings.Split(temp_str, "")
		input = append(input, new)
	}

	for x := range input {
		for y := x + 1; y < len(input); y++ {
			var combined []string

			combined = []string{"["}
			combined = append(combined, input[x]...)
			combined = append(combined, ",")
			combined = append(combined, input[y]...)
			combined = append(combined, "]")

			var opens int
			var resolved bool = false
		outout:
			for {
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
					case "]":
						opens--
					default:
						//don't care
					} //end switch
				} //end inner for
				// fmt.Println("Broke out and got \t", combined)
				// fmt.Println()
				opens = 0
				if resolved {
					resolved = false
				} else {
					// fmt.Println("Finished exploding?")
					done, temp := callSplit(combined)
					if !done {
						combined = temp
						copy(combined, temp)
					} else {
						reduced = combined
						copy(reduced, combined)
						break outout
					}
				}
			} //end outer for one
			old = combined
			copy(old, combined)
			// fmt.Printf("Fuuuuuuuuuuuuuuuully reduced this line and moving to the next!!\n")
			// break
			//end outer for

			//now that that nightmare is over, we have to score this...
			// done:
			for {
				var newnew []string
				var done bool = true
			done:
				for i, e := range reduced {
					switch e {
					case ",":
						if reduced[i+1] != "[" && reduced[i-1] != "]" {
							a, _ := strconv.Atoi(reduced[i-1])
							b, _ := strconv.Atoi(reduced[i+1])
							temp := a*3 + b*2
							temp_str := strconv.Itoa(temp)
							newnew = append(newnew, reduced[:(i-2)]...)
							newnew = append(newnew, temp_str)
							newnew = append(newnew, reduced[(i+3):]...)
							// fmt.Println(newnew)
							reduced = newnew
							copy(reduced, newnew)
							done = false
							break done

						}
					default:
						// newnew = append(newnew, e)
					}

				} //end of inner for
				if done {
					break
				}
				// fmt.Println(newnew)
				// break
			}
			fmt.Println(reduced)
			a, _ := strconv.Atoi(reduced[0])
			if a > answer_pt2 {
				answer_pt2 = a
			}
		}

	}
	fmt.Println(answer_pt2)
}
