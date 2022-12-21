package day21

import (
	"fmt"
	"strings"
)

// Compute the value of the Operation held by a Monkey
func (m *Monkey) Compute() int {
	tokens := strings.Split(m.Operation, " ")

	lhs := m.Pack.Rel[tokens[0]].Infer()
	rhs := m.Pack.Rel[tokens[2]].Infer()

	switch tokens[1] {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "/":
		return lhs / rhs
	case "*":
		return lhs * rhs
	default:
		panic(fmt.Sprintf("Unknown operator: %s", m.Operation))
	}
}

// Infer the value associated to a monkey
func (m *Monkey) Infer() int {
	// If there is no known result, compute the operation
	if m.Result == nil {
		m.Result = &Maybe[int]{m.Compute()}
	}

	return m.Result.Value
}

// Solve the first part by infering the value of root
func solveFirstPart(path string) int {
	input := readInput(path)
	return input.Rel["root"].Infer()
}
