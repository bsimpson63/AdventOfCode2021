package main

import "fmt"

/*
The probe's x,y position starts at 0,0. Then, it will follow some trajectory by moving in steps. On each step, these changes occur in the following order:

The probe's x position increases by its x velocity.
The probe's y position increases by its y velocity.
Due to drag, the probe's x velocity changes by 1 toward the value 0;
  that is, it decreases by 1 if it is greater than 0, increases by 1
  if it is less than 0, or does not change if it is already 0.
Due to gravity, the probe's y velocity decreases by 1.

How many distinct initial velocity values cause the probe to be
within the target area after any step?

find vyMax: 90
find vyMin: -91

for every vy within those bounds find the # that ends up in the zone

simple example:
vyMax: 9
vyMin: -10

*/

//calculate nsteps for various vx0

func calcNSteps(vy0 int, yMin int, yMax int) []int {
	steps := 0
	y := 0
	vy := vy0
	nSteps := make([]int, 0)

	for {
		y = y + vy
		vy -= 1
		steps += 1

		if y <= yMax && y >= yMin {
			nSteps = append(nSteps, steps)
		}

		if y < yMin {
			break
		}
	}
	return nSteps
}

func calcNStepsX(vx0 int, xMin int, xMax int, stepsMax int) []int {
	steps := 0
	x := 0
	vx := vx0
	nSteps := make([]int, 0)

	for {
		x = x + vx
		if vx > 0 {
			vx--
		}
		steps++

		if x <= xMax && x >= xMin {
			nSteps = append(nSteps, steps)
		}

		if vx == 0 && steps > stepsMax {
			// we've stalled
			break
		}

		if x > xMax {
			// we've overshot
			break
		}
	}
	return nSteps
}

func main() {
	/*
		yMin := -10
		yMax := -5
		xMin := 20
		xMax := 30
		vyMin := -10
		vyMax := 9
	*/

	// target area: x=244..303, y=-91..-54
	yMin := -91
	yMax := -54
	xMin := 244
	xMax := 303
	vyMin := -91
	vyMax := 90

	stepMap := make(map[int][]int)

	maxSteps := 0
	for vy0 := vyMin; vy0 <= vyMax; vy0++ {
		nSteps := calcNSteps(vy0, yMin, yMax)
		fmt.Println(vy0, " in range at:", nSteps)

		for _, s := range nSteps {
			stepMap[s] = append(stepMap[s], vy0)
			if s > maxSteps {
				maxSteps = s
			}
		}
	}

	for s, vy0s := range stepMap {
		fmt.Println(s, "steps works for vy0:", vy0s)
	}

	velocities := make(map[[2]int]bool)
	for vx0 := 0; vx0 < 1000; vx0++ {
		nSteps := calcNStepsX(vx0, xMin, xMax, maxSteps)
		for _, s := range nSteps {
			fmt.Printf("%d steps: vx0:%d, vy0:%d\n", s, vx0, stepMap[s])
			for _, vy0 := range stepMap[s] {
				velocities[[2]int{vx0, vy0}] = true
			}

		}
	}
	fmt.Println("answer:", len(velocities))
	fmt.Println("maxSteps:", maxSteps)
}
