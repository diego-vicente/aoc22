package day01

import "fmt"

func Solve(path string) {
	input := readInput(path)

	fmt.Println("Solution to part one:", solveFirstPart(input))
	fmt.Println("Solution to part two:", solveSecondPart(input))
}
