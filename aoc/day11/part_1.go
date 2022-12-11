package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// A control function receives an integer and returns it in reasonable bounds
type ControlFunc = func(int) int

// Apply an operation to an integer
func apply(old int, operation string) int {
	var value int

	tokens := strings.Split(operation, " ")

	// Check right hand side
	if tokens[2] == "old" {
		value = old
	} else {
		rhs, err := strconv.Atoi(tokens[2])
		if err != nil {
			panic(fmt.Sprintf("Error parsing operation %s", operation))
		}
		value = rhs
	}

	// Check operator
	switch tokens[1] {
	case "+":
		return old + value
	case "*":
		return old * value
	default:
		panic("Unknown operation")
	}
}

// Run all rounds for the first part
func performRounds(monkeys [](*Monkey), rounds int, control ControlFunc) {
	for round := 1; round <= rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				item.Worry = apply(item.Worry, monkey.Operation)
				item.Worry = control(item.Worry)

				if item.Worry%monkey.DivisibleBy == 0 {
					monkeys[monkey.IfTrue].Items = append(
						monkeys[monkey.IfTrue].Items,
						item,
					)
				} else {
					monkeys[monkey.IfFalse].Items = append(
						monkeys[monkey.IfFalse].Items,
						item,
					)
				}

				monkey.Inspected++
			}

			monkey.Items = []Item{}
		}
	}
}

// Compute the product of the two busiest monkeys inspections
func monkeyBusiness(monkeys [](*Monkey)) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	return monkeys[0].Inspected * monkeys[1].Inspected
}

// Solve the first part using 20 rounds and dividing by 3
func solveFirstPart(path string) int {
	monkeys := readInput(path)
	performRounds(monkeys, 20, func(w int) int { return w / 3 })
	return monkeyBusiness(monkeys)
}
