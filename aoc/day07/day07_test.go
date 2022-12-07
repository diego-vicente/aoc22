package day07

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 95437

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 07-01 test returned %d, got %d instead`, want, result)
	}
}

func TestSecondPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 24933642

	if result := solveSecondPart(path); result != want {
		t.Fatalf(`Day 07-02 test returned %d, got %d instead`, want, result)
	}
}
