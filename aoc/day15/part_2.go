package day15

// Find the possible spot iteratively searching rows
func findPossibleSpot(input map[Sensor]Beacon, bounds int) Point {
	// TODO: it could be done using 2D ranges?
	for y := 0; y <= bounds; y++ {
		coverage := coverageOnRow(input, y)

		// Taking advantage of the fact that just one point can be it
		if len(coverage.Segments) > 1 {
			return Point{coverage.Segments[0].End + 1, y}
		}
	}

	panic("No beacon found for that input")
}

// Compute the tuning frequency as explained in the statment
func tuningFrequency(p Point) int {
	return 4000000*p.X + p.Y
}

// Solve the second part by finding the beacon
func solveSecondPart(path string) int {
	input := readInput(path)
	location := findPossibleSpot(input, 4000000)
	return tuningFrequency(location)
}
