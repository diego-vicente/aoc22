package day01

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 24000

	input := readInput(path)
	result := solveFirstPart(input)

	if result != want {
		t.Fatalf(`Day 01-01 test returned %d, got %d instead`, want, result)
	}
}

func TestSecondPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 45000

	input := readInput(path)
	result := solveSecondPart(input)

	if result != want {
		t.Fatalf(`Day 01-01 test returned %d, got %d instead`, want, result)
	}
}
