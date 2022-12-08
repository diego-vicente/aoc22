package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A Point is a 2D coordinate position
type Point struct {
	X int
	Y int
}

// A Forest has a height map and its dimensions
type Forest struct {
	Height     map[Point]int
	Visible    map[Point]bool
	Dimensions Point
}

// Parse the input as a Forest
func readInput(path string) Forest {
	heights := map[Point]int{}
	var x, y int

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for i, char := range line {
			x = i
			height, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("Could not parse line %s", line))
			}

			heights[Point{x, y}] = height
		}

		y += 1
	}

	return Forest{
		Height:     heights,
		Visible:    map[Point]bool{},
		Dimensions: Point{X: x + 1, Y: y},
	}
}
