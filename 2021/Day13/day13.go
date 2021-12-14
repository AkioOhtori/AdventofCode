package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
const PART int = 1

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func prettyPrintMatrixLetters(m [][]int) {
	fmt.Println()
	for _, y := range m {
		for i, x := range y {
			if i%5 == 0 {
				fmt.Printf("\t")
			}
			if x == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func prettyPrintMatrix2D(m [][]int) {
	fmt.Println()
	for _, x := range m {
		fmt.Println(x)
	}
	fmt.Println()
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

	var input_int [][]int
	var fold_instructions [][]string
	var answer_pt1 int = 0
	var max_x int
	var max_y int

	//Go through the instructions and convert them to slices of slices, [ORIGIN],[DESTIN]
	for scanner.Scan() {
		temp_str := scanner.Text()
		if len(temp_str) > 0 {
			if temp_str[:1] != "f" {
				input_str := strings.Split(temp_str, ",")

				var temp_int []int
				for _, x := range input_str {
					y, _ := strconv.Atoi(x)
					temp_int = append(temp_int, y)
				}
				input_int = append(input_int, temp_int)
			} else {
				t := strings.Split(temp_str, "=")
				fold_instructions = append(fold_instructions, []string{t[0][11:], t[1]})
			}
		}
	}

	for _, f := range fold_instructions {
		if f[0] == "y" && max_y == 0 {
			max_y, _ = strconv.Atoi(f[1])
			max_y = max_y * 2
		} else if f[0] == "x" && max_x == 0 {
			max_x, _ = strconv.Atoi(f[1])
			max_x = max_x * 2
		}
		if max_x != 0 && max_y != 0 {
			break
		}
	}
	fmt.Printf("I think we're going to make the matrix x = %v, y = %v\n", max_x, max_y)

	/*

		                                   ,----,            ,---,
		         ,--.    ,----..         ,/   .`|         ,`--.' |
		       ,--.'|   /   /   \      ,`   .'  :   ,---,.|   :  :
		   ,--,:  : |  /   .     :   ;    ;     / ,'  .' |'   '  ;
		,`--.'`|  ' : .   /   ;.  \.'___,/    ,',---.'   ||   |  |
		|   :  :  | |.   ;   /  ` ;|    :     | |   |   .''   :  ;
		:   |   \ | :;   |  ; \ ; |;    |.';  ; :   :  |-,|   |  '
		|   : '  '; ||   :  | ; | '`----'  |  | :   |  ;/|'   :  |
		'   ' ;.    ;.   |  ' ' ' :    '   :  ; |   :   .';   |  ;
		|   | | \   |'   ;  \; /  |    |   |  ' |   |  |-,`---'. |
		'   : |  ; .' \   \  ',  /     '   :  | '   :  ;/| `--..`;
		|   | '`--'    ;   :    /      ;   |.'  |   |    \.--,_
		'   : |         \   \ .'       '---'    |   :   .'|    |`.
		;   |.'          `---`                  |   | ,'  `-- -`, ;
		'---'                                   `----'      '---`"

		Apparently not all inputs are created equal and I got the short straw on this one.
		The above code determines the "paper" size based on the fold directions, rather than the dot locations.
		This was discovered after I couldn't get the code to work with my input, but it worked great with a friend's.
		Yes.  I am bitter.  This is why I made this over the top note.
	*/
	//Make the "paper" based on max size + 1 (zero ref)
	paper := make([][]int, max_y+1)
	for i := 0; i < len(paper); i++ {
		paper[i] = make([]int, max_x+1)
	}
	for _, dot := range input_int {
		paper[dot[1]][dot[0]] = 1
	}

	for i := 0; i < len(fold_instructions); i++ {
		if fold_instructions[i][0] == "y" {

			along, _ := strconv.Atoi(fold_instructions[i][1]) //where to fold
			new := make([][]int, along)                       //make a new, blank matrix

			for y := 0; y < along; y++ { //fill the blank matrix
				new[y] = make([]int, len(paper[y]))
				copy(new[y], paper[y])

				for x := 0; x < len(paper[y]); x++ { //fold the paper
					new[y][x] = (new[y][x] | paper[len(paper)-1-y][x])
				}
			}

			paper = new

		} else { //x
			along, _ := strconv.Atoi(fold_instructions[i][1]) //where to fold
			new := make([][]int, len(paper))                  //make a new, correctly sized blank
			for y := 0; y < len(paper); y++ {                 //fill the blank matrix
				new[y] = make([]int, along)
				copy(new[y], paper[y][:along])

				for x := 0; x < along; x++ { //fold the paper
					new[y][x] = (new[y][x] | paper[y][len(paper[y])-1-x])
				}
				// break
			}
			// Make the folded version the real one
			paper = new

		}

		if i == 0 { //Part 1 answer
			for _, y := range paper {
				for _, x := range y {
					answer_pt1 += x
				}
			}
		}

	}
	fmt.Printf("The new, folded matrix is x=%v, y=%v\n", len(paper[0]), len(paper))

	// prettyPrintMatrix2D(paper)
	//Calculate the answer

	prettyPrintMatrix2D(paper)
	prettyPrintMatrixLetters(paper)
	fmt.Printf("The Answer to part 1 is: %v \n", answer_pt1)

} //EOF
