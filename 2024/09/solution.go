package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func parseFile() []string {
	fileBytes, err := ioutil.ReadFile("09/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "")
	lines = lines[:len(lines) - 1]

	return lines
}

func solveA(diskmap []string) int {
	disk := make([]int, 0)
	id := 0

	for i, rawNumber := range diskmap {
		curId := id

		if i % 2 != 0 {
			curId = -1 // free space	
		}

		size, _ := strconv.Atoi(rawNumber)

		if size > 0 {
			files := make([]int, size)

			for j := 0; j < size; j++ {
				files[j] = curId
			}

			disk = append(disk, files...)
		}

		if curId != -1 {
			id++
		}
	}

	i := 0
	j := len(disk) - 1

	for disk[i] != -1 {
		i++
	}

	for j > i {
		disk[i] = disk[j]
		disk[j] = -1

		i--
		j--

		if disk[i] != -1 {
			for disk[i] != -1 {
				i++
			}
		}

		if disk[j] == -1 {
			for disk[j] == -1 {
				j--
			}
		}
	}

	checksum := 0

	for i := 0; disk[i] != -1; i++ {
		checksum += i * disk[i]
	}

	return checksum
}

func solveB(diskmap []string) int {
	totalSize := 0
	freeSlots := make([][2]int, 0) // (startPos, size)
	files := make([][2]int, 0) // (startPos, size)

	for i, rawSize := range diskmap {
		size, _ := strconv.Atoi(rawSize)

		if i % 2 == 0 {
			files = append(files, [2]int{ totalSize, size })
		} else {
			freeSlots = append(freeSlots, [2]int{ totalSize, size })
		}

		totalSize += size
	}

	for i := len(files) - 1; i >= 0; i-- {
		filePos := files[i][0]
		fileSize := files[i][1]


		for j := 0; freeSlots[j][0] < filePos; j++ {
			freeSlotPos := freeSlots[j][0]

			if freeSlotPos >= filePos {
				break
			}

			freeSlotSize := freeSlots[j][1]

			if freeSlotSize < fileSize {
				continue	
			}

			files[i][0] = freeSlotPos

			freeSlots[j][0] += fileSize
			freeSlots[j][1] -= fileSize

			break
		}
	}

	checksum := 0

	for fileId, file := range files {
		filePos := file[0]
		fileSize := file[1]

		for j := 0; j < fileSize; j++ {
			checksum += (filePos + j) * fileId
		}
	}

	return checksum
}

func main() {
	diskmap := parseFile()

	solutionA := solveA(diskmap)
	solutionB := solveB(diskmap)

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
