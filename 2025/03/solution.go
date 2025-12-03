package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const ASCII_ZERO = 48

func parseFile() [][]int {
	fileBytes, err := os.ReadFile("03/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	banks := make([][]int, len(lines))

	for i, line := range lines {
		bank := make([]int, len(line)) 
		banks[i] = bank

		for j, char := range line {
			bank[j]	= int(char - ASCII_ZERO)
		}
	}

	return banks
}

func pow10(exponent int) int {
	if exponent == 0 {
		return 1
	}

	out := 10

	for i := 2; i <= exponent; i++ {
		out *= 10	
	}

	return out
}

func solve(banks [][]int, n int) int {
	out := 0
	maxJoltages := len(banks[0])
	allValid := (1 << n) - 1

	for _, bank := range banks {
		flags := 1 << (n - 1)
		digits := make([]int, n)
		lastValidDigits := make([]int, n)
		digits[0] = bank[0]

		for i, joltage := range bank[1:] {
			maxDistance := min(n, maxJoltages - i - 1)
			minIndex := -1

			for distance := range maxDistance {
				digitIndex := n - distance - 1

				if joltage > digits[digitIndex] {
					minIndex = digitIndex
				}
			}

			if minIndex != -1 {
				flags = (allValid ^ ((1 << (n - minIndex - 1)) - 1))
				digits[minIndex] = joltage

				for j := minIndex + 1; j < n; j++ {
					digits[j] = 0
				}

				if flags == allValid {
					copy(lastValidDigits, digits)
				}
			}
		}

		for i, digit := range lastValidDigits {
			out += digit * pow10(n - i - 1)
		}
	}

	return out
}

func main() {
	start := time.Now()

	banks := parseFile()
	solutionA, solutionB := solve(banks, 2), solve(banks, 12)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A: ", solutionA)
	fmt.Println("B: ", solutionB)
}
