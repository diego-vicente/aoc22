package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A NoOp has no further attributes
type NoOp struct{}

// An AddX instruction has a value to add
type AddX struct{ X int }

// Parse a instruction from input
func parseInstruction(line string) []Executable {
	switch line[:4] {
	case "noop":
		return []Executable{NoOp{}}
	case "addx":
		x, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			panic(fmt.Sprintf("Error parsing line: %s", line))
		}
		// To simulate two cycles executions, we parse it as NoOp + AddX
		return []Executable{NoOp{}, AddX{X: x}}
	default:
		panic(fmt.Sprintf("Error parsing line: %s", line))
	}
}

// Parse the input as a list of instructions
func readInput(path string) []Executable {
	var input []Executable

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseInstruction(line)...)
	}

	return input
}
