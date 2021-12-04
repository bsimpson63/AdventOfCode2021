package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board [10][5]int

func makeBoard(raw []int) Board {
	b := Board{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b[i][j] = raw[5*i+j]
			b[5+j][i] = raw[5*i+j]
		}
	}
	return b
}

func checkBoard(board Board, called []int) bool {
	calledMap := make(map[int]bool)
	for _, c := range called {
		calledMap[c] = true
	}

	for _, row := range board {
		rowIsValid := true
		for _, val := range row {
			isCalled := calledMap[val]
			if !isCalled {
				// this row has a number that has not been called
				rowIsValid = false
				break
			}
		}
		// this row has all its numbers called
		if rowIsValid {
			return true
		}
	}
	return false
}

func scoreBoard(board Board, called []int) int {
	calledMap := make(map[int]bool)
	for _, c := range called {
		calledMap[c] = true
	}

	s := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			v := board[i][j]
			if isCalled := calledMap[v]; !isCalled {
				s += v
			}
		}
	}
	return s * called[len(called)-1]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	order := make([]int, 0)
	boards := make([]Board, 0)
	raw := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if len(order) == 0 {
			parts := strings.Split(line, ",")
			for _, p := range parts {
				num, err := strconv.Atoi(p)
				if err != nil {
					fmt.Println("bad input parsing order", line)
					fmt.Println("bad part", p)
					return
				}
				order = append(order, num)
			}
			continue
		}

		parts := strings.Fields(line)
		for _, p := range parts {
			num, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("bad input parsing board", line)
				fmt.Println("bad part", p)
				return
			}
			raw = append(raw, num)
		}

		if len(raw) == 25 {
			b := makeBoard(raw)
			boards = append(boards, b)
			raw = make([]int, 0)
		}
	}
	for i := 5; i < len(order); i++ {
		if len(boards) == 0 {
			break
		}

		toRemove := make([]int, 0)
		for boardIndex, board := range boards {
			if checkBoard(board, order[0:i]) {
				fmt.Println("called", order[0:i])
				fmt.Println("at turn", i, "board", board, "is a winner")
				fmt.Println("score", scoreBoard(board, order[0:i]))
				toRemove = append(toRemove, boardIndex)
			}
		}

		for j := len(toRemove) - 1; j >= 0; j-- {
			boardIndex := toRemove[j]
			boards = append(boards[:boardIndex], boards[boardIndex+1:]...)
		}
	}
}
