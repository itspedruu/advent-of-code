package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func parseFile() [][]int {
	fileBytes, err := ioutil.ReadFile("10/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	grid := make([][]int, len(lines))

	for i, line := range lines {
		row := make([]int, len(line))

		for j, rawNumber := range line {
			row[j] = int(rawNumber - '0')
		}

		grid[i] = row
	}

	return grid
}

func solve(grid [][]int) (int, int) {
	visitedOut := 0
	out := 0

	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	for row, line := range grid {
		for col, curHeight := range line {
			if curHeight != 0 {
				continue
			}

			queue := make([][3]int, 4) // (row, col, searchHeight)

			queue[0] = [3]int{ row + 1, col, 1 }
			queue[1] = [3]int{ row - 1, col, 1 }
			queue[2] = [3]int{ row, col + 1, 1 }
			queue[3] = [3]int{ row, col - 1, 1 }

			visited := make(map[[2]int]bool)

			for len(queue) > 0 {
				searchRow := queue[0][0]
				searchCol := queue[0][1]
				searchHeight := queue[0][2]

				queue = queue[1:]

				if searchRow < 0 || searchRow > maxRow || searchCol < 0 || searchCol > maxCol || grid[searchRow][searchCol] != searchHeight {
					continue
				}

				if searchHeight == 9 {
					visitedKey := [2]int{ searchRow, searchCol }

					_, isVisited := visited[visitedKey]
					
					if !isVisited {
						visited[visitedKey] = true
						visitedOut++
					}

					out++
				} else {
					queue = append(
						queue,
						[3]int{ searchRow + 1, searchCol, searchHeight + 1 },
						[3]int{ searchRow - 1, searchCol, searchHeight + 1 },
						[3]int{ searchRow, searchCol + 1, searchHeight + 1},
						[3]int{ searchRow, searchCol - 1, searchHeight + 1},
					)
				}
			}
		}
	}

	return visitedOut, out
}

func main() {
	start := time.Now()

	grid := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(grid)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
