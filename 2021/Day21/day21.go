package main

import (
	"fmt"
)

// var path = "input.txt" //path to problem input

const END int = 21

var rolls []int
var outcomes = make(map[int]int)
var a_score int64 = 0
var b_score int64 = 0

//This badboy rolls all possible OUTCOMES and then scores them based on how frequently they'd occur
func roll(score_a int, score_b int, a_position int, b_position int, a_cumulative int, b_cumulative int) {
	for x, xo := range outcomes {
		//every loop, reset all the stuff
		as := score_a
		ap := a_position
		apath := a_cumulative

		//add the roll to the position
		ap += x

		//Score the position (same limitations as Part 1)
		if ap%10 == 0 {
			as += 10
		} else {
			as += ap % 10
		}
		//Score the PATH, this means the number of universes this outcome would have occurred in
		//Example: If A wins by rolling all nines, since nine can only happen once it would be 1*1*1
		//		...however, if A wins by all sixes, each of which can happen 7 ways, it actually happened
		//		...in 7*7*7*7 = 2401 universes.  Anyway, we keep track of that there
		apath *= xo

		//If we've reached the end (21), add the cumulative "path" to A's wins
		//Not ENTIRELY sure why I have to multiply these, but it gives the right answer
		//Guessing because A wins in apath universes it ALSO wins in bpath universes, therefore a*b=true wins
		if as >= END {
			a_score += int64(apath * b_cumulative) //since B hasn't rolled yet...
			continue
		}

		for y, yo := range outcomes {
			//same as above
			bs := score_b
			bp := b_position
			bpath := b_cumulative

			bp += y

			if bp%10 == 0 {
				bs += 10
			} else {
				bs += bp % 10
			}

			bpath *= yo

			if bs >= END {
				b_score += int64(bpath * apath)
				continue
			}

			// Since neither won, we need to roll again
			roll(as, bs, ap, bp, apath, bpath)
		}
	}
}

func main() {
	//Starting positions as per puzzle input
	start_1 := 6
	start_2 := 7

	//sample inputs
	// start_1 = 4
	// start_2 = 8

	var dice int = 1
	score_1 := 0
	score_2 := 0

	p1 := start_1
	p2 := start_2
	//PART 1
	for {
		//"Roll" three times
		p1 += 3*dice + 3
		dice += 3

		// Map current position to score
		if p1%10 == 0 { //since we're tracking absolute position and then modding it...
			score_1 += 10 //we need to deal with 10, as 10%10 = 0 but would score 10
		} else {
			score_1 += (p1 % 10)
		}

		//If we've met the Part 1 winning conditions, cool
		if score_1 >= 1000 || score_2 >= 1000 {
			// fmt.Println(dice, score_1, score_2)
			break
		}

		//Otherwise do all the same things for Player 2
		p2 += 3*dice + 3
		dice += 3

		if p2%10 == 0 {
			score_2 += 10
		} else {
			score_2 += (p2 % 10)
		}
		if score_1 >= 1000 || score_2 >= 1000 {
			// fmt.Println(dice, score_1, score_2)
			break
		}
	}

	//Score per instructions
	var answer_pt1 int
	if score_1 > score_2 {
		answer_pt1 = (dice - 1) * score_2
	} else {
		answer_pt1 = (dice - 1) * score_1
	}
	fmt.Printf("\nThe answer to Part 1 is: %v\n", answer_pt1)

	//PART 2!
	//determine all possible rolls and sort them into their outcomes
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			for z := 1; z <= 3; z++ {
				rolls = append(rolls, x+y+z)
				outcomes[x+y+z] += 1
			}
		}
	}

	fmt.Printf("\nThere are %v possible rolls, which are %v\n", len(rolls), rolls)
	fmt.Printf("However, there are %v possible outcomes, which are %v\n", len(outcomes), outcomes)

	//Recursively roll until we've rolled... everything
	roll(0, 0, start_1, start_2, 1, 1)

	fmt.Printf("\nFor Part 2, Player A (winner) won %v times and Player B won %v times.\n", a_score, b_score)
	fmt.Printf("The answer to Part 2 is %v. \n\n", a_score)

}
