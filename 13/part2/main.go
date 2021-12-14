package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func foldUp(dots map[[2]int]bool, yFold int) map[[2]int]bool {
	newDots := make(map[[2]int]bool)

	for pos := range dots {
		x := pos[0]
		y := pos[1]
		if y == yFold {
			fmt.Println("oops, got dot on a fold!", y, yFold)
			continue
		}

		if y < yFold {
			newDots[pos] = true
			continue
		}

		// reflect this point
		yReflected := 2*yFold - y
		newDots[[2]int{x, yReflected}] = true
	}
	return newDots
}

func foldLeft(dots map[[2]int]bool, xFold int) map[[2]int]bool {
	newDots := make(map[[2]int]bool)

	for pos := range dots {
		x := pos[0]
		y := pos[1]
		if x == xFold {
			fmt.Println("oops, got dot on a fold!", x, xFold)
			continue
		}

		if x < xFold {
			newDots[pos] = true
			continue
		}

		// reflect this point
		xReflected := 2*xFold - x
		newDots[[2]int{xReflected, y}] = true
	}
	return newDots
}

func displayDots(dots map[[2]int]bool) {
	xMax, yMax := 0, 0
	for pos := range dots {
		if pos[0] > xMax {
			xMax = pos[0]
		}
		if pos[1] > yMax {
			yMax = pos[1]
		}
	}
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if isLit := dots[[2]int{x, y}]; isLit {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	dots := make(map[[2]int]bool)
	folds := make([][2]int, 0)
	matcher, _ := regexp.Compile(`fold along ([xy])=(\d+)`)
	readingDots := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readingDots = false
			continue
		}

		if readingDots {
			parts := strings.Split(line, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("bad input")
				return
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("bad input")
				return
			}
			dots[[2]int{x, y}] = true
		} else {
			m := matcher.FindStringSubmatch(line)
			n, err := strconv.Atoi(m[2])
			if err != nil {
				fmt.Println("error parsing line", line)
				return
			}
			if m[1] == "x" {
				folds = append(folds, [2]int{n, 0})
			} else {
				folds = append(folds, [2]int{0, n})
			}
		}
	}

	for _, fold := range folds {
		if fold[0] != 0 {
			dots = foldLeft(dots, fold[0])
		} else {
			dots = foldUp(dots, fold[1])
		}
	}
	displayDots(dots)
}
