package day03

import (
	"bufio"
	"os"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

type Rucksack struct {
	First  dsa.IntSet
	Second dsa.IntSet
}

// Map each item to their priority
func Priority(char rune) int {
	if char < 97 {
		return int(char) - 38
	} else {
		return int(char) - 96
	}
}

// Parse a line as a Rucksack definition
func parseRucksack(line string) Rucksack {
	rucksack := Rucksack{dsa.NewIntSet(), dsa.NewIntSet()}

	size := len(line) / 2

	for _, value := range line[:size] {
		rucksack.First.Add(Priority(value))
	}

	for _, value := range line[size:] {
		rucksack.Second.Add(Priority(value))
	}

	return rucksack
}

// Read each line of the input as a Rucksack definition
func readInput(path string) []Rucksack {
	var input []Rucksack

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseRucksack(line))
	}

	return input
}
