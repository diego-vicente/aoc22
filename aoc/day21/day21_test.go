package day21

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 152

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 21-01 test returned %d, got %d instead`, want, result)
	}
}
