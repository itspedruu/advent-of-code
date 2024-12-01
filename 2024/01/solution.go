package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func abs(number int) int {
	if number > 0 {
		return number;
	}

	return -number;
}

func parseFile() ([]int, []int) {
	fileBytes, err := ioutil.ReadFile("01/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, line := range lines {
		rawNumbers := strings.Split(line, "   ")

		leftNumber, err := strconv.Atoi(rawNumbers[0])

		if err != nil {
			panic(err)
		}

		rightNumber, err := strconv.Atoi(rawNumbers[1])

		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func solveA(leftList []int, rightList []int) int {
	distance := 0

	for i := 0; i < len(leftList); i++ {
		distance += abs(leftList[i] - rightList[i])
	}

	return distance
}

func solveB(leftList []int, rightList[] int) int {
	appearances := make(map[int]int)

	for _, number := range leftList {
		appearances[number] = 0
	}

	for _, number := range rightList {
		_, ok := appearances[number]

		if !ok {
			continue
		}

		appearances[number]++
	}

	out := 0

	for number, count := range appearances {
		out += number * count
	}

	return out
}

func main() {
	leftList, rightList := parseFile()

	solutionA := solveA(leftList, rightList)
	solutionB := solveB(leftList, rightList)

	fmt.Printf("A: %d\n", solutionA)
	fmt.Printf("B: %d\n", solutionB)
}
