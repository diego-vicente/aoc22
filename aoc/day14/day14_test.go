package day14

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 24

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 14-01 test returned %d, got %d instead`, want, result)
	}
}

func TestSecondPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 93

	if result := solveSecondPart(path); result != want {
		t.Fatalf(`Day 14-02 test returned %d, got %d instead`, want, result)
	}
}
