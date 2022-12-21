package day21

import (
	"strings"
)

// Get the operations associated with a Monkey, but do not solve them
func (m *Monkey) getSymbolic() Operable {
	if m.Literal != nil {
		return m.Literal
	} else {
		tokens := strings.Split(m.Operation, " ")
		return Operation{
			Lhs:      m.Pack.Rel[tokens[0]].getSymbolic(),
			Rhs:      m.Pack.Rel[tokens[2]].getSymbolic(),
			Operator: tokens[1],
		}
	}
}

// Compute the value of the Operation held by a Monkey
func (m *Monkey) Compute() int {
	return m.getSymbolic().Compute()
}

// A Monkey is defined if it depends on defined results
func (m *Monkey) IsDefined() bool {
	return m.getSymbolic().IsDefined()
}

// Solve the first part by infering the value of root
func solveFirstPart(path string) int {
	input := readInput(path)
	return input.Rel["root"].Compute()
}
