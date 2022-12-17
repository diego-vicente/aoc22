package main

import (
	"flag"
	"fmt"

	"github.com/diego-vicente/aoc22/aoc/day01"
	"github.com/diego-vicente/aoc22/aoc/day02"
	"github.com/diego-vicente/aoc22/aoc/day03"
	"github.com/diego-vicente/aoc22/aoc/day04"
	"github.com/diego-vicente/aoc22/aoc/day05"
	"github.com/diego-vicente/aoc22/aoc/day06"
	"github.com/diego-vicente/aoc22/aoc/day07"
	"github.com/diego-vicente/aoc22/aoc/day08"
	"github.com/diego-vicente/aoc22/aoc/day09"
	"github.com/diego-vicente/aoc22/aoc/day10"
	"github.com/diego-vicente/aoc22/aoc/day11"
	"github.com/diego-vicente/aoc22/aoc/day12"
	"github.com/diego-vicente/aoc22/aoc/day13"
	"github.com/diego-vicente/aoc22/aoc/day14"
	"github.com/diego-vicente/aoc22/aoc/day15"
	"github.com/diego-vicente/aoc22/aoc/day17"
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
	case 6:
		day06.Solve(*path)
	case 7:
		day07.Solve(*path)
	case 8:
		day08.Solve(*path)
	case 9:
		day09.Solve(*path)
	case 10:
		day10.Solve(*path)
	case 11:
		day11.Solve(*path)
	case 12:
		day12.Solve(*path)
	case 13:
		day13.Solve(*path)
	case 14:
		day14.Solve(*path)
	case 15:
		day15.Solve(*path)
	case 17:
		day17.Solve(*path)
	default:
		fmt.Printf("Day %d has not been implemented yet.\n", *day)
	}
}
