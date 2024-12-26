package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
	"time"
	"math/big"
)

type Gate struct {
	op string
	inputs [2]string
	output string
}

func parseFile() ([]*Gate, map[string]uint, map[string]*Gate) {
	fileBytes, err := ioutil.ReadFile("24/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(fileBytes), "\n"), "\n")

	gates := make([]*Gate, 0)
	values := make(map[string]uint)
	outsTo := make(map[string]*Gate)

	stage := 0

	for _, line := range lines {
		if line == "" {
			stage++
			continue
		}

		if stage == 0 {
			parts := strings.Split(line, ": ")
			number, _ := strconv.Atoi(parts[1])

			values[parts[0]] = uint(number)
		} else {
			parts := strings.Split(line, " ")
			op := parts[1]
			inputs := [2]string{ parts[0], parts[2] }
			output := parts[4]

			if inputs[0] > inputs[1] {
				inputs[0], inputs[1] = inputs[1], inputs[0]
			}
			
			gate := &Gate{ op, inputs, output }

			outsTo[output] = gate
			gates = append(gates, gate)
		}
	}

	return gates, values, outsTo
}

func calculate(gate *Gate, values map[string]uint, outsTo map[string]*Gate) {
	_, ok := values[gate.output]	

	if ok {
		return 
	}

	for _, input := range gate.inputs {
		_, ok := values[input]

		if ok {
			continue
		}

		calculate(outsTo[input], values, outsTo)
	}

	if (gate.op == "AND") {
		values[gate.output] = values[gate.inputs[0]] & values[gate.inputs[1]]
	} else if (gate.op == "OR") {
		values[gate.output] = values[gate.inputs[0]] | values[gate.inputs[1]]
	} else {
		values[gate.output] = values[gate.inputs[0]] ^ values[gate.inputs[1]]
	}
}

func solve(gates []*Gate, values map[string]uint, outsTo map[string]*Gate) (int, string) {
	nBits := 0
	inputsTo := make(map[[2]string]*Gate)

	for _, gate := range gates {
		if gate.output[0] == 'z' {
			nBits++
		}

		calculate(gate, values, outsTo)

		inputsTo[gate.inputs] = gate
	}

	n := big.NewInt(0)

	wrongWires := make([]string, 0)
	var prevAddGate *Gate

	for i := 0; i < nBits; i++ {
		numberKey := fmt.Sprintf("%d", i)

		if i <= 9 {
			numberKey = fmt.Sprintf("0%d", i)
		}

		xKey := "x" + numberKey
		yKey := "y" + numberKey
		zKey := "z" + numberKey

		n.SetBit(n, i, values[zKey])

		addGate := outsTo[zKey]

		if i <= 1 || i == nBits - 1 {
			prevAddGate = addGate
			continue
		}

		if addGate.op != "XOR" {
			wrongWires = append(wrongWires, addGate.output)
			prevAddGate = addGate
			continue
		}

		var bitAddInput string
		var carryInput string

		for _, input := range addGate.inputs {
			curGate := outsTo[input]

			if curGate.inputs == [2]string{ xKey, yKey } && curGate.op == "XOR" {
				bitAddInput = input
			}

			if curGate.op == "OR" {
				carryInput = input
			}
		}

		filteredInputs := make([]string, 0)

		for _, input := range addGate.inputs {
			if bitAddInput == input || carryInput == input {
				continue
			}

			filteredInputs = append(filteredInputs, input)
		}

		if len(filteredInputs) > 0 {
			wrongWires = append(wrongWires, filteredInputs...)
			prevAddGate = addGate
			continue
		}
		
		carryGate := outsTo[carryInput]

		prevAddCarryGate := outsTo[carryGate.inputs[0]]
		prevBitCarryGate := outsTo[carryGate.inputs[1]]

		if prevAddCarryGate.inputs != prevAddGate.inputs || prevAddCarryGate.op != "AND" {
			prevAddCarryGate = outsTo[carryGate.inputs[1]]
			prevBitCarryGate = outsTo[carryGate.inputs[0]]
		}

		if prevAddCarryGate.inputs != prevAddGate.inputs || prevAddCarryGate.op != "AND" {
			wrongWires = append(wrongWires, prevAddCarryGate.output)
			prevAddGate = addGate
			continue
		}

		if (prevBitCarryGate.op != "AND") {
			wrongWires = append(wrongWires, prevBitCarryGate.output)
			prevAddGate = addGate
			continue
		}

		prevAddGate = addGate
	}

	sort.Strings(wrongWires)

	return int(n.Int64()), strings.Join(wrongWires, ",")
}

func main() {
	start := time.Now()

	gates, values, outsTo := parseFile()

	fmt.Println("~ reading file:", time.Since(start))
	start = time.Now()

	solutionA, solutionB := solve(gates, values, outsTo)
	fmt.Println("~ solving:", time.Since(start))

	fmt.Println("A:", solutionA)
	fmt.Println("B:", solutionB)
}
