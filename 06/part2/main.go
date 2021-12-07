package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCount(fishies map[int]int) int {
	s := 0
	for _, count := range fishies {
		s += count
	}
	return s
}

func takeStep(fishies map[int]int) map[int]int {
	nextFishies := make(map[int]int)
	for timer, count := range fishies {
		if timer == 0 {
			nextFishies[6] += count
			nextFishies[8] += count
		} else {
			nextFishies[timer-1] += count
		}
	}
	return nextFishies
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fishies := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for _, p := range parts {
			f, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("bad input")
				return
			}
			fishies[f]++
		}
	}
	for i := 0; i <= 256; i++ {
		fmt.Println(i, getCount(fishies))
		fishies = takeStep(fishies)
	}
}
