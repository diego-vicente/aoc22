package day05

// Move all the crates in a instruction in the original order
func moveInOrder(stacks CrateStacks, movement Movement) CrateStacks {
	current := []int{}

	for i := 1; i <= movement.Amount; i++ {
		current = append(current, stacks[movement.From].Pop())
	}

	for i := (movement.Amount - 1); i >= 0; i-- {
		stacks[movement.To].Push(current[i])
	}

	return stacks
}

// Return the top crates after applying all the movements in the input
func solveSecondPart(path string) string {
	stacks, movements := readInput(path)

	for _, movement := range movements {
		stacks = moveInOrder(stacks, movement)
	}

	return topCrates(stacks)
}
