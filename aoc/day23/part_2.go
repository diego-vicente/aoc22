package day23

// Solve the second part by letting the rounds converge
func solveSecondPart(path string) int {
	input := readInput(path)
	round, stable := 0, false

	for !stable {
		stable = input.RunRound(round)
		round++
	}

	return round
}
