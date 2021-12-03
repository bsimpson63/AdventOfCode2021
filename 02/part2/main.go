package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	depth, horizontal, aim := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("bad input")
			return
		}
		if direction == "up" {
			aim -= num
		} else if direction == "down" {
			aim += num
		} else if direction == "forward" {
			horizontal += num
			depth += (aim * num)
		} else {
			fmt.Println("bad direction", direction)
			return
		}
	}
	fmt.Println("final position:", depth, horizontal)
	fmt.Println("product", depth*horizontal)
}
