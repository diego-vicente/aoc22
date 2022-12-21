package day21

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// A Maybe type holds a value and allows to make it a nil value
type Maybe[T any] struct {
	Value T
}

// A Monkey has an Id and an Operation
type Monkey struct {
	Id        string
	Operation string
	Result    *Maybe[int]
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
			Result:    nil,
		}
	} else {
		return Monkey{
			Id:        sides[0],
			Operation: sides[1],
			Result:    &Maybe[int]{x},
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
