package day20

// Solve the second part by decrypting as instructed
func solveSecondPart(path string) int {
	input := readInput(path)

	// Multiply by decryption key
	for _, node := range input.OriginalNodes() {
		node.Value *= 811589153
	}

	// Move all numbers 10 times
	for i := 0; i < 10; i++ {
		for _, node := range input.OriginalNodes() {
			input.MoveNode(node, node.Value)
		}
	}

	return getCoordinates(input)
}
