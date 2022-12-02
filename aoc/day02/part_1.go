package day02

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

func (match Match) PlayerResult() Result {
	diff := (match.Player.Points() - match.Opponent.Points() + 3) % 3

	switch diff {
	case 1:
		return Victory
	case 2:
		return Defeat
	default:
		return Draw
	}
}

func (match Match) playerScore() int {
	// Count the player movement points
	points := match.Player.Points()

	// Count the match result points
	switch match.PlayerResult() {
	case Draw:
		points += 3
	case Victory:
		points += 6
	}

	return points
}

// Return the maximum number of calories carried by a single elf.
func solveFirstPart(path string) int {
	var result int

	input := readInput(path)

	for _, match := range input {
		result += match.playerScore()
	}

	return result
}
