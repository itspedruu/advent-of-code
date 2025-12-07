package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const ASCII_ZERO = 48

type Operation struct {
	numbers []int
	method byte
}

func parseFile() ([]*Operation, []*Operation) {
	fileBytes, err := os.ReadFile("06/input.txt")	

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines) - 1]
	totalLines := len(lines)

	// Columns
	totalColumns := len(lines[0])
	columnCount, columnI, columnSpan := 1, 1, map[int]int{}

	for ; columnI < totalColumns; columnI++ {
		if lines[totalLines - 1][columnI] == ' ' {
			columnCount++
		} else {
			columnSpan[columnI - columnCount] = columnCount - 1
			columnCount = 1
		}
	}

	columnSpan[columnI - columnCount] = columnCount

	// Initialize operations
	totalOperations := len(columnSpan)
	operationsA := make([]*Operation, totalOperations)
	operationsB := make([]*Operation, totalOperations)

	operationI := 0
	
	for columnI := range totalColumns {
		if lines[totalLines - 1][columnI] == ' ' {
			continue
		}

		method := lines[totalLines - 1][columnI];

		operationsA[operationI] = &Operation{
			make([]int, totalLines - 1),
			method,
		}

		operationsB[operationI] = &Operation{
			make([]int, columnSpan[columnI]),
			method,
		}

		operationI++
	}

	// Numbers
	for i, line := range lines[:totalLines - 1] {
		operationI := -1
		lastSpanStart := -1

		for j, char := range line {
			_, ok := columnSpan[j]
			
			if ok {
				lastSpanStart = j
				operationI++
			}

			if char == ' ' {
				continue
			}

			operationsA[operationI].numbers[i] *= 10
			operationsA[operationI].numbers[i] += int(char - ASCII_ZERO)

			operationsB[operationI].numbers[j - lastSpanStart] *= 10
			operationsB[operationI].numbers[j - lastSpanStart] += int(char - ASCII_ZERO)
		}
	}

	return operationsA, operationsB
}

func solve(operations []*Operation) int {
	out := 0

	for _, operation := range operations {
		total := 0

		if operation.method == '*' {
			total = 1
		}

		for _, number := range operation.numbers {
			if operation.method == '*' {
				total *= number
			} else {
				total += number
			}
		}

		out += total
	}

	return out	
}

func main() {
	start := time.Now()

	operationsA, operationsB := parseFile()
	solutionA, solutionB := solve(operationsA), solve(operationsB)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A: ", solutionA)
	fmt.Println("B: ", solutionB)
}
