package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
)

const MAX_INT = int(^uint(0) >> 1)

const ADV = 0
const BXL = 1
const BST = 2
const JNZ = 3
const BXC = 4
const OUT = 5
const BDV = 6
const CDV = 7

func parseFile() ([]int, map[rune]int) {
	fileBytes, err := ioutil.ReadFile("17/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")

	program := make([]int, 0)
	registers := make(map[rune]int)


	for i, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")

		if i == 4 {
			rawNumbers := strings.Split(parts[1], ",")

			for _, rawNumber := range rawNumbers {
				number, _ := strconv.Atoi(rawNumber)

				program = append(program, number)
			}
		} else {
			number, _ := strconv.Atoi(parts[1])

			registers[rune(parts[0][len(parts[0]) - 1])] = number
		}
	}

	return program, registers
}

func solveA(program []int, registers map[rune]int) string {
	ip := 0

	out := make([]string, 0)

	for ip < len(program) {
		op := program[ip]
		input := program[ip + 1]
		combo := input

		if combo >= 4 && combo <= 6 {
			combo = registers[rune('A' + (combo - 4))]
		}

		switch op {
			case ADV:
				registers['A'] /= (1 << combo)
			case BXL:
				registers['B'] = registers['B'] ^ input
			case BST:
				registers['B'] = combo % 8
			case JNZ:
				if registers['A'] != 0 {
					ip = input
					continue
				}
			case BXC:
				registers['B'] = registers['B'] ^ registers['C']
			case OUT:
				out = append(out, strconv.Itoa(combo % 8))
			case BDV:
				registers['B'] = registers['A'] / (1 << combo)
			case CDV:
				registers['C'] = registers['A'] / (1 << combo)
		}

		ip += 2
	}

	return strings.Join(out, ",")
}

func solveB(program []int) int {
	// reverse engineer :)
	// b = (((a % 8) ^ 1) ^ 4) ^ (a / 2^((a % 8) ^ 1))

	queue := make([][2]int, 1)
	queue[0] = [2]int{ 1, len(program) - 1 }
	minA := MAX_INT

	for len(queue) > 0 {
		node := queue[0]
		startingA := node[0]
		ip := node[1]

		queue = queue[1:]

		for n := 0; n <= 7; n++ {
			a := startingA + n
			b := (((a % 8) ^ 1) ^ 4) ^ (a / (1 << ((a % 8) ^ 1)))

			if b % 8 == program[ip] {
				if ip == 0 && a < minA {
					minA = a
				} else if ip != 0 {
					queue = append(queue, [2]int{ a * 8, ip - 1 })
				}
			}
		}
	}

	return minA
}

func main() {
	start := time.Now()

	program, registers := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA := solveA(program, registers)
	fmt.Println("~ solving A:", time.Since(start))
	start = time.Now()

	solutionB := solveB(program)
	fmt.Println("~ solving B:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
