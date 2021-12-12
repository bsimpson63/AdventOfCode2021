package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isUpper(s string) bool {
	for _, r := range s {
		if unicode.ToUpper(r) != r {
			return false
		}
	}
	return true
}

func hasVisitedSmallTwice(path string) bool {
	visitCount := make(map[string]int)
	parts := strings.Split(path, ",")
	for _, p := range parts {
		if isUpper(p) {
			continue
		}
		if visitCount[p] >= 1 {
			return true
		}
		visitCount[p]++
	}
	return false
}

func findPaths(routes map[string]map[string]bool) {
	// find the number of distinct paths that start at start,
	// end at end, and don't visit small caves more than once
	// except a single small cave that can be visited twice!
	paths := make([]string, 0)
	stack := make([]string, 0)
	seen := make(map[string]bool)

	stack = append(stack, "start")

	for {
		if len(stack) == 0 {
			break
		}

		n := len(stack) - 1
		partialPath := stack[n]
		parts := strings.Split(partialPath, ",")
		lastPos := parts[len(parts)-1]
		stack = append(stack[:n], stack[n+1:]...)

		for d := range routes[lastPos] {
			if d == "start" {
				continue
			}

			if d == "end" {
				path := partialPath + "," + d
				paths = append(paths, path)
				seen[path] = true
				continue
			}

			if !isUpper(d) {
				if strings.Contains(partialPath, d) {
					// check if we've been any small cave twice already
					if hasVisitedSmallTwice(partialPath) {
						continue
					}
				}
			}

			path := partialPath + "," + d
			if alreadySeen := seen[path]; alreadySeen {
				continue
			}

			stack = append(stack, path)
			seen[path] = true
		}
	}
	fmt.Println("Found", len(paths), "paths")
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	routes := make(map[string]map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")

		if routes[parts[0]] == nil {
			routes[parts[0]] = make(map[string]bool)
		}
		routes[parts[0]][parts[1]] = true

		if routes[parts[1]] == nil {
			routes[parts[1]] = make(map[string]bool)
		}
		routes[parts[1]][parts[0]] = true
	}
	findPaths(routes)
}
