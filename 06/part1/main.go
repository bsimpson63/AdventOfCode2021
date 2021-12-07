package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func takeStep(fishies []int) []int {
	newFishies := make([]int, 0)
	for _, fish := range fishies {
		if fish == 0 {
			newFishies = append(newFishies, 8)
			newFishies = append(newFishies, 6)
		} else {
			newFishies = append(newFishies, fish-1)
		}
	}
	return newFishies
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fishies := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for _, p := range parts {
			f, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("bad input")
				return
			}
			fishies = append(fishies, f)
		}
	}
	for i := 0; i <= 80; i++ {
		fmt.Println(i, len(fishies))
		fishies = takeStep(fishies)
	}
}
