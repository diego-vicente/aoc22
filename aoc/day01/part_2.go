package day01

// A ranking is a sorted slice of integers.
type ranking = []int

// For a given sorted ranking, insert the new element in the correct order (if
// it is larger than any of the values)
func updateRanking(ranking *ranking, newValue int) {
	candidate := newValue

	for idx, value := range *ranking {
		if candidate > value {
			(*ranking)[idx] = candidate
			candidate = value
		}
	}
}

// Return the maximum number of calories carried by a the top-3 elves.
func solveSecondPart(path string) int {
	var result int
	top := ranking{0, 0, 0}

	input := readInput(path)
	for _, inventory := range input {
		updateRanking(&top, totalCalories(inventory))
	}

	for _, value := range top {
		result += value
	}

	return result
}
