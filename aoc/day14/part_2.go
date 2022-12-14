package day14

// Drop sand to the cave until it falls to the floor (cave.Depth + 2)
func (cave *Cave) dropGrainToFloor(origin Point) {
	grain := origin
	resting := false

	// As long as the grains are falling before finding the floor (a horizontal
	// line at cave.Depth + 2), let them fall
	for !resting && grain.Y < cave.Depth+1 {
		resting = grain.fall(*cave)
	}

	cave.Scan[grain] = Sand
}

// Drop grains of sand from origin until they fill the cave
func (cave *Cave) fillWithSand(origin Point) int {
	grains := 0

	for cave.isEmtpy(origin) {
		cave.dropGrainToFloor(origin)
		// cave.Display(490, 510)
		grains++
	}

	return grains
}

// Solve the second part by filling the cave with sand
func solveSecondPart(path string) int {
	cave := readInput(path)
	return cave.fillWithSand(Point{500, 0})
}
