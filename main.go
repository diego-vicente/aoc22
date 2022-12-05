package main

import (
	"flag"
	"fmt"

	"github.com/diego-vicente/aoc22/aoc/day01"
	"github.com/diego-vicente/aoc22/aoc/day02"
	"github.com/diego-vicente/aoc22/aoc/day03"
	"github.com/diego-vicente/aoc22/aoc/day04"
	"github.com/diego-vicente/aoc22/aoc/day05"
)

func main() {
	day := flag.Int("day", 1, "day to solve")
	path := flag.String("input", "", "path to input file")

	flag.Parse()

	fmt.Println("Welcome to Advent of Code 2022!")
	fmt.Println("You have asked to solve day:", *day)

	switch *day {
	case 1:
		day01.Solve(*path)
	case 2:
		day02.Solve(*path)
	case 3:
		day03.Solve(*path)
	case 4:
		day04.Solve(*path)
	case 5:
		day05.Solve(*path)
	default:
		fmt.Printf("Day %d has not been implemented yet.\n", *day)
	}
}
