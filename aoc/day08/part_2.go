package day08

// Count how many trees are visible before the view is blocked in a direction
func countVisibleFrom(forest Forest, start Point, update func(Point) Point) int {
	result := 0
	ptr := HeightPointer{start, -1}

	for {
		// Move the pointer
		ptr.Point = update(ptr.Point)

		// If we are outside of the forest, we are done
		height, ok := forest.Height[ptr.Point]
		if !ok {
			break
		}

		// Count one more tree
		result++

		// If the view is blocked, we are done
		if height >= forest.Height[start] {
			break
		}
	}

	return result
}

// Compute the scenic score by multiplying visible trees in each direction
func computeScenicScore(forest Forest, start Point) int {
	right := countVisibleFrom(forest, start, traverseRight)
	left := countVisibleFrom(forest, start, traverseLeft)
	up := countVisibleFrom(forest, start, traverseUp)
	down := countVisibleFrom(forest, start, traverseDown)

	return right * left * up * down
}

// Find the tree with the highest score
func solveSecondPart(path string) int {
	forest := readInput(path)

	topScore := 0
	for point := range forest.Height {
		if score := computeScenicScore(forest, point); score > topScore {
			topScore = score
		}
	}

	return topScore
}
