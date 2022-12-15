package day15

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

// A Sensor is defined by its location
type Sensor struct {
	Position Point
}

// A Beacon is defined by its location
type Beacon struct {
	Position Point
}

// Parse a point from the input format
func parsePoint(rawX string, rawY string) Point {
	x, err := strconv.Atoi(rawX)
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s, %s", rawX, rawY))
	}

	y, err := strconv.Atoi(rawY)
	if err != nil {
		panic(fmt.Sprintf("Error parsing point %s, %s", rawX, rawY))
	}

	return Point{x, y}
}

// Parse a pair of Sensor and nearby Beacon from input
func parseLine(line string) (Sensor, Beacon) {
	tokens := strings.Split(line, " ")

	xRaw, yRaw := tokens[2], tokens[3]
	sensorPosition := parsePoint(xRaw[2:len(xRaw)-1], yRaw[2:len(yRaw)-1])

	xRaw, yRaw = tokens[8], tokens[9]
	beaconPosition := parsePoint(xRaw[2:len(xRaw)-1], yRaw[2:])

	return Sensor{sensorPosition}, Beacon{beaconPosition}
}

// Read the input file into a map of Sensors to Beacons
func readInput(path string) map[Sensor]Beacon {
	input := map[Sensor]Beacon{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sensor, beacon := parseLine(line)
		input[sensor] = beacon
	}

	return input
}
