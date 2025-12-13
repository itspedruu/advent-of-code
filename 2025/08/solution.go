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
	z int
}

type Pair struct {
	p1 *Position
	p2 *Position
	distance int
}

func parseFile() []*Position {
	fileBytes, err := os.ReadFile("08/input.txt")	

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
		z, _ := strconv.Atoi(parts[2])

		positions[i] = &Position{
			x,
			y,
			z,
		}
	}

	return positions
}

func distance3(p1 *Position, p2 *Position) int {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	deltaZ := p1.z - p2.z

	return deltaX * deltaX + deltaY * deltaY + deltaZ * deltaZ;
}

func getTotalPairs(totalPositions int) int {
	out := 0

	for i := 1; i < totalPositions; i++ {
		out += i
	}

	return out
}

func find(parents map[*Position]*Position, p1 *Position) *Position {
	cur := p1

	for parents[cur] != cur {
		parents[cur] = parents[parents[cur]]
		cur = parents[cur]
	}

	return cur
}

func union(parents map[*Position]*Position, size map[*Position]int, p1 *Position, p2 *Position) {
	rootP1 := find(parents, p1)
	rootP2 := find(parents, p2)

	if rootP1 == rootP2 {
		return
	}

	if size[rootP1] > size[rootP2] {
		parents[rootP2] = rootP1
		size[rootP1] += size[rootP2]
		delete(size, rootP2)
	} else {
		parents[rootP1] = rootP2
		size[rootP2] += size[rootP1]
		delete(size, rootP1)
	}
}

func solve(positions []*Position) (int, int) {
	// Pairs
	totalPositions := len(positions)
	totalPairs := getTotalPairs(totalPositions)
	pairs := make([]Pair, totalPairs)
	pairI := 0

	for i, p1 := range positions {
		for j := i + 1; j < totalPositions; j++ {
			p2 := positions[j]

			pairs[pairI] = Pair{ p1, p2, distance3(p1, p2) }
			pairI++
		}
	}

	slices.SortFunc(pairs, func(a Pair, b Pair) int {
		return a.distance - b.distance
	})

	// UF
	parents := map[*Position]*Position{}
	size := map[*Position]int{}

	for _, position := range positions {
		parents[position] = position
		size[position] = 1
	}

	pairI = 0
	solutionA, solutionB := 0, 0

	for len(size) != 1 {
		pair := pairs[pairI]
		union(parents, size, pair.p1, pair.p2)
		pairI++

		if pairI == 1000 {
			topCircuits := [3]int{ 1, 1, 1 }

			for _, count := range size {
				if count > topCircuits[0] {
					topCircuits[2] = topCircuits[1]
					topCircuits[1] = topCircuits[0]
					topCircuits[0] = count
				} else if count > topCircuits[1] {
					topCircuits[2] = topCircuits[1]
					topCircuits[1] = count
				} else if count > topCircuits[2] {
					topCircuits[2] = count
				}
			}

			solutionA = topCircuits[0] * topCircuits[1] * topCircuits[2]
		}
	}

	solutionB = pairs[pairI - 1].p1.x * pairs[pairI - 1].p2.x

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	positions := parseFile()
	solutionA, solutionB := solve(positions)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}

