package day02

// Compute the points associated with a move
func (shape HandShape) Points() int {
	switch shape {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic("Impossible move")
	}
}

// Compute the points associated with a result
func (result Result) Points() int {
	// Count the match result points
	switch result {
	case Defeat:
		return 0
	case Draw:
		return 3
	case Victory:
		return 6
	default:
		panic("Impossible result")
	}
}

// Compute the result of a match for the player
func (match Match) PlayerResult() Result {
	diff := (match.Player.Points() - match.Opponent.Points() + 3) % 3

	switch diff {
	case 0:
		return Draw
	case 1:
		return Victory
	case 2:
		return Defeat
	default:
		panic("Impossible case")
	}
}

// Compute the score for the player
func (match Match) PlayerScore() int {
	// Count the player movement points
	points := match.Player.Points()

	// Count the result points
	points += match.PlayerResult().Points()

	return points
}

// Return the total number of points for the player
func solveFirstPart(path string) int {
	var result int

	input := readFirstPartInput(path)

	for _, match := range input {
		result += match.PlayerScore()
	}

	return result
}
