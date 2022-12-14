package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A Point is represented as a 2D coordinate
type Point struct {
	X int
	Y int
}

// Materials can be Rock or Sand
type Material = string

const (
	Rock Material = "#"
	Sand Material = "o"
)

// A Cave has a materials Scan and a max depth known
type Cave struct {
	Scan  map[Point]Material
	Depth int
}

// Parse a point from the input format
func parsePoint(token string) Point {
	values := strings.Split(token, ",")

	x, err := strconv.Atoi(values[0])
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s", token))
	}

	y, err := strconv.Atoi(values[1])
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s", token))
	}

	return Point{x, y}
}

// Parse a set of corners from a line in the input file
func parseLine(line string) []Point {
	values := strings.Split(line, " -> ")
	corners := []Point{}

	for _, value := range values {
		corners = append(corners, parsePoint(value))
	}

	return corners
}

// Add Rock lines to a Cave.Scan following a set of corners
func (cave *Cave) addRockLines(corners []Point) {
	var start, end Point

	for i := 0; i < len(corners)-1; i++ {
		start, end = corners[i], corners[i+1]

		if start.X == end.X {
			// Define a vertical line of rock
			if start.Y > end.Y {
				start, end = end, start
			}

			for x, y := start.X, start.Y; y <= end.Y; y++ {
				cave.Scan[Point{x, y}] = Rock
			}
		} else if start.Y == end.Y {
			// Define a horizontal line of rock
			if start.X > end.X {
				start, end = end, start
			}

			for x, y := start.X, start.Y; x <= end.X; x++ {
				cave.Scan[Point{x, y}] = Rock
			}
		} else {
			panic("Cannot handle diagonal rock definition")
		}
	}
}

// Read the input file into a Cave object
func readInput(path string) Cave {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cave := Cave{Scan: map[Point]string{}, Depth: 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		corners := parseLine(line)

		// Update the largest known depth if necessary
		for _, corner := range corners {
			if corner.Y > cave.Depth {
				cave.Depth = corner.Y
			}
		}

		cave.addRockLines(corners)
	}

	return cave
}
