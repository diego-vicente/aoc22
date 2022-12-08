package day03

import "github.com/diego-vicente/aoc22/aoc/dsa"

// Return the total priority the common items in both compartments
func solveSecondPart(path string) int {
	var result int

	input := readInput(path)

	var currentGroup dsa.Set[int]
	groupCounter := 0

	for _, rucksack := range input {
		// An elf's items are all those in their rucksack.
		elf := rucksack.First.Union(rucksack.Second)
		groupCounter += 1

		// If this is the first elf of the group, take it as the first value and
		// continue
		if groupCounter == 1 {
			currentGroup = elf
			continue
		}

		// Keep only the common items of the current elf and the known group
		currentGroup = currentGroup.Intersection(elf)

		// If the group is complete, take the badge and add it to the result.
		// Also, reset the group counter for the next group.
		if groupCounter == 3 {
			badge := currentGroup.Values()[0]
			result += badge
			groupCounter = 0
		}
	}

	return result
}
