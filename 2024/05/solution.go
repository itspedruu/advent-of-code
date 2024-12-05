package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func parseFile() ([][2]int, [][]int) {
	fileBytes, err := ioutil.ReadFile("05/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	firstSection := true

	priorities := make([][2]int, 0)
	updates := make([][]int, 0)

	for _, line := range lines {
		if line == "" {
			firstSection = false
			continue
		}

		if firstSection {
			rawNumbers := strings.Split(line, "|")

			leftNumber, _ := strconv.Atoi(rawNumbers[0])
			rightNumber, _ := strconv.Atoi(rawNumbers[1])

			priorities = append(priorities, [2]int{ leftNumber, rightNumber })
		} else {
			rawNumbers := strings.Split(line, ",")

			numbers := make([]int, 0)

			for _, rawNumber := range rawNumbers {
				number, _ := strconv.Atoi(rawNumber)

				numbers = append(numbers, number)
			}

			updates = append(updates, numbers)
		}
	}

	return priorities, updates
}

func buildDependencies(priorities [][2]int) map[int]map[int]bool {
	dependencies := make(map[int]map[int]bool)

	for _, priority := range priorities {
		_, ok := dependencies[priority[1]]

		if ok {
			dependencies[priority[1]][priority[0]] = true
		} else {
			newDependencies := make(map[int]bool)
			newDependencies[priority[0]] = true

			dependencies[priority[1]] = newDependencies
		}
	}

	return dependencies
}

func isUpdateOrdered(update []int, dependencies map[int]map[int]bool) bool {
	mustError := make(map[int]bool)

	for _, page := range update {
		_, errors := mustError[page]

		if errors {
			return false
		}

		dependencies, dependenciesOk := dependencies[page]

		if dependenciesOk {
			for dependency := range dependencies {
				mustError[dependency] = true
			}
		}
	}

	return true
}

func solveA(priorities [][2]int, updates [][]int) int {
	dependencies := buildDependencies(priorities)

	out := 0

	for _, update := range updates {
		if isUpdateOrdered(update, dependencies) {
			out += update[int((len(update) - 1) / 2)]
		}
	}

	return out
}

func orderUpdate(update []int, dependencies map[int]map[int]bool) []int {
	out := make([]int, 1)
	out[0] = update[0]

	for _, numberToInsert := range update[1:] {
		ok := false

		for i, number := range out {
			_, dependenciesOk := dependencies[number]

			if dependenciesOk {
				_, isDependency := dependencies[number][numberToInsert]

				if isDependency {
					newOut := make([]int, 0)
					newOut = append(newOut, out[:i]...)
					newOut = append(newOut, numberToInsert)
					newOut = append(newOut, out[i:]...)

					out = newOut
					ok = true
					break
				}
			}
		}

		if !ok {
			out = append(out, numberToInsert)
		}
	}

	return out
}

func solveB(priorities [][2]int, updates [][]int) int {
	dependencies := buildDependencies(priorities)

	out := 0

	for _, update := range updates {
		if !isUpdateOrdered(update, dependencies) {
			newUpdate := orderUpdate(update, dependencies)

			out += newUpdate[int((len(newUpdate) - 1) / 2)]
		}
	}

	return out
}

func main() {
	priorities, updates := parseFile()

	solutionA := solveA(priorities, updates)
	solutionB := solveB(priorities, updates)

	fmt.Printf("A: %d\n", solutionA)
	fmt.Printf("B: %d\n", solutionB)
}
