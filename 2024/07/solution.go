package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"math/big"
)

func parseFile() map[int][]int {
	fileBytes, err := ioutil.ReadFile("07/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]

	equations := make(map[int][]int)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rawNumbers := strings.Split(parts[1], " ")

		rightSide := make([]int, 0)

		for _, rawNumber := range rawNumbers {
			number, _ := strconv.Atoi(rawNumber)
			rightSide = append(rightSide, number)
		}

		testValue, _ := strconv.Atoi(parts[0])
		equations[testValue] = rightSide
	}

	return equations
}

func solveA(equations map[int][]int) int {
	out := 0

	placeholder := big.NewInt(1)

	for testValue, numbers := range equations {
		ops := big.NewInt(0)
		maxOps := 1 << (len(numbers) - 1)

		for i := 1; i <= maxOps; i++ {
			value := numbers[0]

			for j := 1; j < len(numbers); j++ {
				op := ops.Bit(j - 1)

				if op == 0 {
					value += numbers[j]
				} else {
					value *= numbers[j]
				}

				if value > testValue {
					break
				}
			}

			if value == testValue {
				out += testValue
				break
			}

			ops.Add(placeholder, ops)
		}
	}

	return out
}

func addToOps(ops []uint64) {
	i := len(ops) - 1	

	for i >= 0 {
		ops[i] = (ops[i] + 1) % 3

		if ops[i] != 0 {
			break	
		}

		i--;
	}
}

func units(a int) int {
	out := 1

	for a >= 10 {
		a /= 10
		out++
	}

	return out;
}

func concatenate(a int, b int) int {
	return a * int(math.Pow(float64(10), float64(units(b)))) + b
}

func solveB(equations map[int][]int) int {
	out := 0

	for testValue, numbers := range equations {
		ops := make([]uint64, len(numbers) - 1)
		maxOps := int(math.Pow(float64(3), float64((len(numbers) - 1))))

		for i := 0; i < maxOps; i++ {
			value := numbers[0]

			for j := 1; j < len(numbers); j++ {
				op := ops[j - 1]

				if op == 0 {
					value += numbers[j]
				} else if op == 1 {
					value *= numbers[j]
				} else {
					value = concatenate(value, numbers[j])
				}

				if value > testValue {
					break
				}
			}

			if value == testValue {
				out += testValue

				break
			}

			addToOps(ops)
		}
	}

	return out
}

func main() {
	equations := parseFile()

	solutionA := solveA(equations)
	solutionB := solveB(equations)

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
