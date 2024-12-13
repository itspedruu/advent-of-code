package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

type Coord struct {
	x int
	y int
}

type Machine struct {
	buttons [2]Coord
	prize Coord
}

func parseFile() []Machine {
	fileBytes, err := ioutil.ReadFile("13/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")

	machines := make([]Machine, len(lines) / 4)

	stage := 0
	machineIndex := 0

	machines[machineIndex] = Machine{}

	for _, line := range lines {
		if line == "" {
			stage = 0
			machineIndex++
			continue
		}

		parts := strings.Split(line, ", ")
		y, _ := strconv.Atoi(parts[1][2:])

		if stage == 2 {
			x, _ := strconv.Atoi(parts[0][9:])

			machines[machineIndex].prize = Coord{ x, y }
		} else {
			x, _ := strconv.Atoi(parts[0][12:])

			machines[machineIndex].buttons[stage] = Coord{ x, y }
		}

		stage++
	}

	return machines
}

func solve(machines []Machine) (int, int) {
	solutionA, solutionB := 0, 0

	for _, machine := range machines {
		for i := 0; i < 2; i++ {
			aPresses := float64(machine.buttons[1].y * machine.prize.x - machine.buttons[1].x * machine.prize.y) / float64(machine.buttons[1].y * machine.buttons[0].x - machine.buttons[1].x * machine.buttons[0].y)
			bPresses := float64(machine.prize.x - machine.buttons[0].x * int(aPresses)) / float64(machine.buttons[1].x)

			if aPresses == float64(int(aPresses)) && bPresses == float64(int(bPresses)) {
				tokens := int(aPresses * 3 + bPresses)

				if i == 0 {
					solutionA += tokens
				} else {
					solutionB += tokens
				}
			}

			machine.prize.x += 10000000000000
			machine.prize.y += 10000000000000
		}
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	machines := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(machines)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
