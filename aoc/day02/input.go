package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Represent as enum the different possible shapes
type HandShape string

const (
	Rock     HandShape = "Rock"
	Paper              = "Paper"
	Scissors           = "Scissors"
)

// A Match is defined by a player's move and an opponent's move
type Match struct {
	Player   HandShape
	Opponent HandShape
}

// Represent as enum the different possible results of a match
type Result string

const (
	Defeat  Result = "Defeat"
	Draw           = "Draw"
	Victory        = "Victory"
)

// Parse a line as a match definition
func parseMatch(line string) Match {
	var player HandShape
	var opponent HandShape

	values := strings.Split(line, " ")

	switch values[0] {
	case "A":
		opponent = Rock
	case "B":
		opponent = Paper
	case "C":
		opponent = Scissors
	default:
		panic(fmt.Sprintf("Unknown move %s", values[0]))
	}

	switch values[1] {
	case "X":
		player = Rock
	case "Y":
		player = Paper
	case "Z":
		player = Scissors
	default:
		panic(fmt.Sprintf("Unknown move %s", values[0]))
	}

	return Match{Player: player, Opponent: opponent}
}

// Read each line of the input as a match definition
func readInput(path string) []Match {
	var input []Match

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseMatch(line))
	}

	return input
}
