package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseFile() []string {
	fileBytes, err := ioutil.ReadFile("08/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	return lines
}

func getAntennas(grid []string) map[rune][][2]int {
	antennas := make(map[rune][][2]int)

	for row, line := range grid {
		for col, char := range line {
			if char == '.' {
				continue
			}

			locations, ok := antennas[char]

			pos := [2]int{ row, col }

			if ok {
				antennas[char] = append(locations, pos)
			} else {
				newLocations := make([][2]int, 1)
				newLocations[0] = pos

				antennas[char] = newLocations
			}
		}
	}

	return antennas
}

func solveA(grid []string) int {
	antennas := getAntennas(grid)

	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	visited := make(map[[2]int]bool)
	out := 0

	for _, locations := range antennas {
		for i, locationA := range locations {
			for j := i + 1; j < len(locations); j++ {
				locationB := locations[j]

				drow := locationA[0] - locationB[0]
				dcol := locationA[1] - locationB[1]

				antinodes := [2][2]int{
					[2]int{ locationA[0] + drow, locationA[1] + dcol },
					[2]int{ locationB[0] - drow, locationB[1] - dcol },
				}

				for _, antinode := range antinodes {
					_, isVisited := visited[antinode]

					if !isVisited && antinode[0] >= 0 && antinode[0] <= maxRow && antinode[1] >= 0 && antinode[1] <= maxCol {
						out++
						visited[antinode] = true
					}
				}
			}
		}
	}

	return out
}

func solveB(grid []string) int {
	antennas := getAntennas(grid)

	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	visited := make(map[[2]int]bool)
	out := 0

	for _, locations := range antennas {
		for i, locationA := range locations {
			for j := i + 1; j < len(locations); j++ {
				locationB := locations[j]

				drow := locationA[0] - locationB[0]
				dcol := locationA[1] - locationB[1]

				t := 0
				done := false

				for !done {
					done = true

					antinodes := [2][2]int{
						[2]int{ locationA[0] + drow * t, locationA[1] + dcol * t },
						[2]int{ locationB[0] - drow * t, locationB[1] - dcol * t },
					}

					for _, antinode := range antinodes {
						_, isVisited := visited[antinode]

						if antinode[0] >= 0 && antinode[0] <= maxRow && antinode[1] >= 0 && antinode[1] <= maxCol {
							if !isVisited {
								visited[antinode] = true
								out++
							}

							done = false
						}
					}

					t++
				}
			}
		}
	}

	return out
}

func main() {
	grid := parseFile()

	solutionA := solveA(grid)
	solutionB := solveB(grid)

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
