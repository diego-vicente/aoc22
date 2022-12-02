package day01

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 24000

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 01-01 test returned %d, got %d instead`, want, result)
	}
}

func TestSecondPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 45000

	if result := solveSecondPart(path); result != want {
		t.Fatalf(`Day 01-01 test returned %d, got %d instead`, want, result)
	}
}
