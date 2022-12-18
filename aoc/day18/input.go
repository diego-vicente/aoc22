package day18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A Point is represented by their 3D coordinates
type Point struct {
	X int
	Y int
	Z int
}

// Parse a single line of the input into a Point
func parsePoint(line string) Point {
	values := strings.Split(line, ",")

	x, err := strconv.Atoi(values[0])
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s", line))
	}

	y, err := strconv.Atoi(values[1])
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s", line))
	}

	z, err := strconv.Atoi(values[2])
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s", line))
	}

	return Point{x, y, z}
}

// Read the inputa as a set of 3D points
func readInput(path string) dsa.Set[Point] {
	input := dsa.NewSet[Point]()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input.Add(parsePoint(line))
	}

	return input
}
