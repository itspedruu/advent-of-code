package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const XMAS = "XMAS"

func parseFile() []string {
	fileBytes, err := ioutil.ReadFile("04/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	return lines
}

func solveA(lines []string) int {
	queue := make([][5]int, 0)

	for row, line := range lines {
		for column, letter := range line {
			if letter == 'X' {
				north := [5]int{ 0, -1, 0, row, column }
				south := [5]int{ 0, 1, 0, row, column }
				east := [5]int{ 0, 0, 1, row, column }
				west := [5]int{ 0, 0, -1, row, column }
				northeast := [5]int{ 0, -1, 1, row, column }
				northwest := [5]int{ 0, -1, -1, row, column }
				southeast := [5]int{ 0, 1, 1, row, column }
				southwest := [5]int{ 0, 1, -1, row, column }

				queue = append(queue, north, south, east, west, northeast, northwest, southeast, southwest)
			}
		}
	}

	count := 0
	maxRow := len(lines) - 1
	maxColumn := len(lines[0]) - 1

	for len(queue) > 0 {
		searchIndex := queue[0][0]
		rowSpeed := queue[0][1]
		columnSpeed := queue[0][2]
		row := queue[0][3]
		column := queue[0][4]

		queue = queue[1:]

		newRow := row + rowSpeed
		newColumn := column + columnSpeed

		if newRow > maxRow || newColumn > maxColumn || newRow < 0 || newColumn < 0 {
			continue
		}

		if lines[newRow][newColumn] == XMAS[searchIndex + 1] {
			if searchIndex + 1 == 3 {
				count++
			} else {
				newNode := [5]int{ searchIndex + 1, rowSpeed, columnSpeed, newRow, newColumn }

				queue = append(queue, newNode)
			}
		}
	}

	return count
}

func solveB(lines []string) int {
	queue := make([][5]int, 0)

	for row, line := range lines {
		for column, letter := range line {
			if letter == 'M' {
				northeast := [5]int{ 1, -1, 1, row, column }
				northwest := [5]int{ 1, -1, -1, row, column }
				southeast := [5]int{ 1, 1, 1, row, column }
				southwest := [5]int{ 1, 1, -1, row, column }

				queue = append(queue, northeast, northwest, southeast, southwest)
			}
		}
	}

	ends := make(map[[2]int][][2]int)
	maxRow := len(lines) - 1
	maxColumn := len(lines[0]) - 1

	for len(queue) > 0 {
		searchIndex := queue[0][0]
		rowSpeed := queue[0][1]
		columnSpeed := queue[0][2]
		row := queue[0][3]
		column := queue[0][4]

		queue = queue[1:]

		newRow := row + rowSpeed
		newColumn := column + columnSpeed

		if newRow > maxRow || newColumn > maxColumn || newRow < 0 || newColumn < 0 {
			continue
		}

		if lines[newRow][newColumn] == XMAS[searchIndex + 1] {
			if searchIndex + 1 == 3 {
				initialPosition := [2]int{ row - rowSpeed, column - columnSpeed }
				endPosition := [2]int{ newRow, newColumn }

				curEnds, ok := ends[initialPosition]

				if ok {
					ends[initialPosition] = append(curEnds, endPosition)
				} else {
					newEnds := make([][2]int, 1)
					newEnds[0] = endPosition

					ends[initialPosition] = newEnds
				}
			} else {
				newNode := [5]int{ searchIndex + 1, rowSpeed, columnSpeed, newRow, newColumn }

				queue = append(queue, newNode)
			}
		}
	}

	count := 0
	visited := make(map[[2]int]bool)

	for initialPosition, endPositions := range ends {
		_, isVisited := visited[initialPosition]

		if isVisited {
			continue
		}

		visited[initialPosition] = true

		for _, endPosition := range endPositions {
			possibleReverseInitials := [2][2]int{
				[2]int{ endPosition[0], initialPosition[1] },
				[2]int{ initialPosition[0], endPosition[1] },
			}

			possibleReverseEnds := [2][2]int{
				[2]int{ initialPosition[0], endPosition[1] },
				[2]int{ endPosition[0], initialPosition[1] },
			}

			for i, possibleReverseInitial := range possibleReverseInitials {
				if visited[possibleReverseInitial] {
					continue
				}

				reverseEnds, ok := ends[possibleReverseInitial]

				found := false

				if ok {
					for _, reverseEnd := range reverseEnds {
						if reverseEnd == possibleReverseEnds[i] {
							found = true
							break
						}
					}
				}

				if found {
					count++
					break
				}
			}
		}
	}

	return count
}

func main() {
	lines := parseFile()

	solutionA := solveA(lines)
	solutionB := solveB(lines)

	fmt.Printf("A: %d\n", solutionA)
	fmt.Printf("B: %d\n", solutionB)
}
