package day12

import (
	"bufio"
	"os"
)

// A point is represented by 2D coordinates
type Point struct {
	X int
	Y int
}

// A hill map has a start, end, and a point mapping to their heights
type HillMap struct {
	Start  Point
	End    Point
	Height map[Point]int
}

// Translate char-represented heights to integers
func charToHeight(char rune) int {
	return int(char) - 97
}

// Parse the input as a HillMap
func readInput(path string) HillMap {
	var start, end Point
	height := map[Point]int{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for currentY := 0; scanner.Scan(); currentY++ {
		line := scanner.Text()
		for currentX, char := range line {
			currentPoint := Point{currentX, currentY}
			switch string(char) {
			case "S":
				start = currentPoint
				height[currentPoint] = charToHeight('a')
			case "E":
				end = currentPoint
				height[currentPoint] = charToHeight('z')
			default:
				height[currentPoint] = charToHeight(char)
			}
		}
	}

	return HillMap{
		Start:  start,
		End:    end,
		Height: height,
	}
}
