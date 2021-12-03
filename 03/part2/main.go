package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func convertToDecimal(s string) int {
	num := 0
	for i, d := range s {
		if d != 49 {
			continue
		}
		exp := len(s) - i - 1
		num += int(math.Pow(float64(2), float64(exp)))
	}
	return num
}

func mostCommonAtPos(lines []string, pos int) int {
	count := 0
	for _, line := range lines {
		if line[pos] == 48 {
			count -= 1
		} else {
			count += 1
		}
	}

	if count >= 0 {
		return 1
	} else {
		return 0
	}
}

func filterLines(lines []string, pos int, val int) []string {
	var filtered []string
	for _, line := range lines {
		if (line[pos] == 48 && val == 0) || (line[pos] == 49 && val == 1) {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	oxyLines := make([]string, len(lines))
	co2Lines := make([]string, len(lines))

	copy(oxyLines, lines)
	copy(co2Lines, lines)

	pos := 0
	for {
		if len(oxyLines) <= 1 {
			break
		}

		val := mostCommonAtPos(oxyLines, pos)
		oxyLines = filterLines(oxyLines, pos, val)
		pos++
	}

	pos = 0
	for {
		if len(co2Lines) <= 1 {
			break
		}

		val := mostCommonAtPos(co2Lines, pos)
		if val == 0 {
			val = 1
		} else {
			val = 0
		}
		co2Lines = filterLines(co2Lines, pos, val)
		pos++
	}

	fmt.Println("oxygen", oxyLines, convertToDecimal(oxyLines[0]))
	fmt.Println("co2", co2Lines, convertToDecimal(co2Lines[0]))
	fmt.Println("product", convertToDecimal(oxyLines[0])*convertToDecimal(co2Lines[0]))
}
