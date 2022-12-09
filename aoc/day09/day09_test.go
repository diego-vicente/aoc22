package day09

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example_1.txt"
	want := 13

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 09-01 test returned %d, got %d instead`, want, result)
	}
}

func TestSecondPart(t *testing.T) {
	path := ASSETS_DIR + "example_2.txt"
	want := 36

	if result := solveSecondPart(path); result != want {
		t.Fatalf(`Day 09-02 test returned %d, got %d instead`, want, result)
	}
}
