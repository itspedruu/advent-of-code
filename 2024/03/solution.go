package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

func parseFile() []string {
	fileBytes, err := ioutil.ReadFile("03/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	return lines
}

func solveA(lines []string) int {
	out := 0

	mulRegex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	for _, line := range lines {
		matches := mulRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])

			out += left * right
		}
	}

	return out
}

func solveB(lines []string) int {
	out := 0

	instructionRegex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
	skip := false

	for _, line := range lines {
		matches := instructionRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				skip = false
			} else if match[0] == "don't()" {
				skip = true
			} else if !skip {

				left, _ := strconv.Atoi(match[1])
				right, _ := strconv.Atoi(match[2])

				out += left * right
			}
		}
	}

	return out
}

func main() {
	lines := parseFile()

	solutionA := solveA(lines)
	solutionB := solveB(lines)

	fmt.Printf("A: %d\n", solutionA)
	fmt.Printf("B: %d\n", solutionB)
}
