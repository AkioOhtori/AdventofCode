package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Function to calculate the winner because I didn't have the heart to put this inline
func calcWinner(card [5][]string, winning_no int) int {
	var a int = 0
	for _, row := range card {
		for _, val := range row {
			if val == "" {
				continue
			} else {
				v, _ := strconv.Atoi(val)
				a += v
			}
		}
	}
	return a * winning_no
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

	// PART 1
	if PART == 1 {
		var master []string
		// Load all of input into memory (master)
		for scanner.Scan() {
			master = append(master, scanner.Text())
		}

		number_of_cards := ((len(master) - 1) / 6)

		//Split out input
		input := strings.Split(master[0], ",")

		var new_card [5][]string
		var card_num int = 0
		cards := make([][5][]string, number_of_cards)
		answers := make([][2][5]int, number_of_cards)

		var ind int = 0

		//build out deck of bingo cards for checking
		//cards are stored in "cards" which is a 100x5x5 slice except for when it isnt'
		//Had to do some hard coding because Go gonna Go
		for i := 2; i < len(master); i++ {
			if master[i] == "" {
				//do end of card stuff
				card_num++
				ind = 0
			} else {
				new_card[ind] = strings.Fields(master[i])
				cards[card_num][ind] = new_card[ind]
				ind++
			}
		}

		// Ok so now we have all the bingo cards loaded in, format:
		// cards[card no][row][column]
		// So now we iterate over the input in a bunch of nested loops?
	out:
		//on the fourth day of sAdvent, my true love gave to me:
		//four nested for loops
		//three dimentional arrys
		//two potential answers
		//and a sneaky function call in a print statemtn
		for _, val := range input {
			for card_no, card := range cards {
				for row := 0; row < 5; row++ {
					for col := 0; col < 5; col++ {
						if card[row][col] == val {
							fmt.Printf("hit! %v, %v, %v\n", card_no, row, col)
							answers[card_no][0][row]++
							answers[card_no][1][col]++
							card[row][col] = ""
							if answers[card_no][0][row] == 5 || answers[card_no][1][col] == 5 {
								w := [3]int{card_no, row, col}
								ii, _ := strconv.Atoi(val)
								fmt.Printf("We have a winner and it is probably %v for a winning number of %v\n", w, calcWinner(card, ii))
								break out
							}
						}

					} //end col
				} //end row
			} //end cards
		}
		// And we're done with part 1!
	} else {
		var master []string
		// Load all of input into memory (master)
		for scanner.Scan() {
			master = append(master, scanner.Text())
		}

		number_of_cards := ((len(master) - 1) / 6)

		//Split out input
		input := strings.Split(master[0], ",")

		var new_card [5][]string
		var card_num int = 0
		cards := make([][5][]string, number_of_cards)
		answers := make([][2][5]int, number_of_cards)
		var winners int = 0
		empty := []string{"", "", "", "", ""}
		erase := [5][]string{empty, empty, empty, empty, empty}

		var ind int = 0

		//build out deck of bingo cards for checking
		//cards are stored in "cards" which is a 100x5x5 slice except for when it isnt'
		//Had to do some hard coding because Go gonna Go
		for i := 2; i < len(master); i++ {
			if master[i] == "" {
				//do end of card stuff
				card_num++
				ind = 0
			} else {
				new_card[ind] = strings.Fields(master[i])
				cards[card_num][ind] = new_card[ind]
				ind++
			}
		}

		// Ok so now we have all the bingo cards loaded in, format:
		// cards[card no][row][column]
		// So now we iterate over the input in a bunch of nested loops?
	outout:
		//on the fourth day of sAdvent, my true love gave to me:
		//four nested for loops
		//three dimentional arrys
		//two potential answers
		//and a sneaky function call in a print statemtn
		for _, val := range input {
			for card_no, card := range cards {
				for row := 0; row < 5; row++ {
					for col := 0; col < 5; col++ {
						if card[row][col] == val {
							fmt.Printf("hit! %v, %v, %v\n", card_no, row, col)
							answers[card_no][0][row]++
							answers[card_no][1][col]++
							card[row][col] = ""
							if answers[card_no][0][row] == 5 || answers[card_no][1][col] == 5 {
								winners++
								//check if all cards have won yet
								if winners == number_of_cards {
									w := [3]int{card_no, row, col}
									ii, _ := strconv.Atoi(val)
									fmt.Printf("We have a winner and it is probably %v for a winning number of %v\n", w, calcWinner(card, ii))
									break outout
								} else { //So this card can't score again, erase it
									card = erase
									cards[card_no] = erase
									break
								}
							} //end of winner checking
						} //end of value checking
					} //end col
				} //end row
			} //end cards
		} //end of input value checking
	} //end of part 2

} //EOF
