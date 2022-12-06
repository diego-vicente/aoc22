package day06

import (
	"os"
)

// Read the input as a string
func readInput(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(buf)
}
