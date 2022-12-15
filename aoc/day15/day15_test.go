package day15

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 26

	input := readInput(path)
	if result := coverageOnRow(input, 10).Size(); result != want {
		t.Fatalf(`Day 15-01 test returned %d, got %d instead`, want, result)
	}
}

func TestSecondPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 56000011

	input := readInput(path)
	if result := tuningFrequency(findPossibleSpot(input, 20)); result != want {
		t.Fatalf(`Day 15-01 test returned %d, got %d instead`, want, result)
	}
}
