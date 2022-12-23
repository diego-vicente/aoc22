package day23

import (
	"bufio"
	"os"
)

// A Point is represented by 2D coordinates
type Point struct {
	X int
	Y int
}

// A BoundingBox is defined by two corners
type BoundingBox struct {
	Min Point
	Max Point
}

// A Direction is a direction in the 2D space
type Direction = string

const (
	North     Direction = "N"
	NorthEast Direction = "NE"
	East      Direction = "E"
	SouthEast Direction = "SE"
	South     Direction = "S"
	SouthWest Direction = "SW"
	West      Direction = "W"
	NorthWest Direction = "NW"
)

// An elf holds their current position
type Elf struct {
	Position Point
}

// A grove includes a relation of elves' positions
type Grove struct {
	Elves map[Point](*Elf)
}

// Read the input as a Grove
func readInput(path string) Grove {
	elves := map[Point](*Elf){}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, char := range line {
			switch string(char) {
			case "#":
				pos := Point{x, y}
				elf := Elf{pos}
				elves[pos] = &elf
			case ".":
				continue
			default:
				panic("Error parsing input")
			}
		}
	}

	return Grove{Elves: elves}
}
