package day06

import (
	"testing"
)

const ASSETS_DIR = "./assets/"

func TestFirstPart(t *testing.T) {
	path := ASSETS_DIR + "example.txt"
	want := 7

	if result := solveFirstPart(path); result != want {
		t.Fatalf(`Day 06-01 test returned %d, got %d instead`, want, result)
	}
}

func TestHeaderIndex1(t *testing.T) {
	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	want := 5

	if result := headerIndex(input, 4); result != want {
		t.Fatalf(`Expected %d, got %d instead`, want, result)
	}
}

func TestHeaderIndex2(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"
	want := 6

	if result := headerIndex(input, 4); result != want {
		t.Fatalf(`Expected %d, got %d instead`, want, result)
	}
}

func TestHeaderIndex3(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	want := 10

	if result := headerIndex(input, 4); result != want {
		t.Fatalf(`Expected %d, got %d instead`, want, result)
	}
}

func TestHeaderIndex4(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	want := 11

	if result := headerIndex(input, 4); result != want {
		t.Fatalf(`Expected %d, got %d instead`, want, result)
	}
}
