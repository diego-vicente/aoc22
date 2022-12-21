package day21

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A type is Operable if it can yield a value somehow
type Operable interface {
	Compute() int
	IsDefined() bool
}

// A result holds a value
type Result struct {
	Value int
}

func (r Result) Compute() int  { return r.Value }
func (Result) IsDefined() bool { return true }

// A literal is a given value from the statement or a explicit incognita
type Literal struct {
	Value   int
	Defined bool
}

func (l Literal) Compute() int    { return l.Value }
func (l Literal) IsDefined() bool { return l.Defined }

// An Operator is represented by a string
type Operator = string

const (
	Sum       Operator = "+"
	Substract Operator = "-"
	Multiply  Operator = "*"
	Divide    Operator = "/"
)

// An operation is formed by two Operables and an Operator
type Operation struct {
	Lhs      Operable
	Rhs      Operable
	Operator Operator
}

func (op Operation) Compute() int {
	switch op.Operator {
	case Sum:
		return op.Lhs.Compute() + op.Rhs.Compute()
	case Substract:
		return op.Lhs.Compute() - op.Rhs.Compute()
	case Multiply:
		return op.Lhs.Compute() * op.Rhs.Compute()
	case Divide:
		return op.Lhs.Compute() / op.Rhs.Compute()
	default:
		panic(fmt.Sprintf("Unknown operator: %s", op.Operator))
	}
}

// An operation is defined if both sides are defined
func (op Operation) IsDefined() bool {
	return op.Lhs.IsDefined() && op.Rhs.IsDefined()
}

// A Monkey has an Id and an Operation
type Monkey struct {
	Id        string
	Operation string
	Literal   *Literal
	Pack      *MonkeyPack
}

// A MonkeyPack has a relation of Ids to Monkeys
type MonkeyPack struct {
	Rel map[string](*Monkey)
}

// Parse a single line as a monkey
func parseMonkey(line string) Monkey {
	sides := strings.Split(line, ": ")
	x, err := strconv.Atoi(sides[1])
	if err != nil {
		return Monkey{
			Id:        sides[0],
			Operation: sides[1],
			Literal:   nil,
		}
	} else {
		return Monkey{
			Id:        sides[0],
			Operation: sides[1],
			Literal:   &Literal{x, true},
		}
	}
}

// Read the input as a MonkeyPack
func readInput(path string) MonkeyPack {
	input := MonkeyPack{
		Rel: map[string](*Monkey){},
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		monkey := parseMonkey(line)
		monkey.Pack = &input
		input.Rel[monkey.Id] = &monkey
	}

	return input
}
