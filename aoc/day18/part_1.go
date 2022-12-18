package day18

import "github.com/diego-vicente/aoc22/aoc/dsa"

// Get all neighbors of a Point
func (p Point) Neighbors() []Point {
	deltas := []Point{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	neigbors := []Point{}

	for _, delta := range deltas {
		neighbor := Point{
			X: p.X + delta.X,
			Y: p.Y + delta.Y,
			Z: p.Z + delta.Z,
		}
		neigbors = append(neigbors, neighbor)
	}

	return neigbors
}

// Count how many sides of a Point are not adjacent to another
func countExposedSides(points dsa.Set[Point], p Point) int {
	exposed := 0

	for _, neighbor := range p.Neighbors() {
		if !points.Contains(neighbor) {
			exposed++
		}
	}

	return exposed
}

// Count all the exposed sides of a set of points
func ApproximateSurface(points dsa.Set[Point]) int {
	surface := 0
	for _, point := range points.Values() {
		surface += countExposedSides(points, point)
	}
	return surface
}

// Solve the first part by approximating the surface
func solveFirstPart(path string) int {
	input := readInput(path)
	return ApproximateSurface(input.Points)
}
