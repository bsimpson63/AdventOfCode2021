package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calcSingleFuel(delta int) int {
	f := 0
	for i := 1; i <= delta; i++ {
		f += i
	}
	return f
}

func calcFuel(crabs []int, pos int) int {
	f := 0
	for _, p := range crabs {
		if p > pos {
			f += calcSingleFuel(p - pos)
		} else if p < pos {
			f += calcSingleFuel(pos - p)
		}
	}
	return f
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crabs := make([]int, 0)
	s := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for _, p := range parts {
			f, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("bad input")
				return
			}
			crabs = append(crabs, f)
			s += f
		}
	}
	sort.Ints(crabs)

	minPos := 0
	minFuel := calcFuel(crabs, minPos)
	for pos := crabs[0]; pos <= crabs[len(crabs)-1]; pos++ {
		fuel := calcFuel(crabs, pos)
		if fuel < minFuel {
			minFuel = fuel
			minPos = pos
		}
	}
	fmt.Println("min fuel at", minPos, "is", minFuel)

}
