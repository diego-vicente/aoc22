package day19

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A Material can be ore, clay, obsidian, or geode
type Material = string

const (
	Ore      string = "ore"
	Clay     string = "clay"
	Obsidian string = "obsidian"
	Geode    string = "geode"
)

// A list of requirements map materials to their needed amount
type Requirements = map[Material]int

// A blueprint contains an ID and a list of
type Blueprint struct {
	Id    int
	Robot map[Material]Requirements
}

// Parse a line of input into a Blueprint
func parseBlueprint(line string) Blueprint {
	tokens := strings.Split(line, " ")

	rawValues := []string{
		tokens[1][:len(tokens[1])-1],
		tokens[6],
		tokens[12],
		tokens[18],
		tokens[21],
		tokens[27],
		tokens[30],
	}

	parsedValues := []int{}

	for _, raw := range rawValues {
		parsed, err := strconv.Atoi(raw)
		if err != nil {
			panic(fmt.Sprintf("Error parsing line: %s", raw))
		}
		parsedValues = append(parsedValues, parsed)
	}

	return Blueprint{
		Id: parsedValues[0],
		Robot: map[Material]Requirements{
			Ore:      {Ore: parsedValues[1]},
			Clay:     {Ore: parsedValues[2]},
			Obsidian: {Ore: parsedValues[3], Clay: parsedValues[4]},
			Geode:    {Ore: parsedValues[5], Obsidian: parsedValues[6]},
		},
	}

}

// Read the input as a list of Blueprints
func readInput(path string) []Blueprint {
	input := []Blueprint{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseBlueprint(line))
	}

	return input
}
