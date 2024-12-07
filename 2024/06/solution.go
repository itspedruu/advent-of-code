package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseFile() []string {
	fileBytes, err := ioutil.ReadFile("06/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	return lines
}

func solveA(grid []string) int {
	obstacles := make(map[[2]int]bool)

	guardPos := [2]int{ 0, 0 }

	for row, line := range grid {
		for column, char := range line {
			if char == '#' {
				pos := [2]int{ row, column }

				obstacles[pos] = true
			} else if char == '^' {
				guardPos[0] = row
				guardPos[1] = column
			}
		}
	}

	maxRow := len(grid) - 1
	maxColumn := len(grid[0]) - 1

	visited := make(map[[2]int]bool)
	out := 0

	dRow := -1
	dColumn := 0

	for guardPos[0] >= 0 && guardPos[0] <= maxRow && guardPos[1] >= 0 && guardPos[1] <= maxColumn {
		futurePos := [2]int{ guardPos[0] + dRow, guardPos[1] + dColumn }

		_, isVisited := visited[guardPos]
		_, isObstacle := obstacles[futurePos]

		if isObstacle {
			temp := dColumn

			dColumn = -dRow
			dRow = temp
		}

		if !isVisited {
			out++
			visited[guardPos] = true
		}

		guardPos[0] += dRow
		guardPos[1] += dColumn
	}

	return out
}

func doesItLoop(pos [2]int, dRow int, dColumn int, maxRow int, maxColumn int, obstaclesByRow map[int]map[int]bool, obstaclesByCol map[int]map[int]bool) bool {
	visited := make(map[[4]int]bool)

	done := false

	for !done {
		done = true

		key := [4]int{ pos[0], pos[1], dRow, dColumn }

		_, isVisited := visited[key]

		if isVisited {
			return true
		}

		visited[key] = true

		if dRow == 0 {
			for column := pos[1]; column >= 0 && column <= maxColumn; column += dColumn {
				exists := obstaclesByRow[pos[0]][column]

				if exists {
					pos[1] = column - dColumn

					temp := dColumn
					dColumn = -dRow
					dRow = temp

					done = false

					break
				}
			}
		} else {
			for row := pos[0]; row >= 0 && row <= maxRow; row += dRow {
				exists := obstaclesByCol[pos[1]][row]

				if exists {
					pos[0] = row - dRow

					temp := dColumn
					dColumn = -dRow
					dRow = temp
					
					done = false

					break
				}
			}
		}
	}

	return false
}

func solveB(grid []string) int {
	obstaclesByRow := make(map[int]map[int]bool)
	obstaclesByCol := make(map[int]map[int]bool)

	pos := [2]int{ 0, 0 }
	maxRow := len(grid) - 1
	maxColumn := len(grid[0]) - 1

	for row := 0; row <= maxRow; row++ {
		rowMap := make(map[int]bool)

		for col := 0; col <= maxColumn; col++ {
			rowMap[col] = false
		}

		obstaclesByRow[row] = rowMap;
	}

	for col:= 0; col <= maxColumn; col++ {
		colMap := make(map[int]bool)

		for row := 0; row <= maxRow; row++ {
			colMap[row] = false
		}

		obstaclesByCol[col] = colMap
	}

	for row, line := range grid {
		for column, char := range line {
			if char == '#' {
				obstaclesByRow[row][column] = true
				obstaclesByCol[column][row] = true
			} else if char == '^' {
				pos[0] = row
				pos[1] = column
			}
		}
	}

	dRow := -1
	dColumn := 0

	out := 0
	startPos := pos

	tested := make(map[[2]int]bool)

	for pos[0] >= 0 && pos[0] <= maxRow && pos[1] >= 0 && pos[1] <= maxColumn {
		futurePos := [2]int{ pos[0] + dRow, pos[1] + dColumn }

		isObstacle := obstaclesByRow[futurePos[0]][futurePos[1]]

		if isObstacle {
			temp := dColumn

			dColumn = -dRow
			dRow = temp

			continue
		} else if futurePos[0] >= 0 && futurePos[0] <= maxRow && futurePos[1] >= 0 && futurePos[1] <= maxColumn {
			obstaclesByRow[futurePos[0]][futurePos[1]] = true
			obstaclesByCol[futurePos[1]][futurePos[0]] = true

			_, hasBeenTested := tested[futurePos]

			if futurePos != startPos && !hasBeenTested && doesItLoop([2]int{ pos[0], pos[1] }, dColumn, -dRow, maxRow, maxColumn, obstaclesByRow, obstaclesByCol) {
				out++
			}

			tested[futurePos] = true

			obstaclesByRow[futurePos[0]][futurePos[1]] = false
			obstaclesByCol[futurePos[1]][futurePos[0]] = false
		}

		pos[0] += dRow
		pos[1] += dColumn
	}

	return out
}

func main() {
	grid := parseFile()

	solutionA := solveA(grid)
	solutionB := solveB(grid)

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
