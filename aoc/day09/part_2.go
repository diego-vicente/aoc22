package day09

// Solve the second part using a larger rope
func solveSecondPart(path string) int {
	input := readInput(path)

	rope := NewRope(10)
	for _, movement := range input {
		rope.Move(movement)
	}

	return rope.tailVisited.Size()
}
