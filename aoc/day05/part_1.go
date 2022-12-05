package day05

// Move the crates following a Movement instruction
func move(stacks CrateStacks, movement Movement) CrateStacks {
	for i := 1; i <= movement.Amount; i++ {
		elem := stacks[movement.From].Pop()
		stacks[movement.To].Push(elem)
	}

	return stacks
}

// Return the string of the top crate in each stack
func topCrates(stacks CrateStacks) string {
	nStacks := len(stacks)

	result := ""
	for i := 1; i <= nStacks; i++ {
		result += string(rune(stacks[i].Peek()))
	}

	return result
}

// Return the top crates after all movements in the input
func solveFirstPart(path string) string {
	stacks, movements := readInput(path)

	for _, movement := range movements {
		stacks = move(stacks, movement)
	}

	return topCrates(stacks)
}
