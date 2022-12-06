package day06

// Solve the second part using headers of length 14
func solveSecondPart(path string) int {
	input := readInput(path)
	return headerIndex(input, 14)
}
