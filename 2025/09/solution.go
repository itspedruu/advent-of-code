package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	x int
	y int
}

func parseFile() []*Position {
	fileBytes, err := os.ReadFile("09/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	positions := make([]*Position, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		positions[i] = &Position{
			x,
			y,
		}
	}

	return positions
}

func abs(n int) int {
	if n < 0 {
		return -n;	
	}

	return n
}

func solveA(positions []*Position) int {
	maxArea := 0
	totalPositions := len(positions)

	for i, p1 := range positions {
		for j := i + 1; j < totalPositions; j++ {
			p2 := positions[j]
			area := (1 + abs(p1.x - p2.x)) * (1 + abs(p1.y - p2.y))

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func findCompressed(list []int, search int) (compressed int) {
	for i, value := range list {
		if value == search {
			compressed = i
			break
		}
	}

	return compressed
}

func findSubsum(prefixSum [][]int, p1 Position, p2 Position) int {
	return prefixSum[p2.x + 1][p2.y + 1] - prefixSum[p1.x][p2.y + 1] - prefixSum[p2.x + 1][p1.y] + prefixSum[p1.x][p1.y]
}

func solveB(positions []*Position) (maxArea int) {
	// sorted xs and ys
	xMap := map[int]bool{}
	yMap := map[int]bool{}

	for _, position := range positions {
		xMap[position.x] = true
		yMap[position.y] = true
	}

	xs := make([]int, len(xMap))
	ys := make([]int, len(yMap))

	i := 0

	for x := range xMap {
		xs[i] = x
		i++
	}

	i = 0

	for y := range yMap {
		ys[i] = y
		i++
	}

	slices.Sort(xs)
	slices.Sort(ys)

	// compressed grid
	grid := make([][]int, len(xs))

	for i := range xs {
		grid[i] = make([]int, len(ys))
	}

	for i, p1 := range positions {
		p2 := positions[(i + 1) % len(positions)]
		compressedP1 := Position{ findCompressed(xs, p1.x), findCompressed(ys, p1.y) }
		compressedP2 := Position{ findCompressed(xs, p2.x), findCompressed(ys, p2.y) }

		for cx := min(compressedP1.x, compressedP2.x); cx <= max(compressedP1.x, compressedP2.x); cx++ {
			for cy := min(compressedP1.y, compressedP2.y); cy <= max(compressedP1.y, compressedP2.y); cy++ {
				grid[cx][cy] = 1
			}
		}
	}
	
	// check outside coords
	gridCorners := [4]Position{
		{ 0, 0 },
		{ len(grid) - 1, 0 },
		{ 0, len(grid[0]) - 1 },
		{ len(grid) - 1, len(grid[0]) - 1 },
	}
	queue := make([]Position, 0)

	for _, gridCorner := range gridCorners {
		if grid[gridCorner.x][gridCorner.y] == 0 {
			queue = append(queue, gridCorner)
		}
	}

	outsideCoords := map[Position]bool{}

	for len(queue) > 0 {
		curPos := queue[0]
		queue = queue[1:]

		_, seen := outsideCoords[curPos]

		if seen {
			continue
		}

		outsideCoords[curPos] = true

		neighbours := [4]Position{
			{ curPos.x + 1, curPos.y },
			{ curPos.x - 1, curPos.y },
			{ curPos.x, curPos.y + 1 },
			{ curPos.x, curPos.y - 1 },
		}

		for _, neighbour := range neighbours {
			if neighbour.x < 0 || neighbour.x >= len(grid) || neighbour.y < 0 || neighbour.y >= len(grid[0]) {
				continue
			}

			if grid[neighbour.x][neighbour.y] == 1 {
				continue
			}

			queue = append(queue, neighbour)
		}
	}

	for cx, line := range grid {
		for cy := range line {
			_, outside := outsideCoords[Position{ cx, cy }]

			if !outside {
				grid[cx][cy] = 1
			}
		}
	}

	// build prefix sum
	prefixSum := make([][]int, len(xs) + 1)

	for i := range prefixSum {
		prefixSum[i] = make([]int, len(ys) + 1)
	}

	for cx, line := range grid {
		for cy := range line {
			prefixSum[cx + 1][cy + 1] = prefixSum[cx][cy + 1] + prefixSum[cx + 1][cy] - prefixSum[cx][cy] + grid[cx][cy]
		}
	}

	// find max area
	for i, p1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			p2 := positions[j]

			vertexes := [4]Position{
				{ findCompressed(xs, p1.x), findCompressed(ys, p1.y) },
				{ findCompressed(xs, p1.x), findCompressed(ys, p2.y) },
				{ findCompressed(xs, p2.x), findCompressed(ys, p2.y) },
				{ findCompressed(xs, p2.x), findCompressed(ys, p1.y) },
			}

			flag := true

			for i, from := range vertexes {
				to := vertexes[(i + 1) % len(vertexes)]

				var subsum int

				if (from.x == to.x && from.y <= to.y) || (from.y == to.y && from.x <= to.x) {
					subsum = findSubsum(prefixSum, from, to)
				} else {
					subsum = findSubsum(prefixSum, to, from)
				}

				var expectedEdgeSize int

				if from.x == to.x {
					expectedEdgeSize = 1 + abs(from.y - to.y)
				} else {
					expectedEdgeSize = 1 + abs(from.x - to.x)
				}

				if subsum != expectedEdgeSize {
					flag = false
					break	
				}
			}

			if !flag {
				continue
			}

			area := (1 + abs(p1.x - p2.x)) * (1 + abs(p1.y - p2.y))

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	start := time.Now()

	positions := parseFile()
	solutionA := solveA(positions)
	solutionB := solveB(positions)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
