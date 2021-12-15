package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// I couldn't figure this one out, found solution at
// https://www.reddit.com/r/adventofcode/comments/rfzq6f/comment/hohc8vc/?utm_source=reddit&utm_medium=web2x&context=3
// key observation is that ordering doesn't matter, only the counts of pairs
// matters

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readingTemplate := true
	polymer := make([]byte, 0)
	rules := make(map[[2]byte]byte)
	for scanner.Scan() {
		line := scanner.Text()
		if readingTemplate {
			for i := 0; i < len(line); i++ {
				b := line[i]
				polymer = append(polymer, b)
			}
			readingTemplate = false
			continue
		}

		if line == "" {
			continue
		}

		parts := strings.Split(line, " -> ")
		f := parts[0][0]
		s := parts[0][1]
		rules[[2]byte{f, s}] = parts[1][0]
	}

	// position doesn't matter, only pairs
	pairCounts := make(map[[2]byte]int)
	for i := 0; i < len(polymer)-1; i++ {
		pair := [2]byte{polymer[i], polymer[i+1]}
		pairCounts[pair]++
	}

	counts := make(map[byte]int)
	for i := 0; i < len(polymer); i++ {
		counts[polymer[i]]++
	}

	for step := 1; step <= 40; step++ {
		nextPairCounts := make(map[[2]byte]int)
		for pair, count := range pairCounts {
			m := rules[pair]
			leftPair := [2]byte{pair[0], m}
			rightPair := [2]byte{m, pair[1]}
			nextPairCounts[leftPair] += count
			nextPairCounts[rightPair] += count
			counts[m] += count
		}
		pairCounts = nextPairCounts
	}

	min, max := 2191039569603, -1
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	fmt.Println(max - min)
}
