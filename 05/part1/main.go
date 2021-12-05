package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	matcher, _ := regexp.Compile(`(\d+),(\d+) -> (\d+),(\d+)`)
	grid := make(map[[2]int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m := matcher.FindStringSubmatch(line)
		x1, err := strconv.Atoi(m[1])
		if err != nil {
			fmt.Println("error parsing line", line)
			return
		}
		y1, err := strconv.Atoi(m[2])
		if err != nil {
			fmt.Println("error parsing line", line)
			return
		}
		x2, err := strconv.Atoi(m[3])
		if err != nil {
			fmt.Println("error parsing line", line)
			return
		}
		y2, err := strconv.Atoi(m[4])
		if err != nil {
			fmt.Println("error parsing line", line)
			return
		}

		// lines are only vertical or horizontal
		dx, dy := 0, 0
		if x1 == x2 {
			dx = 0
		} else if x1 < x2 {
			dx = 1
		} else {
			dx = -1
		}
		if y1 == y2 {
			dy = 0
		} else if y1 < y2 {
			dy = 1
		} else {
			dy = -1
		}

		if dx != 0 && dy != 0 {
			// not horizontal or vertical
			continue
		}

		x, y := x1, y1
		for {
			pos := [2]int{x, y}
			grid[pos]++

			if x == x2 && y == y2 {
				break
			}

			x += dx
			y += dy
		}
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			c := grid[[2]int{x, y}]
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(c)
			}
		}
		fmt.Println()
	}

	total := 0
	for _, overlapCount := range grid {
		if overlapCount >= 2 {
			//fmt.Println(pos, "overlaps", overlapCount)
			total += 1
		}
	}
	fmt.Println("total overlaps >2", total)
}
