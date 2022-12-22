package day22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A Point is represented by 2D coordinates
type Point struct {
	X int
	Y int
}

// A Direction is an ordinal of the 4 possible orientations
type Direction = int

const (
	East  Direction = 0
	South Direction = 1
	West  Direction = 2
	North Direction = 3
)

// A Turn is able to do arithmetic on the Direction values
type Turn = int

const (
	Left  Turn = -1
	Right Turn = 1
)

// A Point's terrain can be empty, a solid wall or non-existent
type Terrain = string

const (
	Empty   Terrain = "."
	Wall    Terrain = "#"
	Nothing Terrain = " "
)

// A Jungle contains a map of points to their contents
type Jungle struct {
	Map map[Point]Terrain
}

// An instruction has an amount of steps to take or turn to make
type Instruction struct {
	Amount    int
	Direction Turn
}

// Parse an instruction list from the input file
func parseInstructions(line string) []Instruction {
	instructions := []Instruction{}
	token := ""

	for _, byte := range line {
		if char := string(byte); char == "R" || char == "L" {
			// If it is a turning instruction, parse the previous and add this one
			if token != "" {
				value, err := strconv.Atoi(token)
				if err != nil {
					panic(fmt.Sprintf("Error parsing token %s", token))
				}

				instructions = append(instructions, Instruction{Amount: value})
				token = ""
			}

			if char == "R" {
				instructions = append(instructions, Instruction{Direction: Right})
			} else if char == "L" {
				instructions = append(instructions, Instruction{Direction: Left})
			} else {
				panic(fmt.Sprintf("Unknown direction %s", char))
			}
		} else {
			// If not keep on accumulating the number
			token += string(byte)
		}
	}

	// Parse the remaining instruction (if any)
	if token != "" {
		value, err := strconv.Atoi(token)
		if err != nil {
			panic(fmt.Sprintf("Error parsing token %s", token))
		}

		instructions = append(instructions, Instruction{Amount: value})
	}

	return instructions
}

// Read the input as a Jungle map and an Instruction list
func readInput(path string) (Jungle, []Instruction) {
	jungleMap := map[Point]Terrain{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		if line == "" {
			break
		}

		for x, char := range line {
			switch string(char) {
			case ".":
				jungleMap[Point{x, y}] = Empty
			case "#":
				jungleMap[Point{x, y}] = Wall
			default:
				continue
			}
		}
	}

	scanner.Scan()
	instructions := parseInstructions(scanner.Text())

	return Jungle{jungleMap}, instructions
}
