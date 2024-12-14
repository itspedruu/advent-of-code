package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

const WIDTH = 101
const HEIGHT = 103
const CENTER_X = (WIDTH - 1) / 2
const CENTER_Y = (HEIGHT - 1) / 2

type Robot struct {
	startX int
	startY int
	vx int
	vy int
}

func parseFile() []Robot {
	fileBytes, err := ioutil.ReadFile("14/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	robots := make([]Robot, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")

		for partIndex, part := range parts {
			partNumbers := strings.Split(part, ",")

			leftNumber, _ := strconv.Atoi(partNumbers[0][2:])
			rightNumber, _ := strconv.Atoi(partNumbers[1])

			if partIndex == 0 {
				robots[i].startX = leftNumber
				robots[i].startY = rightNumber
			} else {
				robots[i].vx = leftNumber
				robots[i].vy = rightNumber
			}
		}
	}

	return robots
}

func solveA(robots []Robot) int {
	quadrants := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		deltaX := robot.startX + robot.vx * 100

		if deltaX < 0 && deltaX % WIDTH != 0 {
			deltaX = WIDTH + (deltaX % WIDTH)
		} else {
			deltaX = deltaX % WIDTH
		}

		deltaY := robot.startY + robot.vy * 100

		if deltaY < 0 && deltaY % HEIGHT != 0 {
			deltaY = HEIGHT + (deltaY % HEIGHT)
		} else {
			deltaY = deltaY % HEIGHT
		}

		if deltaX != CENTER_X && deltaY != CENTER_Y {
			if deltaX < CENTER_X {
				if deltaY < CENTER_Y {
					quadrants[0]++
				} else {
					quadrants[1]++
				}
			} else {
				if deltaY < CENTER_Y {
					quadrants[2]++
				} else {
					quadrants[3]++
				}
			}
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func solveB(robots []Robot) {
	for seconds := 0; seconds < 10000; seconds++ {
		fmt.Println("seconds", seconds)

		var grid [HEIGHT][WIDTH]int;

		for _, robot := range robots {
			deltaX := robot.startX + robot.vx * seconds

			if deltaX < 0 && deltaX % WIDTH != 0 {
				deltaX = WIDTH + (deltaX % WIDTH)
			} else {
				deltaX = deltaX % WIDTH
			}

			deltaY := robot.startY + robot.vy * seconds

			if deltaY < 0 && deltaY % HEIGHT != 0 {
				deltaY = HEIGHT + (deltaY % HEIGHT)
			} else {
				deltaY = deltaY % HEIGHT
			}

			grid[deltaY][deltaX]++
		}

		for _, line := range grid {
			for _, number := range line {
				if number == 0 {
					fmt.Print(".")
				} else {
					fmt.Print("1")
				}
			}

			fmt.Println("")
		}
	}
}

func main() {
	start := time.Now()

	robots := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA := solveA(robots)
	fmt.Println("~ solving A:", time.Since(start))
	fmt.Println("A:", solutionA)

	solveB(robots)
}
