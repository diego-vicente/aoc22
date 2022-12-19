package day19

// Solve the second part by running the first 3 blueprints
func solveSecondPart(path string) int {
	input := readInput(path)

	return FindOptimum(input[0], 32) *
		FindOptimum(input[1], 32) *
		FindOptimum(input[2], 32)
}
