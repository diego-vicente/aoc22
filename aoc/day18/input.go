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

// A Bounding Box represents the minimum and maximum points of a shape
type BoundingBox struct {
	Min Point
	Max Point
}

// A Droplet is their set of lava points and a bounding box
type Droplet struct {
	Points dsa.Set[Point]
	Bounds BoundingBox
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

// Read the input as a lava droplet
func readInput(path string) Droplet {
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

	result := Droplet{
		Points: input,
		Bounds: ComputeBoundingBox(input),
	}

	return result
}
