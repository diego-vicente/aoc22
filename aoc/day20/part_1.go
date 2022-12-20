package day20

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// Get the final coordinates from the Ring
func getCoordinates(ring dsa.Ring[int]) int {
	zeroIdx := ring.FindIndex(0)

	return ring.GetValue(1000+zeroIdx) +
		ring.GetValue(2000+zeroIdx) +
		ring.GetValue(3000+zeroIdx)
}

// Solve the first part moving the numbers
func solveFirstPart(path string) int {
	input := readInput(path)

	for _, node := range input.OriginalNodes() {
		input.MoveNode(node, node.Value)
	}

	return getCoordinates(input)
}
