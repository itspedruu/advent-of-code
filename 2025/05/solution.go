package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type IdRange struct {
	start int
	end int
}

func intersects(a *IdRange, b *IdRange) bool {
	return ((a.start >= b.start && a.start <= b.end) ||
		(a.end >= b.start && a.end <= b.end) ||
		(b.start >= a.start && b.start <= a.end) ||
		(b.end >= a.start && b.end <= a.end));
}

func binarySearch(idRanges []*IdRange, searchRange *IdRange) int {
	left := 0
	right := len(idRanges) - 1
	middle := 0

	for left <= right {
		middle = left + (right - left) / 2
		selectedRange := idRanges[middle]

		if intersects(searchRange, selectedRange) {
			return middle
		}

		if searchRange.start > selectedRange.end {
			left = middle + 1
		} else if searchRange.start < selectedRange.start {
			right = middle - 1
		}
	}

	return middle
}

func insert(idRanges []*IdRange, element *IdRange) []*IdRange {
	if len(idRanges) == 0 {
		return append(idRanges, element)
	}

	index := binarySearch(idRanges, element)
	selectedRange := idRanges[index]
	elementIndex := index

	if intersects(element, selectedRange) {
		element.end = max(element.end, selectedRange.end)
		element.start = min(element.start, selectedRange.start)
	} else if element.start > selectedRange.start {
		idRanges = slices.Insert(idRanges, index + 1, element)
		elementIndex = index + 1
	} else {
		idRanges = slices.Insert(idRanges, index, element)
		elementIndex = index
	}

	sliceL, sliceR := elementIndex, elementIndex
	slideL, slideR := elementIndex, elementIndex
	finalIndex := len(idRanges) - 1
	slideLFlag, slideRFlag := elementIndex != 0, elementIndex != finalIndex

	for slideLFlag || slideRFlag {
		if slideLFlag {
			slideL--;
			
			if slideL == 0 {
				slideLFlag = false
			}

			if intersects(element, idRanges[slideL]) {
				element.start = min(idRanges[slideL].start, element.start)
				sliceL = slideL
			} else {
				slideLFlag = false
			}
		}

		if slideRFlag {
			slideR++

			if slideR == finalIndex {
				slideRFlag = false
			}

			if intersects(element, idRanges[slideR]) {
				element.end = max(idRanges[slideR].end, element.end)
				sliceR = slideR
			} else {
				slideLFlag = false
			}
		}
	}

	newIdRanges := make([]*IdRange, 0)
	newIdRanges = append(newIdRanges, idRanges[:sliceL]...);
	newIdRanges = append(newIdRanges, element)
	newIdRanges = append(newIdRanges, idRanges[sliceR + 1:]...);

	return newIdRanges
}

func parseFile() ([]*IdRange, []int) {
	fileBytes, err := os.ReadFile("05/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	total := len(lines) - 1
	lines = lines[:total]
	i := 0
	
	freshRanges := make([]*IdRange, 0)

	for ; i < total; i++ {
		line := lines[i]

		if line == "" {
			i++
			break
		}

		parts := strings.Split(line, "-")
		from, err := strconv.Atoi(parts[0])

		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		freshRanges = insert(freshRanges, &IdRange{ from, to })
	}

	availableIds := make([]int, total - i)
	j := 0

	for ; i < total; i++ {
		id, err := strconv.Atoi(lines[i])

		if err != nil {
			panic(err)
		}

		availableIds[j] = id
		j++
	}

	return freshRanges, availableIds
}

func solveA(idRanges []*IdRange, ids []int) int {
	out := 0

	for _, id := range ids {
		index := binarySearch(idRanges, &IdRange{ id, id })
		selectedRange := idRanges[index]

		if id >= selectedRange.start && id <= selectedRange.end {
			out++
		}
	}

	return out
}

func solveB(idRanges []*IdRange) int {
	out := 0

	for _, idRange := range idRanges {
		out += idRange.end - idRange.start + 1
	}

	return out
}

func main() {
	start := time.Now()

	freshRanges, availableIds := parseFile()
	solutionA := solveA(freshRanges, availableIds)
	solutionB := solveB(freshRanges)

	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A: ", solutionA)
	fmt.Println("B: ", solutionB)
}
