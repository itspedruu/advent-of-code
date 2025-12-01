package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Path struct {
	direction int
	amount int
}

func parseFile() []*Path {
	fileBytes, err := os.ReadFile("01/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	paths := make([]*Path, len(lines))

	for i, line := range lines {
		direction := 1

		if line[0] == 'L' {
			direction = -1	
		}

		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		paths[i] = &Path{
			direction,
			amount,
		}
	}

	return paths
}

func solve(paths []*Path) (int, int) {
	digit := 50
	count := 0
	rotations := 0

	for _, path := range paths {
		if path.direction == 1 || digit == 0 {
			rotations += (digit + path.amount) / 100
		} else {
			rotations += ((100 - digit) + path.amount) / 100
		}

		digit = (digit + path.amount * path.direction) % 100

		if digit < 0 {
			digit += 100	
		}

		if digit == 0 {
			count++
		}
	}

	return count, rotations
}

func main() {
	start := time.Now()

	paths := parseFile()
	solutionA, solutionB := solve(paths)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
