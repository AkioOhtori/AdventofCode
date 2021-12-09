package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var path = "sample2.txt" //path to problem input
const PART int = 1
const PIPE int = 10

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func sortString(str string) string {
	var ordered string

	temp_slice := strings.Split(str, "")
	sort.Strings(temp_slice)
	// fmt.Println(temp_slice)

	//do I recombine or leave as an array?  Recombine!
	ordered = strings.Join(temp_slice, "")
	return ordered
}

func solveNumbers(in []string, subt_number string, start int, end int) (segment string, number string) {
	for i := start; i <= end; i++ {

		redu := subtractSegments(in[i], subt_number)

		if len(redu) == 1 {
			segment = redu
			number = in[i]
			return
		}
	}
	segment = ""
	number = ""
	return
}

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

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_str [][]string
	var answer int = 0

	//Go through the instructions and make them usable
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		input_str = append(input_str, tmp)
	}

	if PART == 1 {
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
		fmt.Printf("The anser to Part 1 is: %v\n", answer)
		answer = 0
	}
	/* Part 2 Notes
	*Can ID d segment via 7-1
	*Can ID 3 as it is the only qty5 with both 1 segments in it
	*Can ID 9 as it is the only qty6 with all but 1 segment from 3
	*Can ID e from above
	*Can ID f from 4 - 1 - e
	*Can ID c from 9 - 4 - d
	*Can ID g from whatever 9 is missing
	*Can ID 0 from not f
	*Can ID 6 by process of elimination (6char)
	*Can ID a from not in 6
	Can ID b from not a (or 9 - adef)
	Can ID c a bunch of ways but lets go with 4 + d - 9

	e and f CAN be id'd based on unique occurance count, which will lead to a

	Which should give us everything!  But how TF to program that?!
	1) Pull line which is a slice 0-9 = segs, 11-14 = Answers
	2) Organize segs alphabetically
	3) Determine 1, 4, 7, and 8 and add to translation dict
	4) Solve
	5) Churn out answer

	*/

	for _, line := range input_str {
		var segments = map[string]string{"a": "", "b": "", "c": "", "d": "", "e": "", "f": "", "g": ""}
		// var oppo = map[string]string{"a": "", "b": "", "c": "", "d": "", "e": "", "f": "", "g": ""}
		// var oppo [10]string
		var numbers = map[int]string{0: "", 1: "", 2: "", 3: "", 4: "", 5: "", 6: "", 7: "", 8: "", 9: ""}

		//Sort segs alphabetically
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

		// fmt.Println(input)

		//Store numbers of unique length
		numbers[1] = input[0]
		numbers[7] = input[1]
		numbers[4] = input[2]
		numbers[8] = input[9]

		//Find d segment by subtracting 1 from 7
		segments["d"] = subtractSegments(numbers[7], numbers[1])

		//len=5 segments (2,3,5) are 3,4,5 in input when sorted
		//this finds 3 based on it being the only len5 with "1" in it
		for i := 3; i <= 5; i++ {
			if strings.Contains(input[i], numbers[1]) {
				numbers[3] = input[i]

				break
			}
		}

		//len=6 segments (0,6,9) are 6,7,8 in input when sorted
		//this finds 9 as it is the only len6 with with 9-3=len1
		//This is fucking dirty and gross but it works so heh
		//we do get e for free as it is the excluded segement
		segments["e"], numbers[9] = solveNumbers(input, numbers[3], 6, 8)

		//we can get f by removing e and 1 from 4
		segments["f"] = (subtractSegments(numbers[4], numbers[1]))
		segments["f"] = (subtractSegments(segments["f"], segments["e"]))

		//we can get c from 9 - 4 - d
		segments["c"] = subtractSegments(subtractSegments(numbers[9], numbers[4]), segments["d"])

		//Can ID g from whatever 9 is missing
		for x := range segments {
			if !strings.Contains(numbers[9], x) {
				segments["g"] = x
			}
		}

		// Can ID 0 from not f
		for i := 6; i <= 8; i++ {
			if input[i] != "" && !strings.Contains(input[i], segments["f"]) {
				numbers[0] = input[i]

			}
		}

		//Can ID 6 by process of elimination (6char)
		for i := 6; i <= 8; i++ {
			if input[i] != numbers[0] && input[i] != numbers[9] && input[i] != "" {
				numbers[6] = input[i]

			}
		}

		//Can ID a from not in 6
		for x := range segments {
			if !strings.Contains(numbers[6], x) {
				segments["a"] = x
			}
		}
		//Can ID b from 1-a
		segments["b"] = (subtractSegments(numbers[1], segments["a"]))

		numbers[2] = sortString(segments["a"] + segments["c"] + segments["d"] + segments["e"] + segments["g"])
		numbers[5] = sortString(segments["a"] + segments["b"] + segments["d"] + segments["f"] + segments["g"])

		//switch?
		// for x, y := range numbers {
		// 	oppo[y] = x
		// }
		oppo := map[string]int{numbers[0]: 0, numbers[1]: 1, numbers[2]: 2, numbers[3]: 3, numbers[4]: 4, numbers[5]: 5, numbers[6]: 6, numbers[7]: 7, numbers[8]: 8, numbers[9]: 9}

		var temps string
		// var tempi int
		for _, x := range output {
			temps += strconv.Itoa(oppo[x])
		}
		answer, _ = strconv.Atoi(temps)

		if true {
			fmt.Println(line)
			fmt.Println(segments)
			fmt.Println(numbers)
			fmt.Println(answer)
			break
		}
	}

}
