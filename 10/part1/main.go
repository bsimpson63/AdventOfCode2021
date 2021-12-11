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

	openingChar := map[string]string{
		">": "<",
		")": "(",
		"]": "[",
		"}": "{",
	}

	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()

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
				// this line is corrupt
				fmt.Println(line, "is corrupt:", scores[c])
				score += scores[c]
				isCorrupt = true
				break
			}

			n := len(openStack) - 1
			openStack = append(openStack[:n], openStack[n+1:]...)
		}
		if !isCorrupt {
			fmt.Println(line, "not corrupt")
		}
	}
	fmt.Println("score:", score)
}
