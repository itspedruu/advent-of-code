package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type RTreeNode struct {
	value rune
	nodes []*RTreeNode
}

func parseFile() (*RTreeNode, []string) {
	fileBytes, err := ioutil.ReadFile("19/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")
	tree := &RTreeNode{ '*', make([]*RTreeNode, 0) }

	patterns := strings.Split(lines[0], ", ")

	for _, pattern := range patterns {
		curNode := tree

		for _, char := range pattern {
			var found *RTreeNode

			for _, node := range curNode.nodes {
				if node != nil && node.value == char {
					found = node
					break
				}
			}

			if found == nil {
				newNode := &RTreeNode{ char, make([]*RTreeNode, 0) }
				curNode.nodes = append(curNode.nodes, newNode)	
				curNode = newNode
			} else {
				curNode = found
			}
		}

		found := false

		for _, node := range curNode.nodes {
			if node == nil {
				found = true
				break
			}
		}

		if !found {
			curNode.nodes = append(curNode.nodes, nil)	
		}
	}

	return tree, lines[2:]
}

type MemoNode struct {
	design string	
	pattern *RTreeNode
}

func countPatterns(design string, curPattern *RTreeNode, tree *RTreeNode, memo map[MemoNode]int) int {
	memoNode := MemoNode{ design, curPattern }
	_, ok := memo[memoNode]

	if ok {
		return memo[memoNode]
	}

	count := 0

	for _, node := range curPattern.nodes {
		if node == nil {
			if len(design) == 0 {
				count = 1
				break
			}
			
			count += countPatterns(design, tree, tree, memo)

			continue
		} 

		if len(design) == 0 {
			continue
		}

		if node.value == rune(design[0]) {
			count += countPatterns(design[1:], node, tree, memo)
		}
	}

	memo[memoNode] = count

	return count
}

func solve(patterns *RTreeNode, designs []string) (int, int) {
	solutionA := 0
	solutionB := 0
	memo := make(map[MemoNode]int)

	for _, design := range designs {
		n := countPatterns(design, patterns, patterns, memo)

		if n > 0 {
			solutionA += 1
		}

		solutionB += n
	}

	return solutionA, solutionB
}

func main() {
	start := time.Now()

	patterns, designs := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(patterns, designs)
	fmt.Println("~ solving:", time.Since(start))
	start = time.Now()

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
