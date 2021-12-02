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
	prev := -1
	increaseCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("bad input")
			return
		}
		if prev == -1 {
			prev = num
			continue
		}
		if num > prev {
			increaseCount++
		}
		prev = num
	}
	fmt.Println("increase count is", increaseCount)
}
