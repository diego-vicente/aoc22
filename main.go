package main

import (
	"flag"
	"fmt"
)

func main() {
	day := flag.Int("day", 1, "day to solve")

	flag.Parse()

	fmt.Println("Welcome to Advent of Code 2022!")
	fmt.Println("You have asked to solve day:", *day)
}