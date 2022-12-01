package day01

import (
	"bufio"
	"os"
	"strconv"
)

// An inventory is a set of caloric values per elf, represented by an slice of
// integer values.
type Inventory = []int

// The input for Day 01 is represented by a list of values, representing each
// the calories of an item. Consecutive lines represent a single elf's
// inventory, and empty new lines represent that the inventory of the last elf
// is finished.
func readInput(path string) []Inventory {
	var input []Inventory
	var current Inventory

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			current = append(current, calories)
		} else {
			input = append(input, current)
			current = Inventory{}
		}

	}

	return input
}
