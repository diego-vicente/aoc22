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
	Knots       []Point
	tailVisited dsa.Set[Point]
}

// Create a Rope with a given number of knots
func NewRope(nKnots int) Rope {
	var knots []Point

	for i := 0; i < nKnots; i++ {
		knots = append(knots, Point{0, 0})
	}

	return Rope{
		Knots:       knots,
		tailVisited: dsa.NewSet[Point](),
	}
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

// Check if two knots are too far apar
func areStretched(first, tail Point) bool {
	xDelta := first.X - tail.X
	yDelta := first.Y - tail.Y
	return !(xDelta <= 1 && xDelta >= -1 && yDelta <= 1 && yDelta >= -1)
}

// Ensure that all deltas are in the [-1, 1] interval
func normalizeDeltas(xDelta, yDelta int) (int, int) {
	if xDelta > 0 {
		xDelta /= xDelta
	} else if xDelta < 0 {
		xDelta /= -xDelta
	}

	if yDelta > 0 {
		yDelta /= yDelta
	} else if yDelta < 0 {
		yDelta /= -yDelta
	}

	return xDelta, yDelta
}

// Move the tail to the next position defined by its head
func compress(first, tail Point) Point {
	xDelta, yDelta := normalizeDeltas(first.X-tail.X, first.Y-tail.Y)

	return Point{
		X: tail.X + xDelta,
		Y: tail.Y + yDelta,
	}
}

// Move a rope according to one Movement, and update its knots
func (rope *Rope) Move(movement Movement) {
	for i := 0; i < movement.Times; i++ {
		// Move the head of the rope
		rope.Knots[0] = rope.Knots[0].Move(movement.To)

		for i := 1; i < len(rope.Knots); i++ {
			if areStretched(rope.Knots[i-1], rope.Knots[i]) {
				rope.Knots[i] = compress(rope.Knots[i-1], rope.Knots[i])
			} else {
				// If there is no update, stop checking the rope
				break
			}
		}

		rope.tailVisited.Add(rope.Knots[len(rope.Knots)-1])
	}
}

// Solve the first part by counting the tail's positions
func solveFirstPart(path string) int {
	input := readInput(path)

	rope := NewRope(2)
	for _, movement := range input {
		rope.Move(movement)
	}

	return rope.tailVisited.Size()
}
