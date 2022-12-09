package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Represent all the possible directions
type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
	Down  Direction = "D"
	Up    Direction = "U"
)

// A movement is represented by a direction and the times it happens
type Movement struct {
	To    Direction
	Times int
}

// Parse a Movement as defined in the input
func parseMovement(line string) Movement {
	parts := strings.Split(line, " ")

	times, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("Could not parse line: %s", line))
	}

	var to Direction
	switch parts[0] {
	case "U":
		to = Up
	case "D":
		to = Down
	case "L":
		to = Left
	case "R":
		to = Right
	}

	return Movement{To: to, Times: times}
}

// Parse the input as a list of Movements
func readInput(path string) []Movement {
	var input []Movement

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseMovement(line))
	}

	return input
}
