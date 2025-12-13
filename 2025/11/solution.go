package main

import (
	"fmt"
	"maps"
	"os"
	"sort"
	"strings"
	"time"
)

func parseFile() map[string][]string {
	fileBytes, err := os.ReadFile("11/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	graph := map[string][]string{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		from := parts[0][:len(parts[0]) - 1]

		graph[from] = make([]string, len(parts) - 1)
		copy(graph[from], parts[1:])
	}

	return graph
}

func countPaths(graph map[string][]string, in string, out string, memo map[string]int) int {
	cached, seen := memo[in]

	if seen {
		return cached
	}

	sum := 0

	for _, to := range graph[in] {
		if to == out {
			sum++
		} else {
			sum += countPaths(graph, to, out, memo)
		}
	}

	memo[in] = sum

	return sum
}

func getMemoKeyForRequirements(in string, reqsPassed map[string]bool) string {
	values := make([]string, len(reqsPassed))

	i := 0

	for req := range reqsPassed {
		values[i] = req
		i++
	}

	sort.Strings(values)
	memoKey := in

	for _, value := range values {
		memoKey = fmt.Sprintf("%s.%s", memoKey, value)
	}

	return memoKey
}

func countPathsWithRequirements(graph map[string][]string, in string, out string, memo map[string]int, reqs []string, reqsPassed map[string]bool) int {
	cached, seen := memo[getMemoKeyForRequirements(in, reqsPassed)]

	if seen {
		return cached
	}

	sum := 0

	for _, to := range graph[in] {
		if to == out {
			if len(reqsPassed) == len(reqs) {
				sum++
			}
		} else {
			newReqsPassed := map[string]bool{}
			maps.Copy(newReqsPassed, reqsPassed)

			for _, req := range reqs {
				if to == req {
					newReqsPassed[to] = true
				}
			}

			sum += countPathsWithRequirements(graph, to, out, memo, reqs, newReqsPassed)
		}
	}

	memo[getMemoKeyForRequirements(in, reqsPassed)] = sum

	return sum
}

func solve(graph map[string][]string) (int, int) {
	return countPaths(graph, "you", "out", map[string]int{}), countPathsWithRequirements(graph, "svr", "out", map[string]int{}, []string{"dac", "fft"}, map[string]bool{})
}

func main() {
	start := time.Now()

	graph := parseFile()
	solutionA, solutionB := solve(graph)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
