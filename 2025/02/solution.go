package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type IdRange struct {
	start int
	end int
}

func parseFile() []*IdRange {
	fileBytes, err := os.ReadFile("02/input.txt")

	if err != nil {
		panic(err)
	}

	content := string(fileBytes)
	rawRanges := strings.Split(content[:len(content) - 1], ",")

	idRanges := make([]*IdRange, len(rawRanges))

	for i, rawRange := range rawRanges {
		parts := strings.Split(rawRange, "-")
		start, err := strconv.Atoi(parts[0])

		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(parts[1])

		idRanges[i] = &IdRange{
			start,
			end,
		}
	}

	return idRanges
}

func flog(n int) int {
	digits := 0

	for n >= 1 {
		n /= 10
		digits++
	}

	return digits
}

func pow10(exponent int) int {
	out := 10

	for i := 2; i <= exponent; i++ {
		out *= 10	
	}

	return out
}

func factors(n int) []int {
	out := make([]int, 0)

	for x := 1; x < n; x++ {
		if n % x == 0 {
			out = append(out, x)
		}
	}

	return out
}

func repeat(x int, n int) int {
	digits := flog(x)
	totalDigits := digits * n
	out := x

	for i := 1; i < n; i++ {
		out += x * pow10(totalDigits - digits * i)
	}

	return out
}

func solveA(idRanges []*IdRange) int {
	out := 0

	for _, idRange := range idRanges {
		start := idRange.start
		end := idRange.end
		startDigits := flog(start)
		endDigits := flog(end)

		// check if both odds in digits
		if startDigits % 2 != 0 && endDigits % 2 != 0 {
			continue
		}

		// normalize range in even digits
		if startDigits % 2 != 0 {
			startDigits++
			start = pow10(startDigits - 1)
		}

		if endDigits % 2 != 0 {
			endDigits--
			end = pow10(endDigits) - 1
		}

		// invalid loop ranges
		splitter := pow10(startDigits / 2)
		invalidStart, invalidStartPart := start / splitter, start % splitter
		invalidEnd, invalidEndPart := end / splitter, end % splitter
		
		if invalidStartPart > invalidStart {
			invalidStart++
		}

		if invalidEndPart < invalidEnd {
			invalidEnd--
		}

		// loop invalid ids
		for ; invalidStart <= invalidEnd; invalidStart++ {
			out += invalidStart * splitter + invalidStart
		}
	}

	return out
}

func solveB(idRanges []*IdRange) int {
	// normalize ranges
	normalizedRanges := make([][2]int, 0)

	for _, idRange := range idRanges {
		start := idRange.start
		end := idRange.end
		startDigits := flog(start)
		endDigits := flog(end)

		if startDigits == endDigits {
			normalizedRanges = append(normalizedRanges, [2]int{ start, end })
			continue
		}

		if startDigits > 1 {
			normalizedEnd := pow10(endDigits - 1) - 1
			normalizedRanges = append(normalizedRanges, [2]int{ start, normalizedEnd })
		}

		if endDigits > 1 {
			normalizedStart := pow10(startDigits)
			normalizedRanges = append(normalizedRanges, [2]int{ normalizedStart, end })
		}
	}

	out := 0
	seen := map[int]bool{}

	for _, normalizedRange := range normalizedRanges {
		start := normalizedRange[0]
		end := normalizedRange[1]
		digits := flog(start)
		digitsFactors := factors(digits)

		// go through each repeatable sequence
		for _, factor := range digitsFactors {
			splitter := pow10(digits - factor)
			endPattern := end / splitter
			repeatN := digits / factor


			for pattern := start / splitter; pattern <= endPattern; pattern++ {
				n := repeat(pattern, repeatN)

				_, ok := seen[n]

				if n >= start && n <= end && !ok {
					seen[n] = true
					out += n
				}
			}
		}
	}

	return out
}

func main() {
	start := time.Now()

	idRanges := parseFile()
	solutionA := solveA(idRanges)
	solutionB := solveB(idRanges)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A: ", solutionA)
	fmt.Println("B: ", solutionB)
}
