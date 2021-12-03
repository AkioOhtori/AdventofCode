package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const PART int = 2
const OXYGEN bool = true
const C02 bool = false

var path = "input.txt" //path to problem input
var one string = "1"   //a thing we need because Go
var length int = 0     //not sure this is still needed but...

// var zero string = "0"

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func checkDominant(master []string, i int, mode bool) []string {
	var new_zero []string //where we store all the values in when master[x][i] == 0
	var new_one []string  //where we store all the values in when master[x][i] == 1

	// Iterates over the length of the dataset (1000 to start)
	// Requires the location of the examined bit (i)
	for x := 0; x < len(master); x++ {
		if master[x][i:i+1] == one { //if bit is 1, store it in the 1s slice
			new_one = append(new_one, master[x])
		} else { //otherwise store it in the 0s slice
			new_zero = append(new_zero, master[x])
		}
	}
	// Checking Oxygen Rules
	if mode {
		if len(new_one) >= len(new_zero) {
			master = new_one
		} else {
			master = new_zero
		}

	} else { // Checking CO2 Rules
		if len(new_zero) > len(new_one) {
			master = new_one
		} else {
			master = new_zero
		}
	} //I feel like there is a way to do this without the conditions but heh
	return master
}

func checkBits(z []string, mode bool) []string {
	for i := 0; i < len(z[0]); i++ {
		if len(z) == 1 {
			break
		}
		z = checkDominant(z, i, mode)
	}
	return z
}

func main() {

	var places = [12]int{}
	//Open input file
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// PART 1
	if PART == 1 {
		// Check to see how many 1s
		for scanner.Scan() {
			length++
			txt := (scanner.Text())
			for i := 0; i < len(txt); i++ {
				if txt[i:(i+1)] == one {
					places[i]++
				}
			}
		}

		// Little pre-math
		thresh := length / 2
		var gamma string = ""
		var epsilon string = ""

		//Iterate over the "strings" to determine which is more used
		for i := 0; i < len(places); i++ {
			if places[i] > thresh {
				gamma += "1"
				epsilon += "0"
			} else {
				gamma += "0"
				epsilon += "1"
			}
		}

		// Convert Binary strings to decimal for mathing
		gamma_dec, _ := strconv.ParseInt(gamma, 2, 64)
		epsilon_dec, _ := strconv.ParseInt(epsilon, 2, 64)

		// And we're done with part 1!
		fmt.Printf("\nI think we're finished and the gamma is %v and epsilon is %v.\n", gamma, epsilon)
		fmt.Printf("The answer is %v!", gamma_dec*epsilon_dec)
		// BEGIN PART 2
	} else {
		var master []string

		// Load all of the input into memory (master)
		for scanner.Scan() {
			master = append(master, scanner.Text())
		}

		ox := checkBits(master, OXYGEN)
		cO2 := checkBits(master, C02)

		// Convert binary string to decimal int
		ox_dec, _ := strconv.ParseInt(ox[0], 2, 64)
		cO2_dec, _ := strconv.ParseInt(cO2[0], 2, 64)

		// Print the answer
		fmt.Printf("I think we're finished and the oxygen is %v and CO2 is %v.\n", ox_dec, cO2_dec)
		fmt.Printf("The answer is %v!", ox_dec*cO2_dec)
	}

}
