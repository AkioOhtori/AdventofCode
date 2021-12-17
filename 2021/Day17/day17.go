package main

import (
	"fmt"
)

const PART int = 2

func calculateTrajectory(x_vel int, y_vel int, target_x [2]int, target_y [2]int) (y_max int, hit bool) {
	/*
		1) The probe's x position increases by its x velocity.
		2) The probe's y position increases by its y velocity.
		3) Due to drag, the probe's x velocity changes by 1 toward the value 0;
			that is, it decreases by 1 if it is greater than 0,
			increases by 1 if it is less than 0, or does not change if it is already 0.
		4) Due to gravity, the probe's y velocity decreases by 1.
	*/
	var x_pos int = 0
	var y_pos int = 0
	hit = false
	// var y_max int = 0

	for i := 0; i < 1000; i++ {
		// fmt.Printf("Started step %v with position %v, %v and velocity %v, %v \n", i, x_pos, y_pos, x_vel, y_vel)
		if y_pos > y_max {
			y_max = y_pos
		}

		x_pos += x_vel
		y_pos += y_vel

		y_vel--
		if x_vel > 0 {
			x_vel--
		} else if x_vel < 0 {
			x_vel++
		}

		// fmt.Printf("And ended with position %v, %v and velocity %v, %v \n", x_pos, y_pos, x_vel, y_vel)

		if x_pos > target_x[1] || y_pos < target_y[1] { //tells me if we've passed in either x or y
			// fmt.Println(x_pos, target_x[1], y_pos, target_y[1])
			y_max = 0
			return
		} else if x_pos >= target_x[0] && y_pos <= target_y[0] {
			if x_pos <= target_x[1] && y_pos >= target_y[1] {
				hit = true
				return
			}
		}
	}
	y_max = 0
	return
}

func calcX(x_vel int, target_x [2]int) (int, int) {
	var x_pos int = 0
	for i := 0; i < 1000; i++ {
		// fmt.Printf("Starting step %v with vel = %v at pos = %v ", i, x_vel-i, x_pos)
		if (x_vel - i) < 0 {
			x_vel = 0
			// fmt.Printf("and ending\n")
			if x_pos > target_x[0] {
				return (i - 1), 0
			} else {
				return -100, -100
			}

		} else {
			x_pos += (x_vel - i)
		}
		// fmt.Printf("and ending with %v, %v\n", x_vel-(i+1), x_pos)

		if x_pos > target_x[1] {
			return (i - 1), (x_vel + (i - 1))
		}
	}

	return 0, 0

}

func main() {
	//sample input
	target_x := [2]int{20, 30}
	target_y := [2]int{-5, -10}

	//cheating and just hard coding the input
	target_x = [2]int{57, 116}
	target_y = [2]int{-148, -198}

	x_steps := []int{}
	var answer_pt2 int = 0

	for x := 1; x <= target_x[1]; x++ {
		//steps, velocity
		s, tv := calcX(x, target_x)
		if PART == 1 {
			if s > 1 && tv == 0 {
				x_steps = append(x_steps, x)
			}
		} else {
			if s >= 0 {
				x_steps = append(x_steps, x)
			}
		}
	}

	ym := 0
	for _, x := range x_steps {
		for y := (target_y[1]); y < -target_y[1]; y++ {
			t, h := calculateTrajectory(x, y, target_x, target_y)
			if t > ym {
				ym = t
			}
			if h {
				answer_pt2++
			}
		}
	}

	fmt.Printf("The answer to Part 1 is y = %v and Part 2 is %v\n", ym, answer_pt2)
}
