package day04

// Return if one turn completely contains another one
func (turn Turn) contains(other Turn) bool {
	return turn.Start <= other.Start && turn.End >= other.End
}

// Return the total number of pairs in which one turn contains the other
func solveFirstPart(path string) int {
	var result int

	input := readInput(path)

	for _, pair := range input {
		if pair.First.contains(pair.Second) || pair.Second.contains(pair.First) {
			result += 1
		}
	}

	return result
}
