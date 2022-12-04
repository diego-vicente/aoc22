package day04

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 2

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 04-01 test returned %d, got %d instead`, want, result)
	}
}
