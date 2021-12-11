package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func findBasinSize(grid map[[2]int]int, loc [2]int) int {
	deltas := [][2]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}

	seen := make(map[[2]int]bool)
	stack := make([][2]int, 0)

	// initialize the stack with all the neighbors
	for _, delta := range deltas {
		neighborLoc := [2]int{loc[0] + delta[0], loc[1] + delta[1]}
		stack = append(stack, neighborLoc)
	}

	for {
		if len(stack) == 0 {
			break
		}

		pos := stack[0]
		stack = append(stack[:0], stack[1:]...)

		if alreadySeen := seen[pos]; alreadySeen {
			continue
		}

		v, isOK := grid[pos]
		if v == 9 || !isOK {
			continue
		}

		seen[pos] = true
		for _, delta := range deltas {
			neighborLoc := [2]int{pos[0] + delta[0], pos[1] + delta[1]}
			stack = append(stack, neighborLoc)
		}
	}
	return len(seen)
}

func findLowPoints(grid map[[2]int]int, nRows int, nCols int) [][2]int {
	lowPoints := make([][2]int, 0)
	deltas := [][2]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
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
				lowPoints = append(lowPoints, selfLoc)
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

	basinSizes := make([]int, 0)
	for _, lowPoint := range lowPoints {
		s := findBasinSize(grid, lowPoint)
		fmt.Println(lowPoint, s)
		basinSizes = append(basinSizes, s)
	}
	sort.Ints(basinSizes)
	p := 1
	for i := len(basinSizes) - 1; i >= len(basinSizes)-3; i-- {
		p *= basinSizes[i]
		fmt.Println(basinSizes[i])
	}
	fmt.Println("product", p)

}
