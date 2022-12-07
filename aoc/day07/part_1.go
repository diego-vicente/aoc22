package day07

// Return the sum of all children Folders up to a given individual limit
func (folder Folder) SizeUpTo(limit int) int {
	result := 0
	current := 0

	for _, child := range folder.Children {
		// Check the child's recursively
		if child, ok := child.(*Folder); ok {
			result += child.SizeUpTo(limit)
		}

		// Also check the current folder's size
		current += child.Size()
	}

	// If it is under the limit, add itself to the result
	if current <= limit {
		result += current
	}

	return result
}

// Solve the first part adding all folder's sizes under 100000
func solveFirstPart(path string) int {
	input := readInput(path)
	return input.SizeUpTo(100000)
}
