package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func parseFile() []string {
	fileBytes, err := os.ReadFile("07/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	return lines
}

func solve(grid []string) (int, int) {
	solutionA := 0
	startColumn := 0

	for i, char := range grid[0] {
		if char == 'S' {
			startColumn = i
			break
		}
	}

	beams := map[int]int{}
	beams[startColumn] = 1
	totalColumns := len(grid[0])

	for _, line := range grid {
		for beamColumn, timelines := range beams {
			if line[beamColumn] != '^' {
				continue	
			}

			solutionA++
			delete(beams, beamColumn)

			if beamColumn > 0 {
				_, ok := beams[beamColumn - 1]

				if ok {
					beams[beamColumn - 1] += timelines
				} else {
					beams[beamColumn - 1] = timelines
				}
			}

			if beamColumn < totalColumns - 1 {
				_, ok := beams[beamColumn + 1]

				if ok {
					beams[beamColumn + 1] += timelines
				} else {
					beams[beamColumn + 1] = timelines
				}
			}
		}
	}

	solutionB := 0

	for _, timelines := range beams {
		solutionB += timelines
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	grid := parseFile()
	solutionA, solutionB := solve(grid)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
