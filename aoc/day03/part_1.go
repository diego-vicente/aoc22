package day03

// Return the total priority the common items in both compartments
func solveFirstPart(path string) int {
	var result int

	input := readInput(path)

	for _, rucksack := range input {
		common := rucksack.First.Intersection(rucksack.Second)
		result += common.Values()[0]
	}

	return result
}
