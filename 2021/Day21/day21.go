package main

import (
	"fmt"
	"sort"
)

// var path = "input.txt" //path to problem input

const END int = 21

var rolls []int
var outcomes = make(map[int]int)
var ascore int64 = 0
var bscore int64 = 0

func roll(score_a int, score_b int, a_position int, b_position int, a_cumulative int, b_cumulative int) {
	for x, xo := range outcomes {
		as := score_a
		ap := a_position
		apath := a_cumulative

		ap += x

		if ap%10 == 0 {
			as += 10
		} else {
			as += ap % 10
		}
		apath *= xo

		if as >= END {
			ascore += int64(apath) * int64(b_cumulative)
			continue
		}

		for y, yo := range outcomes {
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
				bscore += int64(bpath * apath)
				continue
			}

			if as < END && bs < END {
				// fmt.Println(apath, bpath)
				roll(as, bs, ap, bp, apath, bpath)
			} else {
				fmt.Println("ERRORERRORERRORERRORERRORERRORERRORERRORERRORERRORERROR")
			}
		}
	}
}

func main() {
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
		p1 += 3*dice + 3
		// fmt.Printf("Player 1 rolled a %v+%v+%v for a new place of %v and ", dice, dice+1, dice+2, p1%10)
		dice += 3
		if p1%10 == 0 {
			score_1 += 10
		} else {
			score_1 += (p1 % 10)
		}
		if score_1 >= 1000 || score_2 >= 1000 {
			fmt.Println(dice, score_1, score_2)
			break
		}

		p2 += 3*dice + 3
		// fmt.Printf("Player 2 rolled a %v+%v+%v for a new place of %v\n", dice, dice+1, dice+2, p2%10)
		dice += 3

		if p2%10 == 0 {
			score_2 += 10
		} else {
			score_2 += (p2 % 10)
		}
		// fmt.Printf("P1 Score = %v \t P2 Score = %v\n", score_1, score_2)
		if score_1 >= 1000 || score_2 >= 1000 {
			fmt.Println(dice, score_1, score_2)
			break
		}

	}
	var answer_pt1 int
	if score_1 > score_2 {
		answer_pt1 = (dice - 1) * score_2
	} else {
		answer_pt1 = (dice - 1) * score_1
	}
	fmt.Println(answer_pt1)

	p1 = start_1
	p2 = start_2

	//determine all possible rolls
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			for z := 1; z <= 3; z++ {
				rolls = append(rolls, x+y+z)
				outcomes[x+y+z] += 1
			}
		}
	}
	sort.Ints(rolls)

	fmt.Println(len(rolls), rolls)
	fmt.Println(len(outcomes), outcomes)

	roll(0, 0, start_1, start_2, 1, 1)
	fmt.Println(ascore, bscore)

}
