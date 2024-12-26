package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

const PRUNE_CONST = 16777216

func parseFile() []int {
	fileBytes, err := ioutil.ReadFile("22/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")
	secrets := make([]int, len(lines))

	for i, rawNumber := range lines {
		number, _ := strconv.Atoi(rawNumber)
		secrets[i] = number
	}

	return secrets
}

func solve(secrets []int) (int, int) {
	solutionA := 0
	solutionB := 0

	sequences := make(map[[4]int]int)

	for _, secret := range secrets {
		prevPrice := secret % 10

		var deltas [2000]int
		var prices [2000]int

		for i := 0; i < 2000; i++ {
			secret = ((secret * 64) ^ secret) % PRUNE_CONST
			secret = ((secret / 32) ^ secret) % PRUNE_CONST
			secret = ((secret * 2048) ^ secret) % PRUNE_CONST
			
			deltas[i] = (secret % 10) - prevPrice
			prices[i] = secret % 10
			prevPrice = secret % 10
		}

		seen := make(map[[4]int]bool)

		for i := 0; i < 1997; i++ {
			key := [4]int{ deltas[i], deltas[i + 1], deltas[i + 2], deltas[i + 3] }

			if seen[key] {
				continue
			}

			seen[key] = true

			sellValue := prices[i + 3]

			if (sequences[key] + sellValue > solutionB) {
				solutionB = sequences[key] + sellValue
			}

			sequences[key] += sellValue
		}

		solutionA += secret
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	secrets := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(secrets)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
