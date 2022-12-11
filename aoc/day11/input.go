package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// An item simply stores its worry level
type Item struct {
	Worry int
}

// A monkey holds all its attributes from input
type Monkey struct {
	// List of items curently held
	Items []Item
	// Operation to be performed
	Operation string
	// Integer to check for multiplicity
	DivisibleBy int
	// Index of the monkey to throw item to if divisible
	IfTrue int
	// Index of the monkey to throw item to if not
	IfFalse int
	// Number of items inspected
	Inspected int
}

// Return true if the string starts with the prefix
func startsWith(str string, prefix string) bool {
	for i, char := range prefix {
		if char != rune(str[i]) {
			return false
		}
	}

	return true
}

// Parse a Monkey from an input block
func parseMonkey(raw string) *Monkey {
	var items []Item
	var operation string
	var divisibleBy int
	var ifTrue, ifFalse int

	for _, line := range strings.Split(raw, "\n") {
		if startsWith(line, "Monkey ") {
			continue
		} else if startsWith(line, "  Starting items") {
			list := strings.Split(line, ": ")[1]

			for _, value := range strings.Split(list, ", ") {
				worry, err := strconv.Atoi(value)
				if err != nil {
					panic(fmt.Sprintf("Error parsing line: %s", line))
				}

				items = append(items, Item{worry})
			}
		} else if startsWith(line, "  Operation") {
			operation = strings.Split(line, " = ")[1]
		} else if startsWith(line, "  Test") {
			words := strings.Split(line, " ")
			value, err := strconv.Atoi(words[len(words)-1])
			if err != nil {
				panic(fmt.Sprintf("Error parsing line: %s", line))
			}
			divisibleBy = value
		} else if startsWith(line, "    If true") {
			words := strings.Split(line, " ")
			value, err := strconv.Atoi(words[len(words)-1])
			if err != nil {
				panic(fmt.Sprintf("Error parsing line: %s", line))
			}
			ifTrue = value
		} else if startsWith(line, "    If false") {
			words := strings.Split(line, " ")
			value, err := strconv.Atoi(words[len(words)-1])
			if err != nil {
				panic(fmt.Sprintf("Error parsing line: %s", line))
			}
			ifFalse = value
		} else {
			panic(fmt.Sprintf("Error parsing line: %s", line))
		}

	}

	return &Monkey{
		Items:       items,
		Operation:   operation,
		DivisibleBy: divisibleBy,
		IfTrue:      ifTrue,
		IfFalse:     ifFalse,
		Inspected:   0,
	}
}

// Parse the input as a list of Monkeys
func readInput(path string) [](*Monkey) {
	var input [](*Monkey)

	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	for _, monkey := range strings.Split(string(buf), "\n\n") {
		input = append(input, parseMonkey(monkey))
	}

	return input
}
