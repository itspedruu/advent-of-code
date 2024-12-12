package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func parseFile() [][]int {
	fileBytes, err := ioutil.ReadFile("12/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	grid := make([][]int, len(lines))

	for row, line := range lines {
		grid[row] = make([]int, len(line))	

		for col, letter := range line {
			grid[row][col] = int(letter)
		}
	}

	return grid
}

func checkLetter(grid [][]int, row int, col int, letter int) bool {
	if row < 0 || row > len(grid) - 1 || col < 0 || col > len(grid[0]) - 1 {
		return false
	}

	return grid[row][col] == letter
}

func solve(grid [][]int) (int, int) {
	visited := make(map[[2]int]bool)

	solutionA := 0
	solutionB := 0
	
	areas := make(map[int]int)
	areaId := 1000

	areaGrid := make([][]int, len(grid))

	for row, line := range grid {
		areaGrid[row] = make([]int, len(grid[0]))

		for col, _ := range line {
			areaGrid[row][col] = -1
		}
	}

	for row, line := range grid {
		for col, letter := range line {
			_, isVisited := visited[[2]int{ row, col }]	

			if isVisited {
				continue
			}

			queue := make([][2]int, 1)
			queue[0] = [2]int{ row, col }

			areas[areaId] = 0
			perimeter := 0

			for len(queue) > 0 {
				curRow := queue[0][0]
				curCol := queue[0][1]

				queue = queue[1:]

				visitedKey := [2]int{ curRow, curCol }

				_, isVisited := visited[visitedKey]

				if isVisited {
					continue
				}

				visited[visitedKey] = true

				areaGrid[curRow][curCol] = areaId
				areas[areaId]++
				
				if checkLetter(grid, curRow + 1, curCol, letter) {
					queue = append(queue, [2]int{ curRow + 1, curCol })
				} else {
					perimeter++
				}

				if checkLetter(grid, curRow - 1, curCol, letter) {
					queue = append(queue, [2]int{ curRow - 1, curCol })
				} else {
					perimeter++
				}
				
				if checkLetter(grid, curRow, curCol + 1, letter) {
					queue = append(queue, [2]int{ curRow, curCol + 1 })
				} else {
					perimeter++
				}

				if checkLetter(grid, curRow, curCol - 1, letter) {
					queue = append(queue, [2]int{ curRow, curCol - 1 })
				} else {
					perimeter++
				}
			}

			solutionA += perimeter * areas[areaId]

			areaId++
		}
	}

	sides := make(map[int]int)

	for curAreaId := 1000; curAreaId < areaId; curAreaId++ {
		sides[curAreaId] = 0
	}

	for row, line := range areaGrid {
		checkingAreaId := line[0]

		if row == 0 || areaGrid[row - 1][0] != checkingAreaId {
			sides[checkingAreaId]++
		}

		for col, curAreaId := range line {
			if checkingAreaId == curAreaId {
				continue
			}

			if row == 0 || (areaGrid[row - 1][col - 1] == curAreaId && areaGrid[row - 1][col] == curAreaId) || areaGrid[row - 1][col] != curAreaId {
				sides[curAreaId]++
			}

			if row == 0 || (areaGrid[row - 1][col - 1] == checkingAreaId && areaGrid[row - 1][col] == checkingAreaId) || areaGrid[row - 1][col - 1] != checkingAreaId {
				sides[checkingAreaId]++
			}

			checkingAreaId = curAreaId
		}

		if row == 0 || areaGrid[row - 1][len(line) - 1] != checkingAreaId {
			sides[checkingAreaId]++
		}
	}

	for areaId, sides := range sides {
		solutionB += areas[areaId] * sides * 2
	}

	return solutionA, solutionB
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
