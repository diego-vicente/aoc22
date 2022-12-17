package day17

import (
	"os"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A direction is represented by a string
type Direction = string

const (
	Down  string = "Down"
	Up    string = "Up"
	Left  string = "Left"
	Right string = "Right"
)

// An input is a set of Directions, to be used as jets
func readInput(path string) dsa.Queue[Direction] {
	input := dsa.NewQueue[Direction]()

	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	for _, char := range buf {
		switch string(char) {
		case "<":
			input.Add(Left)
		case ">":
			input.Add(Right)
		default:
			panic("Unknown character in input")
		}
	}

	return input
}
