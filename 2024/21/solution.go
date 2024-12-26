package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

const MAX_INT = int(^uint(0) >> 1)

func parseFile() []string {
	fileBytes, err := ioutil.ReadFile("21/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")

	return lines
}

type Key struct {
	row int
	col int
	char rune
}

func abs(n int) int {
	if n >= 0 {
		return n
	}

	return -n
}

func buildPath(distances [2]int, reverse bool) string {
	out := ""

	for i := 0; i < abs(distances[0]); i++ {
		if distances[0] > 0 {
			out += "v"
		} else {
			out += "^"
		}
	}

	for i := 0; i < abs(distances[1]); i++ {
		if distances[1] > 0 {
			out += ">"
		} else {
			out += "<"
		}
	}

	if (reverse) {
		temp := ""

		for i := len(out) - 1; i >= 0; i-- {
			temp += string(out[i])
		}

		out = temp
	}

	return out + "A"
}

func getPaths(keypad []string) map[rune]map[rune]string {
	paths := make(map[rune]map[rune]string)
	nVertices := len(keypad) * 3 - 1
	vertices := make([]Key, nVertices)

	i := 0

	for row, line := range keypad {
		for col, char := range line {
			if (char == '.') {
				continue
			}

			vertices[i] = Key{ row, col, char }
			i++
		}
	}

	for _, from := range vertices {
		paths[from.char] = make(map[rune]string)

		for _, to := range vertices {
			distances := [2]int{ to.row - from.row, to.col - from.col }

			if distances[1] > 0 && keypad[to.row][from.col] != '.' {
				paths[from.char][to.char] = buildPath(distances, false)
			} else if keypad[from.row][to.col] != '.' {
				paths[from.char][to.char] = buildPath(distances, true)
			} else {
				paths[from.char][to.char] = buildPath(distances, false)
			}
		}
	}

	return paths
}

func solve(codes []string) (int, int) {
	numericalKeypad := []string{
		"789",
		"456",
		"123",
		".0A",
	}

	directionalKeypad := []string{
		".^A",
		"<v>",
	}

	transitions := getPaths(numericalKeypad)
	directionalPaths := getPaths(directionalKeypad)

	for fromKey, fromValue := range directionalPaths {
		if transitions[fromKey] != nil {
			for toKey, toValue := range directionalPaths[fromKey] {
				transitions[fromKey][toKey] = toValue
			}
		} else {
			transitions[fromKey] = fromValue
		}
	}

	solutionA := 0
	solutionB := 0

	for _, code := range codes {
		curVertex := 'A'
		segments := make(map[string]int)

		for _, char := range code {
			curTransition := transitions[curVertex][char]
			curVertex = char

			segments[curTransition]++
		}

		codeNumber, _ := strconv.Atoi(code[:len(code) - 1])

		for i := 0; i < 25; i++ {
			if i == 2 {
				for segment, count := range segments {
					solutionA += codeNumber * len(segment) * count
				}
			}

			temp := make(map[string]int)

			for segment, count := range segments {
				curVertex = 'A'

				for _, char := range segment {
					curTransition := transitions[curVertex][char]
					curVertex = char

					temp[curTransition] += count
				}

				segments = temp
			}
		}

		for segment, count := range segments {
			solutionB += codeNumber * len(segment) * count
		}
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	codes := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(codes)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
