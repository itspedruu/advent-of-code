package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Position struct {
	row int
	col int
}

func getSearchingNeighbours(pos Position, rows int, cols int) []Position {
	out := make([]Position, 0)	

	if pos.row != rows - 1 {
		out = append(out, Position{ pos.row + 1, pos.col })

		if pos.col != 0 {
			out = append(out, Position{ pos.row + 1, pos.col - 1 })
		}

		if pos.col != cols - 1 {
			out = append(out, Position{ pos.row + 1, pos.col + 1 })
		}
	}

	if pos.col != cols - 1 {
		out = append(out, Position{ pos.row, pos.col + 1 })
	}

	return out
}

func parseFile() map[Position]map[Position]bool {
	fileBytes, err := os.ReadFile("04/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	graph := map[Position]map[Position]bool{}

	rows := len(lines)
	cols := len(lines[0])

	for row, line := range lines {
		for col, char := range line {
			if char != '@' {
				continue
			}

			curPos := Position{ row, col }
			searchingNeighbours := getSearchingNeighbours(curPos, rows, cols)
			_, ok := graph[curPos]

			if !ok {
				graph[curPos] = map[Position]bool{}
			}

			for _, neighbour := range searchingNeighbours {
				if lines[neighbour.row][neighbour.col] == '@' {
					_, ok := graph[neighbour]

					if !ok {
						graph[neighbour] = map[Position]bool{}
					}

					graph[curPos][neighbour] = true
					graph[neighbour][curPos] = true
				}
			}
		}
	}

	return graph
}

func solve(graph map[Position]map[Position]bool) (int, int) {
	solutionA := 0
	solutionB := 0
	flagSolutionA := true

	for true {
		queue := make([]Position, 0)

		for pos := range graph {
			if len(graph[pos]) < 4 {
				if flagSolutionA {
					solutionA++
				}

				solutionB++
				queue = append(queue, pos)
			}
		}

		if len(queue) == 0 {
			break
		}

		for _, pos := range queue {
			adjacents := graph[pos]

			for adjacent := range adjacents {
				delete(graph[adjacent], pos)
			}

			delete(graph, pos)
		}

		flagSolutionA = false
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	graph := parseFile()
	solutionA, solutionB := solve(graph)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A: ", solutionA)
	fmt.Println("B: ", solutionB)
}
