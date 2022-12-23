package day23

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPartMini(t *testing.T) {
	path := ASSETS_DIR + "example_1.txt"
	want := 25

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 23-01 test returned %d, got %d instead`, want, result)
	}
}

func TestFirstPartExample(t *testing.T) {
	path := ASSETS_DIR + "example_2.txt"
	want := 110

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 23-01 test returned %d, got %d instead`, want, result)
	}
}
