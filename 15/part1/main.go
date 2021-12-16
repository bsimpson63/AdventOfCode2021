package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	pos := [2]int{0, 0}
	minToPos := make(map[[2]int]int)
	stack := make([][3]int, 0) // (x, y, cost)
	first := [3]int{pos[0], pos[1], 0}
	stack = append(stack, first)

	deltas := [][2]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}

	/*
		at each position consider moving to each adjacent position
		if we've already been at the new position and the earlier move
		was cheaper to get there, don't go

	*/

	dest := [2]int{99, 99}
	count := 0
	for {
		if len(stack) == 0 {
			break
		}

		count++
		if count%1000 == 0 {
			fmt.Println(count, minToPos[dest], len(stack))
		}

		n := len(stack) - 1
		item := stack[n]
		pos := [2]int{item[0], item[1]}
		cost := item[2]
		stack = append(stack[:n], stack[n+1:]...)

		for _, delta := range deltas {
			neighbor := [2]int{pos[0] + delta[0], pos[1] + delta[1]}
			if _, isOK := grid[neighbor]; !isOK {
				// this position doesn't exist
				continue
			}

			nextCost := cost + grid[neighbor]
			if bestCost, isOK := minToPos[neighbor]; !isOK {
				minToPos[neighbor] = nextCost
			} else {
				if nextCost < bestCost {
					minToPos[neighbor] = nextCost
				} else {
					// we've already been here for cheaper
					continue
				}
			}

			stack = append(stack, [3]int{neighbor[0], neighbor[1], nextCost})
		}

	}
	fmt.Println(minToPos[dest])
}
