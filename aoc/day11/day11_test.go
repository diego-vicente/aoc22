package day11

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 10605

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 11-01 test returned %d, got %d instead`, want, result)
	}
}
