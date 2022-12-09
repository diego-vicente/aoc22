package day09

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A point is a position in a 2D space
type Point struct {
	X int
	Y int
}

// A rope is defined by its head's and tail's positions
type Rope struct {
	Head        Point
	Tail        Point
	tailVisited dsa.Set[Point]
}

// Compute the final position of moving the point to any position
func (point Point) Move(direction Direction) Point {
	switch direction {
	case Left:
		return Point{point.X - 1, point.Y}
	case Right:
		return Point{point.X + 1, point.Y}
	case Up:
		return Point{point.X, point.Y + 1}
	case Down:
		return Point{point.X, point.Y - 1}
	default:
		panic("Impossible move")
	}
}

// Check if a Rope is in a stretched position
func (rope Rope) IsStretched() bool {
	xDelta := rope.Head.X - rope.Tail.X
	yDelta := rope.Head.Y - rope.Tail.Y
	return !(xDelta <= 1 && xDelta >= -1 && yDelta <= 1 && yDelta >= -1)
}

// Move a rope according to one Movement, and update its tail
func (rope *Rope) Move(movement Movement) {
	for i := 0; i < movement.Times; i++ {
		origin := rope.Head
		rope.Head = rope.Head.Move(movement.To)
		if rope.IsStretched() {
			rope.Tail = origin
		}
		rope.tailVisited.Add(rope.Tail)
	}
}

// Solve the first part by counting the tail's positions
func solveFirstPart(path string) int {
	input := readInput(path)

	rope := Rope{
		Head:        Point{0, 0},
		Tail:        Point{0, 0},
		tailVisited: dsa.NewSet[Point](),
	}

	for _, movement := range input {
		rope.Move(movement)
	}

	return rope.tailVisited.Size()
}
