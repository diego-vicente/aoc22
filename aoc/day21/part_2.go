package day21

import (
	"fmt"
)

// Solve an incognita with as a given equality
//
// The incognita is an Operable expression that contains a single undefined
// value somewhere down the tree. This function contains the logic to
// recursively unroll these symbolic trees to isolate the undefined value.
func solveIncognita(incognita Operable, expected int) int {
	if operation, ok := incognita.(Operation); ok {
		lhs, rhs := operation.Lhs, operation.Rhs

		switch operation.Operator {
		case Sum:
			if lhs.IsDefined() {
				// m + x = n --> x = n - m
				return solveIncognita(rhs, expected-lhs.Compute())
			} else {
				// x + m = n --> x = n - m
				return solveIncognita(lhs, expected-rhs.Compute())
			}
		case Multiply:
			if lhs.IsDefined() {
				// x * m = n --> x = n / m
				return solveIncognita(rhs, expected/lhs.Compute())
			} else {
				// m * x = n --> x = n / m
				return solveIncognita(lhs, expected/rhs.Compute())
			}
		case Divide:
			if lhs.IsDefined() {
				// m / x = n --> x = m / n
				return solveIncognita(rhs, lhs.Compute()/expected)
			} else {
				// x / m = n --> x = n * m
				return solveIncognita(lhs, expected*rhs.Compute())
			}
		case Substract:
			if lhs.IsDefined() {
				// m - x = n --> x = m - n
				return solveIncognita(rhs, lhs.Compute()-expected)
			} else {
				// x - m = n --> x = n + m
				return solveIncognita(lhs, expected+rhs.Compute())
			}
		default:
			panic(fmt.Sprintf("Unknown operator: %s", operation.Operator))
		}
	} else {
		return expected
	}
}

// Solve the second part by solving the 'humn' value
func solveSecondPart(path string) int {
	input := readInput(path)

	// Turn the 'humn' monkey into undefined
	input.Rel["humn"].Literal = &Literal{Defined: false}

	// Build the symbolic tree
	symbolic := input.Rel["root"].getSymbolic().(Operation)

	// Present the equality and solve it recursively
	known, incognita := symbolic.Lhs, symbolic.Rhs
	if !known.IsDefined() {
		known, incognita = incognita, known
	}

	return solveIncognita(incognita, known.Compute())
}
