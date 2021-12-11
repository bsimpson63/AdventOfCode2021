package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		output := parts[1]
		output_parts := strings.Split(output, " ")
		for _, p := range output_parts {
			if len(p) == 2 {
				// this is a 1
				c++
			} else if len(p) == 3 {
				// this is a 7
				c++
			} else if len(p) == 4 {
				// this is a 4
				c++
			} else if len(p) == 5 {
				// this is a 2 or 3 or 5
			} else if len(p) == 6 {
				// this is a 0 or 6 or 9
			} else if len(p) == 7 {
				// this is a 8
				c++
			}
		}
	}
	fmt.Println(c)

}
