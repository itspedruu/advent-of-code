package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func parseFile() ([][5]int, [][5]int) {
	fileBytes, err := ioutil.ReadFile("25/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n\n")

	keys := make([][5]int, 0)
	locks := make([][5]int, 0)
	
	for _, combination := range lines {
		grid := strings.Split(combination, "\n")
		isKey := grid[0][0] == '.'

		i := 0
		maxI := len(grid) - 1
		delta := 1

		if isKey {
			i = len(grid) - 1
			maxI = 0
			delta = -1
		}

		heights := [5]int{ -1, -1, -1, -1, -1 }

		for ; i * delta <= maxI * delta; i += delta {
			flag := false

			for j, char := range grid[i] {
				if char == '#' {
					heights[j]++
					flag = true
				}
			}

			if !flag {
				break
			}
		}

		if isKey {
			keys = append(keys, heights)
		} else {
			locks = append(locks, heights)
		}
	}

	return keys, locks
}

// returns true if a <= b
func compareCombinations(a [5]int, b [5]int) bool {
	for i := range a {
		if a[i] > b[i] {
			return false
		}
	}

	return true
}

func solve(keys [][5]int, locks [][5]int) int {
	maxHeight := 5
	out := 0

	for _, lock := range locks {
		var maxCombination [5]int

		for i, n := range lock {
			maxCombination[i] = maxHeight - n
		}

		for _, key := range keys {
			if compareCombinations(key, maxCombination) {
				out++
			}
		}
	}

	return out
}

func main() {
	start := time.Now()

	keys, locks := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA := solve(keys, locks)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
}
