package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const MAX_INT = int(^uint(0) >> 1)

func parseFile() ([]string, [2]int, [2]int) {
	fileBytes, err := ioutil.ReadFile("16/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	var startPos [2]int
	var endPos [2]int

	for row, line := range lines {
		for col, char := range line {
			if char == 'S' {
				startPos[0] = row
				startPos[1] = col
			} else if char == 'E' {
				endPos[0] = row
				endPos[1] = col

			}
		}
	}

	return lines, startPos, endPos
}

type Node struct {
	row int
	col int
	vrow int
	vcol int
}

func buildGraph(grid []string, corners map[[2]int]bool) map[Node]map[Node]int {
	graph := make(map[Node]map[Node]int)

	for row, line := range grid {
		for col, char := range line {
			if char == '#' {
				continue
			}

			potentialNeighbours := [4][2]int{
				[2]int{ row + 1, col },
				[2]int{ row - 1, col },
				[2]int{ row, col + 1 },
				[2]int{ row, col - 1 },
			}

			neighbours := make([][2]int, 0)

			for _, potentialNeighbour := range potentialNeighbours {
				if grid[potentialNeighbour[0]][potentialNeighbour[1]] == '#' {
					continue
				}

				neighbours = append(neighbours, potentialNeighbour)
			}

			isCorner := false
			patientZero := neighbours[0]

			for _, neighbour := range neighbours {
				if patientZero[0] != neighbour[0] && patientZero[1] != neighbour[1] {
					corners[[2]int{ row, col }] = true
					isCorner = true
				}

				from := Node{ row, col, neighbour[0] - row, neighbour[1] - col }
				to := Node{ neighbour[0], neighbour[1], neighbour[0] - row, neighbour[1] - col }

				graph[from] = make(map[Node]int)
				graph[from][to] = 1
			}

			if isCorner {
				for i, neighbour := range neighbours {
					from := Node{ row, col, neighbour[0] - row, neighbour[1] - col }

					for j, toNeighbour := range neighbours {
						if i == j {
							continue
						}

						to := Node{ row, col, row - toNeighbour[0], col - toNeighbour[1] }

						if from == to {
							continue
						}

						if graph[from] == nil {
							graph[from] = make(map[Node]int)
						}

						if graph[to] == nil {
							graph[to] = make(map[Node]int)
						}

						graph[from][to] = 1000
						graph[to][from] = 1000
					}
				}
			}
		}
	}

	return graph
}

func dijkstra(graph map[Node]map[Node]int, startPos [2]int) (map[Node]int, map[Node]Node) {
	queue := make([]Node, 1)
	distances := make(map[Node]int)
	parents := make(map[Node]Node)

	distances[Node{ startPos[0], startPos[1], 0, -1 }] = 0
	distances[Node{ startPos[0], startPos[1], -1, 0 }] = 1000

	queue[0] = Node{ startPos[0], startPos[1], -1, 0 }

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
				parents[adj] = node
				distances[adj] = minDist + weight
			}
		}
	}

	return distances, parents
}

func solve(grid []string, startPos [2]int, endPos [2]int) (int, int) {
	corners := make(map[[2]int]bool)
	graph := buildGraph(grid, corners)
	distances, parents := dijkstra(graph, startPos)

	startNode := Node{ startPos[0], startPos[1], -1, 0 }
	endNode := Node{ endPos[0], endPos[1], -1, 0 }

	visited := make(map[Node]bool)
	visited[startNode] = true

	uniqueTiles := make(map[[2]int]bool)
	uniqueTiles[startPos] = true

	queue := make([]Node, 1)
	queue[0] = endNode

	for len(queue) > 0 {
		prev := queue[0]
		cur := queue[0]
		queue = queue[1:]

		for cur != startNode {
			if visited[cur] {
				break
			}

			visited[cur] = true

			pos := [2]int{ cur.row, cur.col }

			uniqueTiles[pos] = true
			isCorner := corners[pos]
			
			next := parents[cur]

			if isCorner {
				if next.row == cur.row && next.col == cur.col {
					next = parents[next]
				}

				next := parents[parents[cur]]

				neighbours := [4][2]int{
					[2]int{ cur.row + 1, cur.col },
					[2]int{ cur.row - 1, cur.col },
					[2]int{ cur.row, cur.col + 1 },
					[2]int{ cur.row, cur.col - 1 },
				}

				checkDirections := make([][2]int, 0)

				for _, neighbour := range neighbours {
					if grid[neighbour[0]][neighbour[1]] == '#' || [2]int{ prev.row, prev.col } == neighbour || [2]int{ next.row, next.col } == neighbour {
						continue
					}

					checkDirections = append(checkDirections, [2]int{ cur.row - neighbour[0], cur.col - neighbour[1] })
				}

				for _, direction := range checkDirections {
					scoreDiff := 1002
					
					if direction == [2]int{ cur.vrow, cur.vcol } {
						scoreDiff = 2
					}

					offNeighbour := Node{ cur.row - direction[0], cur.col - direction[1], direction[0], direction[1] }

					if distances[offNeighbour] + scoreDiff == distances[prev] {
						queue = append(queue, offNeighbour)
					}
				}
			}

			prev = cur
			cur = next
		}
	}

	return distances[endNode], len(uniqueTiles)
}

func main() {
	start := time.Now()

	grid, startPos, endPos := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(grid, startPos, endPos)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
