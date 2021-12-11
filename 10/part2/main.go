package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func discardCorrupt(lines []string) []string {
	openingChar := map[string]string{
		">": "<",
		")": "(",
		"]": "[",
		"}": "{",
	}

	remaining := make([]string, 0)

	for _, line := range lines {
		isCorrupt := false
		openStack := make([]string, 0)
		for i := 0; i < len(line); i++ {
			c := string(line[i])

			if c == "<" || c == "{" || c == "(" || c == "[" {
				openStack = append(openStack, c)
				continue
			}

			expected := openingChar[c]

			var actual string
			if len(openStack) > 0 {
				actual = openStack[len(openStack)-1]
			} else {
				actual = "?"
			}

			if expected != actual {
				isCorrupt = true
				break
			}

			n := len(openStack) - 1
			openStack = append(openStack[:n], openStack[n+1:]...)
		}
		if !isCorrupt {
			remaining = append(remaining, line)
		}
	}
	return remaining
}

func scoreIncomplete(line string) int {
	openingChar := map[string]string{
		">": "<",
		")": "(",
		"]": "[",
		"}": "{",
	}
	openStack := make([]string, 0)
	for i := 0; i < len(line); i++ {
		c := string(line[i])

		if c == "<" || c == "{" || c == "(" || c == "[" {
			openStack = append(openStack, c)
			continue
		}

		expected := openingChar[c]

		var actual string
		if len(openStack) > 0 {
			actual = openStack[len(openStack)-1]
		} else {
			actual = "?"
		}

		if expected != actual {
			// this line is corrupt
			fmt.Println(line, "is corrupt:")
			return -1
		}

		n := len(openStack) - 1
		openStack = append(openStack[:n], openStack[n+1:]...)
	}

	scores := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
	s := 0
	for i := len(openStack) - 1; i >= 0; i-- {
		c := openStack[i]
		s *= 5
		s += scores[c]
	}
	return s
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	lines = discardCorrupt(lines)
	scores := make([]int, 0)
	for _, line := range lines {
		scores = append(scores, scoreIncomplete(line))
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
