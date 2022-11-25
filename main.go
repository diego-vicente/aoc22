package main

import (
	"flag"
	"fmt"
)

func main() {
	day := flag.Int("day", 1, "day to solve")
	path := flag.String("input", "", "path to input file")

	flag.Parse()

	fmt.Println("Welcome to Advent of Code 2022!")
	fmt.Println("You have asked to solve day:", *day)

	switch *day {
	// case 1:
	// 	aoc.SolveDay01(*path)
	default:
		fmt.Printf("Day %d has not been implemented yet.", day)
	}
}
