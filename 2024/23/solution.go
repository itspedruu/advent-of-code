package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func parseFile() map[string]map[string]bool {
	fileBytes, err := ioutil.ReadFile("23/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")
	graph := make(map[string]map[string]bool)

	for _, line := range lines {
		vertices := strings.Split(line, "-")
		from := vertices[0]
		to := vertices[1]

		if graph[from] == nil {
			graph[from] = make(map[string]bool)
		}

		if graph[to] == nil {
			graph[to] = make(map[string]bool)
		}

		graph[from][to] = true
		graph[to][from] = true
	}

	return graph
}

func getKey(a string, b string, c string) [3]string {
	if a > c {
		a, c = c, a
	}

	if a > b {
		a, b = b, a
	}

	if b > c {
		b, c = c, b
	}

	return [3]string{ a, b, c}
}

func solveA(graph map[string]map[string]bool) int {
	startingVertices := make([]string, 0)

	for vertice, _ := range graph {
		if vertice[0] == 't' {
			startingVertices = append(startingVertices, vertice)
		}
	}

	networks := make(map[[3]string]bool)

	for _, from := range startingVertices {
		for adj := range graph[from] {
			for to := range graph[from] {
				if adj == to {
					continue
				}

				if (!graph[adj][to]) {
					continue
				}

				networks[getKey(from, adj, to)] = true
			}
		}
	}

	return len(networks)
}

type RTreeNode struct {
	value string
	next []*RTreeNode
}

func connect(graph map[string]map[string]bool, node *RTreeNode, vertex string) {
	if !graph[node.value][vertex] {
		return
	}

	for _, next := range node.next {
		connect(graph, next, vertex)
	}

	node.next = append(node.next, &RTreeNode{ vertex, make([]*RTreeNode, 0) })
}

func getHeight(node *RTreeNode) int {
	if len(node.next) == 0 {
		return 1
	}

	out := 1

	for _, next := range node.next {
		out = max(out, 1 + getHeight(next))
	}

	return out
}

type QueueElement struct {
	node *RTreeNode
	path []string
}

func solveB(graph map[string]map[string]bool) string {
	startingVertices := make([]string, 0)

	for vertice, _ := range graph {
		if vertice[0] == 't' {
			startingVertices = append(startingVertices, vertice)
		}
	}

	queue := make([]QueueElement, 0)
	seen := make(map[string]bool)

	for from, adjs := range graph {
		if seen[from] {
			continue
		}

		seen[from] = true

		network := &RTreeNode{ from, make([]*RTreeNode, 0) }

		initialPath := make([]string, 1)
		initialPath[0] = from

		queue = append(queue, QueueElement{ network, initialPath })

		for to := range adjs {
			seen[to] = true
			connect(graph, network, to)
		}
	}

	out := queue[0].path

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if len(element.path) > len(out) {
			out = element.path
		}

		for _, next := range element.node.next {
			newPath := make([]string, len(element.path) + 1)
			copy(newPath, element.path)
			newPath[len(element.path)] = next.value

			queue = append(queue, QueueElement{ next, newPath })
		}
	}

	sort.Strings(out)

	return strings.Join(out, ",")
}

func main() {
	start := time.Now()

	graph := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA := solveA(graph)
	fmt.Println("~ solving A:", time.Since(start))
	start = time.Now()

	solutionB := solveB(graph)
	fmt.Println("~ solving B:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
