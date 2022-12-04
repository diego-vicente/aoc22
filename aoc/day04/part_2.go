package day04

// Return if one turn overlaps another
func (turn Turn) overlaps(other Turn) bool {
	if turn.Start < other.Start {
		return turn.End >= other.Start
	} else {
		return other.End >= turn.Start
	}
}

// Return the total number of pairs in which the turns overlap
func solveSecondPart(path string) int {
	var result int

	input := readInput(path)

	for _, pair := range input {
		if pair.First.overlaps(pair.Second) {
			result += 1
		}
	}

	return result
}
