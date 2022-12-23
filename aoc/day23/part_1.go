package day23

import (
	"fmt"
	"math"
)

// Get the neighbor of a Point in a given Direction
func (p Point) Neighbor(direction Direction) Point {
	switch direction {
	case North:
		return Point{p.X, p.Y - 1}
	case NorthEast:
		return Point{p.X + 1, p.Y - 1}
	case East:
		return Point{p.X + 1, p.Y}
	case SouthEast:
		return Point{p.X + 1, p.Y + 1}
	case South:
		return Point{p.X, p.Y + 1}
	case SouthWest:
		return Point{p.X - 1, p.Y + 1}
	case West:
		return Point{p.X - 1, p.Y}
	case NorthWest:
		return Point{p.X - 1, p.Y - 1}
	default:
		panic(fmt.Sprintf("Unknown direction %v", direction))
	}
}

// Is a given position empty or not
func (g Grove) IsEmpty(p Point) bool {
	_, ok := g.Elves[p]
	return !ok
}

// Return the order in which Directions should be proposed for a given round
func getCheckOrder(round int) []Direction {
	original := []Direction{North, South, West, East}

	// Offset the original list by the round number
	return []Direction{
		original[(round+0)%4],
		original[(round+1)%4],
		original[(round+2)%4],
		original[(round+3)%4],
	}
}

// Check if an Elf can stay in their current position
func (e Elf) CanStay(grove Grove) bool {
	return grove.IsEmpty(e.Position.Neighbor(North)) &&
		grove.IsEmpty(e.Position.Neighbor(NorthEast)) &&
		grove.IsEmpty(e.Position.Neighbor(East)) &&
		grove.IsEmpty(e.Position.Neighbor(SouthEast)) &&
		grove.IsEmpty(e.Position.Neighbor(South)) &&
		grove.IsEmpty(e.Position.Neighbor(SouthWest)) &&
		grove.IsEmpty(e.Position.Neighbor(West)) &&
		grove.IsEmpty(e.Position.Neighbor(NorthWest))
}

// Check if an Elf can move in a given direction
func (e Elf) CanMove(grove Grove, direction Direction) bool {
	switch direction {
	case North:
		return grove.IsEmpty(e.Position.Neighbor(North)) &&
			grove.IsEmpty(e.Position.Neighbor(NorthEast)) &&
			grove.IsEmpty(e.Position.Neighbor(NorthWest))
	case South:
		return grove.IsEmpty(e.Position.Neighbor(South)) &&
			grove.IsEmpty(e.Position.Neighbor(SouthEast)) &&
			grove.IsEmpty(e.Position.Neighbor(SouthWest))
	case West:
		return grove.IsEmpty(e.Position.Neighbor(West)) &&
			grove.IsEmpty(e.Position.Neighbor(NorthWest)) &&
			grove.IsEmpty(e.Position.Neighbor(SouthWest))
	case East:
		return grove.IsEmpty(e.Position.Neighbor(East)) &&
			grove.IsEmpty(e.Position.Neighbor(NorthEast)) &&
			grove.IsEmpty(e.Position.Neighbor(SouthEast))
	default:
		panic(fmt.Sprintf("Unknown direction %v", direction))
	}
}

// Proposals map a proposed position to the list of elves proposing it
type Proposals = map[Point][](*Elf)

// Add a given proposal to the current relation
func addProposal(elf Elf, d Direction, proposals *Proposals) {
	target := elf.Position.Neighbor(d)

	if list, ok := (*proposals)[target]; ok {
		(*proposals)[target] = append(list, &elf)
	} else {
		(*proposals)[target] = [](*Elf){&elf}
	}
}

// Move an Elf to a given position within the Grove
func (elf *Elf) Move(grove *Grove, target Point) {
	origin := elf.Position

	// Update the grove
	delete(grove.Elves, origin)
	grove.Elves[target] = elf

	// Update the elf itself
	elf.Position = target
}

// Run a given round of the elves algorithm
func (g *Grove) RunRound(round int) {
	candidates := getCheckOrder(round)
	proposals := Proposals{}

	// Ask each of the elves for a proposed position
	for _, elf := range g.Elves {
		// If there is no one around, the elf does not move
		if elf.CanStay(*g) {
			continue
		}

		// Otherwise, try to propose a direction
		for _, direction := range candidates {
			if elf.CanMove(*g, direction) {
				addProposal(*elf, direction, &proposals)
				break
			}
		}
	}

	// If the position is unique, move the elf
	for position, elves := range proposals {
		if len(elves) == 1 {
			elves[0].Move(g, position)
		}
	}
}

// Compute the bounding box of the current Grove disposition
func (g Grove) ComputeBounds() BoundingBox {
	bb := BoundingBox{
		Min: Point{math.MaxInt, math.MaxInt},
		Max: Point{math.MinInt, math.MinInt},
	}

	for point := range g.Elves {
		if point.X < bb.Min.X {
			bb.Min.X = point.X
		}
		if point.Y < bb.Min.Y {
			bb.Min.Y = point.Y
		}
		if point.X > bb.Max.X {
			bb.Max.X = point.X
		}
		if point.Y > bb.Max.Y {
			bb.Max.Y = point.Y
		}
	}

	return bb
}

// Compute the area of a given BoundingBox
func (bb BoundingBox) Area() int {
	return (bb.Max.X - bb.Min.X + 1) * (bb.Max.Y - bb.Min.Y + 1)
}

// Compute how many empty tiles are there
func (g Grove) EmtpyTiles() int {
	total := g.ComputeBounds().Area()
	return total - len(g.Elves)
}

// Display the current state of the Grove
func (g Grove) Display() {
	bb := g.ComputeBounds()

	for y := bb.Min.Y; y <= bb.Max.Y; y++ {
		for x := bb.Min.X; x <= bb.Max.X; x++ {
			if g.IsEmpty(Point{x, y}) {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Solve the first part by running ten rounds
func solveFirstPart(path string) int {
	input := readInput(path)

	for i := 0; i < 10; i++ {
		input.RunRound(i)
	}

	return input.EmtpyTiles()
}
