package day21

import (
	"strings"
)

// Compute the value of the Operation held by a Monkey
func (m *Monkey) Compute() int {
	var operation Operable

	if m.Literal != nil {
		operation = m.Literal
	} else {
		tokens := strings.Split(m.Operation, " ")
		operation = Operation{
			Lhs:      m.Pack.Rel[tokens[0]],
			Rhs:      m.Pack.Rel[tokens[2]],
			Operator: tokens[1],
		}
	}

	return operation.Compute()
}

// Solve the first part by infering the value of root
func solveFirstPart(path string) int {
	input := readInput(path)
	return input.Rel["root"].Compute()
}
