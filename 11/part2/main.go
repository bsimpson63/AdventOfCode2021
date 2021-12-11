package main

import (
	"bufio"
	"fmt"
	"os"
)

func step(grid map[[2]int]int) int {
	/*
		First, the energy level of each octopus increases by 1.
		Then, any octopus with an energy level greater than 9 flashes.
			This increases the energy level of all adjacent octopuses by 1,
			including octopuses that are diagonally adjacent.
			If this causes an octopus to have an energy level greater than 9,
			it also flashes. This process continues as long as new octopuses
			keep having their energy level increased beyond 9.
			(An octopus can only flash at most once per step.)
		Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
	*/

	hasFlashed := make(map[[2]int]bool)

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			grid[[2]int{row, col}]++
		}
	}

	deltas := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	newFlashes := true
	for {
		if !newFlashes {
			break
		}
		newFlashes = false

		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				pos := [2]int{row, col}
				if grid[pos] > 9 && !hasFlashed[pos] {
					hasFlashed[pos] = true
					newFlashes = true

					for _, delta := range deltas {
						neighbor := [2]int{pos[0] + delta[0], pos[1] + delta[1]}
						if _, isOK := grid[neighbor]; isOK {
							grid[neighbor]++
						}
					}
				}
			}
		}
	}

	for pos := range hasFlashed {
		grid[pos] = 0
	}

	return len(hasFlashed)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make(map[[2]int]int)
	row := 0
	col := 0
	for scanner.Scan() {
		line := scanner.Text()
		col = 0
		for _, c := range line {
			num := int(c - '0')
			grid[[2]int{row, col}] = num
			col++
		}
		row++
	}

	i := 0
	for {
		i++
		flashCount := step(grid)
		if flashCount == 100 {
			fmt.Println(i)
			break
		}
	}
}
