package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define a turn with their start and end IDs
type Turn struct {
	Start int
	End   int
}

// An ElfPair is the definition of each elf's turns
type ElfPair struct {
	First  Turn
	Second Turn
}

// Parse a line as an ElfPair definition as per the statement
func parsePair(line string) ElfPair {
	ids := []int{}

	for _, pair := range strings.Split(line, ",") {
		for _, id := range strings.Split(pair, "-") {
			value, err := strconv.Atoi(id)
			if err != nil {
				panic(fmt.Sprintf("Error parsing line '%v'", line))
			}

			ids = append(ids, value)
		}
	}

	return ElfPair{
		First:  Turn{Start: ids[0], End: ids[1]},
		Second: Turn{Start: ids[2], End: ids[3]},
	}
}

// Read each line of the input as an ElfPair definition
func readInput(path string) []ElfPair {
	var input []ElfPair

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parsePair(line))
	}

	return input
}
