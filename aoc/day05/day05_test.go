package day05

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := "CMZ"

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 05-01 test returned %s, got %s instead`, want, result)
	}
}
