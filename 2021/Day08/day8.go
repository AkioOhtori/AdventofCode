package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
const PART int = 1
const PIPE int = 10

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

//Ingests a string and outputs it sorted
func sortString(str string) string {
	var ordered string

	temp_slice := strings.Split(str, "")
	sort.Strings(temp_slice)

	//do I recombine or leave as an array?  Recombine!
	ordered = strings.Join(temp_slice, "")
	return ordered
}

//This function takes elements of "sub" out of "from" and returns the difference
func subtractSegments(from string, sub string) (ret string) {
	for _, x := range strings.Split(sub, "") {
		from = strings.ReplaceAll(from, x, "")
	}
	return from
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
	var input_str [][]string
	var answer int = 0

	//Go through the instructions and make them usable
	for scanner.Scan() {
		input_str = append(input_str, strings.Fields(scanner.Text()))
	}

	//Iterate over all inputs and check for lengths of 2,3,4,and7 (1478)
	for _, line := range input_str {
		for i := PIPE + 1; i < len(line); i++ {

			switch len(line[i]) {
			case 2, 4, 3, 7:
				answer++
			default:
				continue
			}
		}
	}
	//Thats it!  That was all of part 1
	fmt.Printf("The answer to Part 1 is: %v\n", answer)
	answer = 0

	/* Part 2 Notes
	*Can ID a segment via 7-1
	*Can ID 3 as it is the only qty5 that 5-1 = len3
	*Can ID 9 as it is the only qty6 with all but 1 segment from 3
	*Can ID b from above
	*Can ID d from 4 - 1 - b
	*Can ID g from 9 - 4 - a
	*Can ID e from whatever 9 is missing
	*Can ID 0 because it doesn't have d
	*Can ID 6 by process of elimination (6char)
	*Can ID c from not in 6
	*Can ID f from not 1 - c
	e and f CAN be id'd based on unique occurance count (optional)

	Which should give us everything!  But how TF to program that?!
	1) Pull line which is a slice 0-9 = segs, 11-14 = Answers
	2) Organize segs alphabetically and by length
	3) Determine 1, 4, 7, and 8 and add to translation dict
	4) ???
	5) Churn out answer


	/$$   /$$  /$$$$$$  /$$$$$$$$ /$$$$$$$$ /$$
	| $$$ | $$ /$$__  $$|__  $$__/| $$_____/| $$
	| $$$$| $$| $$  \ $$   | $$   | $$      | $$
	| $$ $$ $$| $$  | $$   | $$   | $$$$$   | $$
	| $$  $$$$| $$  | $$   | $$   | $$__/   |__/
	| $$\  $$$| $$  | $$   | $$   | $$
	| $$ \  $$|  $$$$$$/   | $$   | $$$$$$$$ /$$
	|__/  \__/ \______/    |__/   |________/|__/

	My solution to Part 2 is gross and I hate it.
	I'm like... 80% sure this is the absolute WRONG way to solve this
	but
	alas
	It is the solution that I have.
	*/

	//Iterate over the entire input file
	for _, line := range input_str {
		//Create translation maps
		var segments = map[string]string{"a": "", "b": "", "c": "", "d": "", "e": "", "f": "", "g": ""}
		var numbers = map[int]string{0: "", 1: "", 2: "", 3: "", 4: "", 5: "", 6: "", 7: "", 8: "", 9: ""}

		//Sort input and output alphabetically
		var input []string
		for i := 0; i < PIPE; i++ {
			input = append(input, sortString(line[i]))
		}
		var output []string
		for i := PIPE + 1; i < len(line); i++ {
			output = append(output, sortString(line[i]))
		}

		//Sort input by length (stolen code)
		sort.Slice(input, func(i, j int) bool {
			return len(input[i]) < len(input[j])
		})

		//Store numbers of unique length (known)
		numbers[1] = input[0]
		numbers[7] = input[1]
		numbers[4] = input[2]
		numbers[8] = input[9]

		//Find a segment by subtracting 1 from 7
		segments["a"] = subtractSegments(numbers[7], numbers[1])

		//len=5 segments (2,3,5) are 3,4,5 in input when sorted
		//this finds 3 based on it being the only len5 that if you remove "1" you get len3
		for i := 3; i <= 5; i++ {
			if len(subtractSegments(input[i], numbers[1])) == 3 {
				numbers[3] = input[i]
				break
			}
		}

		//len=6 segments (0,6,9) are 6,7,8 in input when sorted
		//this finds 9 as it is the only len6 with with 9-3=len1
		//This is fucking dirty and gross but it works so heh
		//we do get b for free as it is the excluded segement
		for i := 6; i <= 8; i++ {
			x := subtractSegments(input[i], numbers[3])
			if len(x) == 1 {
				numbers[9] = input[i]
				segments["b"] = x
			}
		}

		//we can get d by removing b and 1 from 4
		segments["d"] = (subtractSegments(numbers[4], numbers[1]))
		segments["d"] = (subtractSegments(segments["d"], segments["b"]))

		//we can get g from 9 - 4 - a (same as above, just all one line)
		segments["g"] = subtractSegments(subtractSegments(numbers[9], numbers[4]), segments["a"])

		//Can ID e from whatever 9 is missing
		for x := range segments {
			if !strings.Contains(numbers[9], x) {
				segments["e"] = x
			}
		}

		// Can ID 0 because it is the only one to not have d
		for i := 6; i <= 8; i++ {
			if !strings.Contains(input[i], segments["d"]) {
				numbers[0] = input[i]
			}
		}

		//Can ID 6 by process of elimination (only unknown 6char)
		for i := 6; i <= 8; i++ {
			if input[i] != numbers[0] && input[i] != numbers[9] {
				numbers[6] = input[i]
			}
		}

		//Can ID c because it isn't in 6
		for x := range segments {
			if !strings.Contains(numbers[6], x) {
				segments["c"] = x
			}
		}

		//Can ID f from 1-c
		segments["f"] = (subtractSegments(numbers[1], segments["c"]))

		//Now we have all the segements translated, we can just MAKE our last two numbers
		numbers[2] = sortString(segments["a"] + segments["c"] + segments["d"] + segments["e"] + segments["g"])
		numbers[5] = sortString(segments["a"] + segments["b"] + segments["d"] + segments["f"] + segments["g"])

		//Translating would sure be eaiser if we could just plug and play... so lets!
		oppo := map[string]int{numbers[0]: 0, numbers[1]: 1, numbers[2]: 2, numbers[3]: 3, numbers[4]: 4, numbers[5]: 5, numbers[6]: 6, numbers[7]: 7, numbers[8]: 8, numbers[9]: 9}

		//Lastly, translate the output and add it to the answer
		var temps string
		var tempi int
		for _, x := range output {
			temps += strconv.Itoa(oppo[x])
		}
		tempi, _ = strconv.Atoi(temps)
		answer += tempi
	}
	//Output the answer
	fmt.Printf("The answer to part 2 is probably %v\n", answer)

}
