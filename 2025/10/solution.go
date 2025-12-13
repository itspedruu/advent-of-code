package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	diagram []int
	wiring [][]int
	joltages []int
}

func parseFile() []Machine {
	fileBytes, err := os.ReadFile("10/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	machines := make([]Machine, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		rawJoltages := strings.Split(parts[len(parts) - 1][1:len(parts[len(parts) - 1]) - 1], ",")

		machines[i] = Machine{
			make([]int, len(parts[0]) - 2),
			make([][]int, len(parts) - 2),
			make([]int, len(rawJoltages)),
		}

		for diagramIndex, char := range parts[0][1:len(parts[0]) - 1] {
			if char == '#' {
				machines[i].diagram[diagramIndex] = 1
			} else {
				machines[i].diagram[diagramIndex] = 0
			}
		}

		for wiringIndex, part := range parts[1:len(parts) - 1] {
			rawButtons := strings.Split(part[1:len(part) - 1], ",")
			machines[i].wiring[wiringIndex] = make([]int, len(rawButtons))

			for buttonIndex, rawButton := range rawButtons {
				button, _ := strconv.Atoi(rawButton)
				machines[i].wiring[wiringIndex][buttonIndex] = button
			}
		}

		for joltageIndex, rawJoltage := range rawJoltages {
			joltage, _ := strconv.Atoi(rawJoltage)
			machines[i].joltages[joltageIndex] = joltage
		}
	}

	return machines
}

func equalLights(a []int, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func solveA(machines []Machine) (out int) {
	for _, machine := range machines {
		queue := make([][]int, len(machine.wiring))

		for i := range queue {
			element := make([]int, 1)
			element[0] = i

			queue[i] = element
		}

		for len(queue) > 0 {
			lights := make([]int, len(machine.diagram))
			order := queue[0]
			queue = queue[1:]

			for _, wiringIndex := range order {
				for _, button := range machine.wiring[wiringIndex] {
					lights[button] = (lights[button] + 1) % 2
				}
			}

			if equalLights(machine.diagram, lights) {
				out += len(order)
				break
			}

			for newIndex := order[len(order) - 1] + 1; newIndex < len(machine.wiring); newIndex++ {
				newOrder := make([]int, len(order) + 1)
				copy(newOrder, order)
				newOrder[len(newOrder) - 1] = newIndex
				queue = append(queue, newOrder)
			}
		}
	}

	return out
}

func main() {
	start := time.Now()

	machines := parseFile()
	solutionA := solveA(machines)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
}
