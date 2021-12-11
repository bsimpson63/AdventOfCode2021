package main

import (
	"bufio"
	"fmt"
	"os"
)

func findLowPoints(grid map[[2]int]int, nRows int, nCols int) []int {
	lowPoints := make([]int, 0)
	deltas := [][2]int{
		//{-1, -1},
		{-1, 0},
		//{-1, 1},
		{0, -1},
		{0, 1},
		//{1, -1},
		{1, 0},
		//{1, 1},
	}

	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			selfLoc := [2]int{row, col}
			self := grid[selfLoc]
			isLowPoint := true
			for _, delta := range deltas {
				dx := delta[0]
				dy := delta[1]
				neighborLoc := [2]int{row + dx, col + dy}
				neighbor, isOK := grid[neighborLoc]
				if !isOK {
					continue
				}
				if neighbor <= self {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				lowPoints = append(lowPoints, self)
			}
		}
	}
	return lowPoints
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
	fmt.Println(grid)
	lowPoints := findLowPoints(grid, row, col)
	fmt.Println("low points", lowPoints)
	s := 0
	for _, p := range lowPoints {
		s += 1 + p
	}
	fmt.Println("sum of risk levels", s)
}
