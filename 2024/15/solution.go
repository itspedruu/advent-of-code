package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const WALL = -1
const EMPTY = 0
const BOX = 1
const BOX_OPEN = 2
const BOX_CLOSE = 3

func parseFile() ([][]int, [2]int, [][2]int) {
	fileBytes, err := ioutil.ReadFile("15/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	grid := make([][]int, len(lines))

	var startPos [2]int

	flag := false

	movements := make([][2]int, 0)

	for row, line := range lines {
		if line == "" {
			flag = true
			continue
		}

		if !flag {
			grid[row] = make([]int, len(line))
		}

		for col, char := range line {
			if flag {
				var movement [2]int

				if char == '^' {
					movement[0] = -1
					movement[1] = 0
				} else if char == 'v' {
					movement[0] = 1
					movement[1] = 0
				} else if char == '>' {
					movement[0] = 0
					movement[1] = 1
				} else {
					movement[0] = 0
					movement[1] = -1
				}

				movements = append(movements, movement)
			} else {
				if char == '@' {
					startPos[0] = row
					startPos[1] = col

					grid[row][col] = EMPTY
				} else if char == '#' {
					grid[row][col] = WALL
				} else if char == '.' {
					grid[row][col] = EMPTY
				} else {
					grid[row][col] = BOX
				}
			}
		}
	}

	return grid, startPos, movements
}

func copyGrid(grid [][]int) [][]int {
	clone := make([][]int, len(grid))

	for row, line := range grid {
		clone[row] = make([]int, len(line))

		for col, entity := range line {
			clone[row][col] = entity
		}
	}

	return clone
}

func solveA(grid [][]int, pos [2]int, movements [][2]int) int {
	for _, movement := range movements {
		newRow := pos[0] + movement[0]
		newCol := pos[1] + movement[1]

		if grid[newRow][newCol] == WALL {
			continue
		}

		if grid[newRow][newCol] == BOX {
			row := newRow + movement[0]
			col := newCol + movement[1]
			done := false
			
			for !done {
				done = false

				if grid[row][col] == WALL {
					done = true
				} else if grid[row][col] != BOX {
					grid[row][col] = BOX
					grid[newRow][newCol] = EMPTY

					pos[0] = newRow
					pos[1] = newCol

					done = true
				}

				row += movement[0]
				col += movement[1]
			}
		} else {
			pos[0] = newRow
			pos[1] = newCol
		}
	}

	out := 0

	for row, line := range grid {
		for col, entity := range line {
			if entity == BOX {
				out += row * 100 + col
			}
		}
	}

	return out
}

func getBoxGroup(grid [][]int, row int, col int) [2][2]int {
	var boxGroup [2][2]int

	if grid[row][col] == BOX_OPEN {
		boxGroup[0] = [2]int{ row, col }
		boxGroup[1] = [2]int{ row, col + 1 }
	} else {
		boxGroup[0] = [2]int{ row, col - 1 }
		boxGroup[1] = [2]int{ row, col }
	}

	return boxGroup
}

func solveB(grid [][]int, pos [2]int, movements [][2]int) int {
	newGrid := make([][]int, len(grid))

	for row, line := range grid {
		newGrid[row] = make([]int, len(line) * 2)

		for col, entity := range line {
			if entity == BOX {
				newGrid[row][col * 2] = BOX_OPEN
				newGrid[row][col * 2 + 1] = BOX_CLOSE
			} else {
				newGrid[row][col * 2] = entity	
				newGrid[row][col * 2 + 1] = entity
			}
		}
	}

	grid = newGrid
	pos[1] *= 2

	for _, movement := range movements {
		newRow := pos[0] + movement[0]
		newCol := pos[1] + movement[1]

		if grid[newRow][newCol] == WALL {
			continue
		}

		if grid[newRow][newCol] == BOX_OPEN || grid[newRow][newCol] == BOX_CLOSE {
			row := newRow + movement[0]
			col := newCol + movement[1] * 2
			done := false

			if movement[0] == 0 {
				for !done {
					done = false

					if grid[row][col] == WALL {
						done = true
					} else if grid[row][col] == EMPTY {
						for swapCol := col; swapCol * movement[1] >= newCol * movement[1]; swapCol -= movement[1] {
							grid[row][swapCol] = grid[row][swapCol - movement[1]]
							grid[row][swapCol - movement[1]] = EMPTY
						}

						pos[1] = newCol

						done = true
					}

					col += movement[1] * 2
				}
			} else {
				boxes := make([][2][2]int, 1)
				lastBoxes := make([][2][2]int, 1)

				boxes[0] = getBoxGroup(grid, newRow, newCol)
				lastBoxes[0] = getBoxGroup(grid, newRow, newCol)

				canMove := true

				for len(lastBoxes) > 0 {
					newLastBoxes := make([][2][2]int, 0)

					for _, boxGroup := range lastBoxes {
						openCol := boxGroup[0][1]
						closeCol := boxGroup[1][1]

						if grid[row][openCol] == WALL || grid[row][closeCol] == WALL {
							canMove = false
							break
						} else if grid[row][openCol] == EMPTY && grid[row][closeCol] == EMPTY {
							continue
						} else {
							if grid[row][openCol] != EMPTY {
								newLastBoxes = append(newLastBoxes, getBoxGroup(grid, row, openCol))

								if grid[row][openCol] == BOX_OPEN {
									continue
								}
							}

							if grid[row][closeCol] != EMPTY {
								newLastBoxes = append(newLastBoxes, getBoxGroup(grid, row, closeCol))
							}
						}
					}

					if !canMove {
						break
					}

					lastBoxes = newLastBoxes
					boxes = append(boxes, newLastBoxes...)

					row += movement[0]
				}

				if canMove {
					for i := len(boxes) - 1; i >= 0; i-- {
						row := boxes[i][0][0]
						openCol := boxes[i][0][1]
						closeCol := boxes[i][1][1]

						grid[row][openCol] = EMPTY
						grid[row][closeCol] = EMPTY
						grid[row + movement[0]][openCol] = BOX_OPEN
						grid[row + movement[0]][closeCol] = BOX_CLOSE
					}

					pos[0] = newRow
				}
			}
		} else {
			pos[0] = newRow
			pos[1] = newCol
		}
	}

	out := 0

	for row, line := range grid {
		for col, entity := range line {
			if entity == BOX_OPEN {
				out += row * 100 + col
			}
		}
	}

	return out
}

func main() {
	start := time.Now()

	grid, startPos, movements := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA := solveA(copyGrid(grid), startPos, movements)
	fmt.Println("~ solving A:", time.Since(start))

	solutionB := solveB(copyGrid(grid), startPos, movements)
	fmt.Println("~ solving B:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
