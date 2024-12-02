package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const MIN = 1
const MAX = 3

func parseFile() [][]int {
	fileBytes, err := ioutil.ReadFile("02/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	reports := make([][]int, 0)

	for _, line := range lines {
		rawNumbers := strings.Split(line, " ")
		report := make([]int, 0)

		for _, rawNumber := range rawNumbers {
			number, err := strconv.Atoi(rawNumber)

			if err != nil {
				panic(err)
			}

			report = append(report, number)
		}

		reports = append(reports, report)
	}

	return reports
}

func testReport(report []int) bool {
	order := 1

	if report[0] > report[1] {
		order = -1
	}

	i := 0

	for ; i < len(report) - 1; i++ {
		delta := (report[i + 1] - report[i]) * order

		if delta < MIN || delta > MAX {
			break
		}
	}

	return i == len(report) - 1
}

func solveA(reports [][]int) int {
	safe := 0

	for _, report := range reports {
		if testReport(report) {
			safe++
		}
	}

	return safe
}

func solveB(reports [][]int) int {
	safe := 0

	for _, report := range reports {
		for i := 0; i < len(report); i++ {
			placeholder := make([]int, 0)
			placeholder = append(placeholder, report[:i]...)
			placeholder = append(placeholder, report[(i + 1):]...)

			if testReport(placeholder) {
				safe++
				break
			}
		}
	}

	return safe
}

func main() {
	reports := parseFile()

	solutionA := solveA(reports)
	solutionB := solveB(reports)

	fmt.Printf("A: %d\n", solutionA)
	fmt.Printf("B: %d\n", solutionB)
}
