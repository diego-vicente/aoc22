package day17

import (
	"fmt"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A Point is represented as a 2D coordinate
type Point struct {
	X int
	Y int
}

// Convert a Direction to a delta vector for 2D
func toDelta(d Direction) Point {
	switch d {
	case Down:
		return Point{0, -1}
	case Up:
		return Point{0, 1}
	case Left:
		return Point{-1, 0}
	case Right:
		return Point{1, 0}
	}

	panic(fmt.Sprintf("Unknown direction: %s", d))
}

// A RockShape is just a list of points
type RockShape = []Point

// A Rock has a Position and a Shape
//
// The position represents the lower-left corner of the "bounding box" of the
// shape, and the position are just points relevant to such position that are
// formed by falling rock.
type Rock struct {
	Position Point
	Shape    RockShape
}

// Check if a Rock could move in a given Direction on a Chamber
func (r Rock) IsFree(c Chamber, d Direction) bool {
	delta := toDelta(d)

	for _, pos := range r.Shape {
		toCheck := Point{
			X: r.Position.X + pos.X + delta.X,
			Y: r.Position.Y + pos.Y + delta.Y,
		}

		if !c.IsEmpty(toCheck) {
			return false
		}
	}

	return true
}

// Move the Rock to a given Direction
func (r *Rock) Move(d Direction) {
	delta := toDelta(d)

	r.Position = Point{
		X: r.Position.X + delta.X,
		Y: r.Position.Y + delta.Y,
	}
}

// A Rock is falling if it can move down freely
func (r Rock) IsFalling(c Chamber) bool {
	return r.IsFree(c, Down)
}

// A Chamber has dimensions and the current state of the inside.
type Chamber struct {
	Width  int
	Height int
	Rocks  dsa.Set[Point]
	Jets   dsa.Queue[Direction]
	Shapes dsa.Queue[RockShape]
}

// Generate a new Chamber from a width and a set of jets
func NewChamber(width int, jets dsa.Queue[Direction]) Chamber {
	shapes := dsa.NewQueue[RockShape]()
	// Add the horizontal block
	shapes.Add(RockShape{{0, 0}, {1, 0}, {2, 0}, {3, 0}})
	// Add the cross
	shapes.Add(RockShape{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}})
	// Add the inverted L
	shapes.Add(RockShape{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}})
	// Add the vertical block
	shapes.Add(RockShape{{0, 0}, {0, 1}, {0, 2}, {0, 3}})
	// Add the square
	shapes.Add(RockShape{{0, 0}, {0, 1}, {1, 0}, {1, 1}})

	return Chamber{
		Width:  width,
		Height: 0,
		Rocks:  dsa.NewSet[Point](),
		Jets:   jets,
		Shapes: shapes,
	}

}

// Generate the next Rock falling in the Chamber
func (c *Chamber) NextRock() Rock {
	// Compute the origin position
	origin := Point{2, c.Height + 3}

	// Get the next shape and push it to the end of the queue
	shape := c.Shapes.Pop()
	c.Shapes.Add(shape)

	return Rock{Position: origin, Shape: shape}
}

// Get the Direction of the next jet blowing in the Chamber
func (c *Chamber) NextJet() Direction {
	// Get the next jet and push it to the end of the queue
	jet := c.Jets.Pop()
	c.Jets.Add(jet)

	return jet
}

// Check if a Point is empty (and within bounds) in the Chamber
func (c Chamber) IsEmpty(p Point) bool {
	inBounds := p.X >= 0 && p.X < c.Width && p.Y >= 0
	return inBounds && !c.Rocks.Contains(p)
}

// Fix the position of a Rock within the Chamber
func (c *Chamber) FixRock(r Rock) {
	for _, pos := range r.Shape {
		p := Point{
			X: r.Position.X + pos.X,
			Y: r.Position.Y + pos.Y,
		}

		// Add a new solid rock point
		c.Rocks.Add(p)

		// Update the chamber height if needed
		if c.Height < p.Y+1 {
			c.Height = p.Y + 1
		}
	}
}

// Throw a new Rock to the Chamber
func (chamber *Chamber) ThrowRock() {
	rock := chamber.NextRock()

	for {
		// If possible, move the rock along the jet
		jet := chamber.NextJet()
		if rock.IsFree(*chamber, jet) {
			rock.Move(jet)
		}

		if rock.IsFalling(*chamber) {
			// If the rock can move down, it keeps falling
			rock.Move(Down)
		} else {
			// If not, it is finally resting and we stop
			break
		}
	}

	// Solidify the rock and update the chamber data
	chamber.FixRock(rock)
}

// Solve the first part running 2022 rounds
func solveFirstPart(path string) int {
	input := readInput(path)
	chamber := NewChamber(7, input)

	for i := 0; i < 2022; i++ {
		chamber.ThrowRock()
	}

	return chamber.Height
}
