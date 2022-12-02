package day01

// Return the sum of all the elements of an inventory
func totalCalories(inv Inventory) int {
	result := 0

	for _, elem := range inv {
		result += elem
	}

	return result
}

// Return the maximum number of calories carried by a single elf.
func solveFirstPart(path string) int {
	var result int

	input := readInput(path)
	for _, inventory := range input {
		if total := totalCalories(inventory); total > result {
			result = total
		}
	}

	return result
}
