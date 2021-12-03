package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func convertToDecimal(arr []int) int {
	num := 0
	for i, d := range arr {
		if d != 1 {
			continue
		}
		exp := len(arr) - i - 1
		num += int(math.Pow(float64(2), float64(exp)))
	}
	return num
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	var counts []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(counts) == 0 {
			counts = make([]int, len(line))
		}

		for i, c := range line {
			if c == 48 {
				counts[i] -= 1
			} else {
				counts[i] += 1
			}
		}
	}
	gamma := make([]int, len(counts))
	epsilon := make([]int, len(counts))
	for i, c := range counts {
		if c > 0 {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}
	fmt.Println("gamma", gamma, convertToDecimal(gamma))
	fmt.Println("epsilon", epsilon, convertToDecimal(epsilon))
	fmt.Println("product", convertToDecimal(gamma)*convertToDecimal(epsilon))
}
