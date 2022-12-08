package day08

// A height pointer stores a pointer and the top height seen
type HeightPointer struct {
	Point   Point
	TopSeen int
}

// The LOS is measures from the outside and is stored in the Forest itself
func countLineOfSight(forest *Forest, start Point, update func(Point) Point) int {
	result := 0
	ptr := HeightPointer{start, -1}

	for {
		height, ok := forest.Height[ptr.Point]
		if !ok {
			break
		}

		if height > ptr.TopSeen {
			forest.Visible[ptr.Point] = true
			ptr.TopSeen = height
		}

		ptr.Point = update(ptr.Point)
	}

	return result
}

// Return the right point
func traverseRight(p Point) Point { return Point{p.X + 1, p.Y} }

// Return the left point
func traverseLeft(p Point) Point { return Point{p.X - 1, p.Y} }

// Return the point below
func traverseDown(p Point) Point { return Point{p.X, p.Y + 1} }

// Return the point above
func traverseUp(p Point) Point { return Point{p.X, p.Y - 1} }

// Count the LOS trees from all 4 directions
func solveFirstPart(path string) int {
	forest := readInput(path)

	for y := 0; y < forest.Dimensions.Y; y++ {
		countLineOfSight(&forest, Point{0, y}, traverseRight)
		countLineOfSight(&forest, Point{forest.Dimensions.X - 1, y}, traverseLeft)
	}

	for x := 0; x < forest.Dimensions.X; x++ {
		countLineOfSight(&forest, Point{x, 0}, traverseDown)
		countLineOfSight(&forest, Point{x, forest.Dimensions.Y - 1}, traverseUp)
	}

	return len(forest.Visible)
}
