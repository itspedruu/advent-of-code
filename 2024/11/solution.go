package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"time"
)

func parseFile() []int {
	fileBytes, err := ioutil.ReadFile("11/input.txt")

	if err != nil {
		panic(err)
	}

	rawNumbers := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), " ")
	rawNumbers = rawNumbers[:len(rawNumbers)]

	stones := make([]int, len(rawNumbers))

	for i, rawNumber := range rawNumbers {
		number, _ := strconv.Atoi(rawNumber)

		stones[i] = number
	}

	return stones
}

func units(a int) int {
	out := 1

	for a >= 10 {
		a /= 10
		out++
	}

	return out;
}

func applyRules(stone int, memo map[[2]int]int, ttl int) int {
	if ttl == 0 {
		return 1
	}

	memoKey := [2]int{ stone, ttl }
	memoOut, ok := memo[memoKey]

	if ok {
		return memoOut
	}

	digits := units(stone)

	if stone == 0 {
		out := applyRules(1, memo, ttl - 1)

		memo[memoKey] = out

		return out
	} else if digits % 2 == 0 {
		partDivider := int(math.Pow(float64(10), float64(digits / 2)))
		out := applyRules(stone / partDivider, memo, ttl - 1) + applyRules(stone % partDivider, memo, ttl - 1)	

		memo[memoKey] = out

		return out
	} else {
		out := applyRules(stone * 2024, memo, ttl - 1)

		memo[memoKey] = out

		return out
	}
}

func solveA(stones []int) (int, int) {
	solutionA := 0
	solutionB := 0

	memo := make(map[[2]int]int)

	for _, stone := range stones {
		solutionB += applyRules(stone, memo, 75)
		solutionA += applyRules(stone, memo, 25)
	}
	
	return solutionA, solutionB
}

func main() {
	start := time.Now()

	stones := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solveA(stones)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
