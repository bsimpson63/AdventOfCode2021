package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("bad input")
			return
		}
		numbers = append(numbers, num)
	}
	prev := 0
	increaseCount := 0
	for i, n3 := range numbers {
		if i < 2 {
			continue
		}
		n1 := numbers[i-2]
		n2 := numbers[i-1]
		windowSum := n1 + n2 + n3
		if prev == 0 {
			prev = windowSum
			continue
		}
		if windowSum > prev {
			increaseCount++
		}
		prev = windowSum
	}
	fmt.Println("increase count is", increaseCount)
}
