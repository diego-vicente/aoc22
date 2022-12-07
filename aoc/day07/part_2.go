package day07

import (
	"math"
)

// Return the sum of all children Folders up to a given individual limit
func (folder Folder) DeleteCandidate(needed int) (int, bool) {
	candidate := math.MaxInt

	// Find the children's best candidate
	for _, child := range folder.Children {
		if child, ok := child.(*Folder); ok {
			newCandidate, ok := child.DeleteCandidate(needed)
			if ok && newCandidate < candidate {
				candidate = newCandidate
			}
		}
	}

	// Check if the current folder is a better candidate itself
	if self := folder.Size(); self >= needed && self < candidate {
		candidate = self
	}

	return candidate, candidate >= needed
}

// Find the candidate to be deleted for a given update
func deleteNeeded(totalDisk int, needed int, root Folder) int {
	available := totalDisk - root.Size()
	delete, ok := root.DeleteCandidate(needed - available)
	if !ok {
		panic("Unable to find a suitable candidate")
	}
	return delete
}

// Solve the second part by finding the smalles directory to free enough space
func solveSecondPart(path string) int {
	input := readInput(path)
	return deleteNeeded(70000000, 30000000, input)
}
