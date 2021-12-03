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
		if master[x][i:i+1] == one {
			new_one = append(new_one, master[x])
		} else {
			new_zero = append(new_zero, master[x])
		}
	}
	if mode {
		if len(new_one) >= len(new_zero) {
			master = new_one
		} else {
			master = new_zero
		}

	} else {
		if len(new_zero) > len(new_one) {
			master = new_one
		} else {
			master = new_zero
		}
	}
	return master
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

		// And we're done!
		fmt.Printf("\nI think we're finished and the gamma is %v and epsilon is %v.\n", gamma, epsilon)
		fmt.Printf("The answer is %v!", gamma_dec*epsilon_dec)
	} else {
		var master []string
		var ox []string
		var cO2 []string

		for scanner.Scan() {
			master = append(master, scanner.Text())
		}
		cO2 = master
		ox = master

		for i := 0; i < len(ox[0]); i++ {
			if len(ox) == 1 {
				// fmt.Println("I think we're done?!")
				// fmt.Print(ox)
				break
			}
			ox = checkDominant(ox, i, OXYGEN)
		}
		for i := 0; i < len(cO2[0]); i++ {
			if len(cO2) == 1 {
				// fmt.Println("I think we're done?!")
				// fmt.Print(cO2)
				break
			}
			cO2 = checkDominant(cO2, i, C02)
		}

		fmt.Println(ox, cO2)
		ox_dec, _ := strconv.ParseInt(ox[0], 2, 64)
		cO2_dec, _ := strconv.ParseInt(cO2[0], 2, 64)
		fmt.Printf("I think we're finished and the oxygen is %v and CO2 is %v.\n", ox_dec, cO2_dec)
		fmt.Printf("The answer is %v!", ox_dec*cO2_dec)

		/*
			for i := 0; i < len(master[0]); i++ {
				if len(master) == 1 {
					fmt.Println("I think we're done?!")
					fmt.Print(master)
					break
				} //else
				for x := 0; x < len(master); x++ {
					if master[x][i:i+1] == one {
						new_one = append(new_one, master[x])
					} else {
						new_zero = append(new_zero, master[x])
					}
				}
				if len(new_one) >= (len(master) / 2) {
					master = new_one
				} else {
					master = new_zero
				}
				new_one = nil
				new_zero = nil
			}*/

	}

}
