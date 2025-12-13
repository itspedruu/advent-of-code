package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

type Grid struct {
	size int
	quantities []int
}

func parseFile() ([]int, []Grid) {
	fileBytes, err := os.ReadFile("12/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	presents := make([]int, 6)
	grids := make([]Grid, len(lines) - 30)
	i := 0

	for lineIndex, line := range lines {
		if line == "" {
			i++
			continue
		}

		if i == 6 {
			rows, _ := strconv.Atoi(line[:2])
			cols, _ := strconv.Atoi(line[3:5])
			quantities := make([]int, 6)

			for partI, part := range strings.Split(line, " ")[1:] {
				quantity, _ := strconv.Atoi(part)
				quantities[partI] = quantity
			}

			grids[lineIndex - 30] = Grid{
				rows * cols,
				quantities,
			}
		} else if line[1] == ':' {
			continue
		} else {
			for _, char := range line {
				if char == '#' {
					presents[i]++
				}
			}
		}
	}

	return presents, grids
}

func solve(presents []int, grids []Grid) (out int) {
	for _, grid := range grids {
		presentsSizeNeeded := 0

		for i, quantity := range grid.quantities {
			presentsSizeNeeded += quantity * presents[i]	
		}

		if presentsSizeNeeded < grid.size {
			out++
		}
	}

	return out
}

func main() {
	start := time.Now()

	presents, grids := parseFile()
	solutionA := solve(presents, grids)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
}
