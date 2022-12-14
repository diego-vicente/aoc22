package day14

import "fmt"

// Check if a given Point in the cave is empty or not
func (cave Cave) isEmtpy(point Point) bool {
	_, full := cave.Scan[point]
	return !full
}

// Drop a grain of Sand to the Cave and return if its resting or not
func (grain *Point) fall(cave Cave) bool {
	if cave.isEmtpy(Point{grain.X, grain.Y + 1}) {
		// The grain is falling down
		grain.Y += 1
		return false
	} else if cave.isEmtpy(Point{grain.X - 1, grain.Y + 1}) {
		// If not possible, down to the left
		grain.X -= 1
		grain.Y += 1
		return false
	} else if cave.isEmtpy(Point{grain.X + 1, grain.Y + 1}) {
		// If not possible, down to the right
		grain.X += 1
		grain.Y += 1
		return false
	} else {
		// Otherwise, the grain is resting
		return true
	}
}

// Drop sand to the cave until it falls through the rock
func (cave *Cave) dropGrain(origin Point) bool {
	grain := origin

	// As long as the grains are resting in the known scan, keep on dropping
	// grains of sand from the origin point
	for grain.Y <= cave.Depth {
		resting := grain.fall(*cave)

		if resting {
			cave.Scan[grain] = Sand
			return false
		}
	}

	// If the grain fell through, its because the cave is full
	return true
}

// Drop grains of sand from origin until they fall through the cave
func (cave *Cave) dropSand(origin Point) int {
	grains := 0

	for {
		if full := cave.dropGrain(origin); !full {
			// cave.Display(494, 503)
			grains++
		} else {
			break
		}
	}

	return grains
}

// Display the Cave on a given range of X
func (cave Cave) Display(fromX, toX int) {
	for y := 0; y <= cave.Depth+2; y++ {
		for x := fromX; x <= toX; x++ {
			p := Point{x, y}
			if cave.isEmtpy(p) {
				fmt.Print(".")
			} else {
				fmt.Print(cave.Scan[p])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Solve the first part by dropping sand until it falls through
func solveFirstPart(path string) int {
	cave := readInput(path)
	return cave.dropSand(Point{500, 0})
}
