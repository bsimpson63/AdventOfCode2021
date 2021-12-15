package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
nextPolymer := make([]byte, 0)
		nextPolymer = append(nextPolymer, polymer[0])
		for i := 0; i < len(polymer)-1; i++ {
			f, s := polymer[i], polymer[i+1]
			if b, isOK := rules[[2]byte{f, s}]; isOK {
				nextPolymer = append(nextPolymer, b)
			}
			nextPolymer = append(nextPolymer, s)
		}
		polymer = nextPolymer
		fmt.Println(step)
*/

func takeStep(polymer *[]byte, rules map[[2]byte]byte) {
	i := 1
	for {
		if i >= len(*polymer) {
			break
		}
		f, s := (*polymer)[i-1], (*polymer)[i]
		if b, isOK := rules[[2]byte{f, s}]; isOK {
			*polymer = append((*polymer)[:i+1], (*polymer)[i:]...)
			(*polymer)[i] = b
			i++
		}
		i++
	}
}

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

	for step := 1; step <= 10; step++ {
		takeStep(&polymer, rules)
		fmt.Println(step)
	}
	counts := make(map[byte]int)
	for _, b := range polymer {
		counts[b]++
	}
	min, max := 999999999, -1
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
