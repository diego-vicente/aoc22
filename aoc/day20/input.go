package day20

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// Parse a single line as an integer
func parseItem(line string) int {
	x, err := strconv.Atoi(line)
	if err != nil {
		panic(fmt.Sprintf("Error parsing line: %s", line))
	}

	return x
}

// Read the input as a Ring
func readInput(path string) dsa.Ring[int] {
	input := dsa.NewRing[int]()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input.Append(parseItem(line))
	}

	return input
}
