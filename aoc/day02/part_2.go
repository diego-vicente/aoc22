package day02

// Infer the move that should be played for a rigged match
func (match RiggedMatch) InferMove() HandShape {
	// Generalize the result
	diff := (match.Opponent.Points() + (match.Expected.Points()/3 + 1)) % 3

	switch diff {
	case 0:
		return Rock
	case 1:
		return Paper
	case 2:
		return Scissors
	default:
		panic("Impossible result")
	}
}

// Compute the total points for a rigged match
func (match RiggedMatch) PlayerScore() int {
	// Count the result points
	points := match.Expected.Points()

	// Count the player move points
	points += match.InferMove().Points()

	return points
}

// Return the total number of points for the player
func solveSecondPart(path string) int {
	var result int

	input := readSecondPartInput(path)

	for _, match := range input {
		result += match.PlayerScore()
	}

	return result
}
