package day11

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// Solve the second part with 10000 rounds and without dividing
func solveSecondPart(path string) int {
	monkeys := readInput(path)

	// Compute the LCM to get a new control function
	var divisors []int
	for _, monkey := range monkeys {
		divisors = append(divisors, monkey.DivisibleBy)
	}
	lcm := dsa.LCM(divisors[0], divisors[1], divisors[1:]...)

	performRounds(monkeys, 10000, func(w int) int { return w % lcm })
	return monkeyBusiness(monkeys)
}
