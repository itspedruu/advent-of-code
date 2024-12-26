package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

const MAX_INT = int(^uint(0) >> 1)
const SIZE = 71

type Pos struct {
	row int
	col int
}

func parseFile() []Pos {
	fileBytes, err := ioutil.ReadFile("18/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")

	corrupted := make([]Pos, len(lines))

	for i, line := range lines {
		rawNumbers := strings.Split(line, ",")
		row, _ := strconv.Atoi(rawNumbers[0])
		col, _ := strconv.Atoi(rawNumbers[1])

		corrupted[i] = Pos{ row, col }
	}

	return corrupted
}

func buildGraph(grid [SIZE][SIZE]bool) map[Pos]map[Pos]int {
	graph := make(map[Pos]map[Pos]int)

	for row, line := range grid {
		for col, isCorrupted := range line {
			if isCorrupted {
				continue
			}

			curPos := Pos{ row, col }

			neighbours := [4]Pos{
				Pos{ row + 1, col },	
				Pos{ row - 1, col },
				Pos{ row, col + 1 },
				Pos{ row, col - 1 },
			}

			for _, neighbour := range neighbours {
				if neighbour.row < 0 || neighbour.row >= SIZE || neighbour.col < 0 || neighbour.col >= SIZE || grid[neighbour.row][neighbour.col] {
					continue
				}

				if graph[curPos] == nil {
					graph[curPos] = make(map[Pos]int)
				}

				graph[curPos][neighbour] = 1
			}
		}
	}

	return graph
}

func dijkstra(graph map[Pos]map[Pos]int) map[Pos]int {
	queue := make([]Pos, 1)
	distances := make(map[Pos]int)

	queue[0] = Pos{ 0, 0 }

	for len(queue) > 0 {
		index := 0
		minDist := MAX_INT

		for i, key := range queue {
			if distances[key] < minDist {
				minDist = distances[key]
				index = i
			}
		}

		node := queue[index]
		queue[index] = queue[0]
		queue = queue[1:]
		
		for adj, weight := range graph[node] {
			_, ok := distances[adj]

			if !ok {
				queue = append(queue, adj)
				distances[adj] = MAX_INT
			}

			if minDist + weight < distances[adj] {
				distances[adj] = minDist + weight
			}
		}
	}

	return distances
}

func solve(corrupted []Pos) (int, Pos) {
	var grid [SIZE][SIZE]bool

	for i := 0; i < 1024; i++ {
		corruptedPos := corrupted[i]

		grid[corruptedPos.row][corruptedPos.col] = true
	}

	graph := buildGraph(grid)
	distances := dijkstra(graph)
	endPos := Pos{ SIZE - 1, SIZE - 1}
	solutionA := distances[endPos]

	i := 1024
	j := len(corrupted) - 1

	for i < j {
		middle := i + ((j - i) >> 1)

		var grid [SIZE][SIZE]bool

		for i := 0; i < middle; i++ {
			corruptedPos := corrupted[i]

			grid[corruptedPos.row][corruptedPos.col] = true
		}

		graph = buildGraph(grid)
		distances = dijkstra(graph)

		if distances[endPos] != 0 {
			i = middle + 1
		} else {
			j = middle - 1
		}
	}

	solutionB := corrupted[i - 1]

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	corrupted := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(corrupted)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
