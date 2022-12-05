package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// The crate positions are a map of stacks
type CrateStacks = map[int](*dsa.IntStack)

// Each movement has a quantity, an origin and a destination
type Movement struct {
	From   int
	To     int
	Amount int
}

// Parse the initial list of stacks
func parseStacks(lines []string) CrateStacks {
	// Initialize the map of empty stacks
	stacks := CrateStacks{}
	numStacks := len(strings.Split(lines[len(lines)-1], "   "))
	for i := 1; i <= numStacks; i++ {
		stack := dsa.NewIntStack()
		stacks[i] = &stack
	}

	// Iterate the individual positions of the input lines, in reverse order. If
	// the element is not a blank space, push it to the corresponding stack.
	for i := len(lines) - 1; i >= 0; i-- {
		line := []rune(lines[i])
		for j := 0; j < numStacks; j++ {
			if value := line[1+j*4]; value != ' ' {
				stacks[j+1].Push(int(value))
			}
		}
	}

	return stacks
}

// Parse the list of movement instructions
func parseMovement(line string) Movement {
	// TODO: use regular expressions, instead of fixed positions?
	words := strings.Split(line, " ")

	amount, err := strconv.Atoi(words[1])
	if err != nil {
		panic(fmt.Sprintf("Error parsing line '%v'", line))
	}

	from, err := strconv.Atoi(words[3])
	if err != nil {
		panic(fmt.Sprintf("Error parsing line '%v'", line))
	}

	to, err := strconv.Atoi(words[5])
	if err != nil {
		panic(fmt.Sprintf("Error parsing line '%v'", line))
	}
	return Movement{Amount: amount, From: from, To: to}
}

// Parse both the initial positions and the movements from the input
func readInput(path string) (CrateStacks, []Movement) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rawStacks []string
	var stacks CrateStacks
	var movements []Movement

	stackFlag := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// A blank line is the separation between the initial position of the
		// crates and the list of movements
		if line == "" {
			stacks = parseStacks(rawStacks)
			stackFlag = false
			continue
		}

		if stackFlag {
			rawStacks = append(rawStacks, line)
		} else {
			movements = append(movements, parseMovement(line))
		}

	}

	return stacks, movements
}
