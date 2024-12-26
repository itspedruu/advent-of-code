package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Pos struct {
	row int
	col int
}

func parseFile() ([]string, int, Pos, Pos) {
	fileBytes, err := ioutil.ReadFile("20/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")

	var startPos Pos
	var endPos Pos

	tracks := 0

	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				continue
			}

			if char == 'S' {
				startPos.row = row
				startPos.col = col
			} else if char == 'E' {
				endPos.row = row
				endPos.col = col
			}

			tracks++
		}
	}

	return lines, tracks, startPos, endPos
}

func abs(n int) int {
	if (n >= 0) {
		return n
	}
	
	return -n
}

func distance(a Pos, b Pos) int {
	return abs(a.row - b.row) + abs(a.col - b.col)
}

func solve(grid []string, tracks int, startPos Pos, endPos Pos) (int, int) {
	curPos := startPos
	curDistance := 0

	directions := [4][2]int{
		[2]int{ 1, 0 },
		[2]int{ -1, 0 },
		[2]int{ 0, 1 },
		[2]int{ 0, -1 },
	}

	visited := make(map[Pos]bool)
	path := make([]Pos, tracks)

	for curPos != endPos {
		visited[curPos] = true
		path[curDistance] = curPos

		for _, direction := range directions {
			neighbour := Pos{ curPos.row + direction[0], curPos.col + direction[1] }

			if grid[neighbour.row][neighbour.col] == '#' || visited[neighbour] {
				continue
			}

			curPos = neighbour
			break
		}

		curDistance++
	}

	endDistance := curDistance
	path[curDistance] = endPos

	solutionA := 0
	solutionB := 0

	for curDistance := 0; curDistance < endDistance - 102; curDistance++ {
		curPos := path[curDistance]

		for futureDistance := curDistance + 102; futureDistance <= endDistance; futureDistance++ {
			radius := distance(path[futureDistance], curPos)

			if radius >= 2 && radius <= 20 {
				diff := futureDistance - curDistance - radius

				if diff >= 100 {
					if radius == 2 {
						solutionA++
					}

					solutionB++
				}
			}
		}
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	grid, tracks, startPos, endPos := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(grid, tracks, startPos, endPos)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
