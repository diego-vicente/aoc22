package day06

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// Check if the last seen characters are the header
func isHeader(seen dsa.Queue[int], length int) bool {
	set := dsa.NewSet[int]()
	for _, elem := range seen.Values() {
		set.Add(elem)
	}

	return set.Size() >= length
}

// Return the index where the header of the message ends
func headerIndex(msg string, length int) int {
	seen := dsa.NewQueue[int]()

	for index, char := range msg {
		if seen.Size() == length {
			seen.Pop()
		}

		seen.Add(int(char))

		if isHeader(seen, length) {
			return index + 1
		}
	}

	return -1
}

// Solve the first part using headers of length 4
func solveFirstPart(path string) int {
	input := readInput(path)
	return headerIndex(input, 4)
}
