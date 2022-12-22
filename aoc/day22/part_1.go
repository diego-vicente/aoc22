package day22

import (
	"fmt"
	"math"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// Return the opposite direction of any given one
func Backwards(d Direction) Direction {
	return (d + 2) % 4
}

// An Agent has a position and an orientation
type Agent struct {
	Position    Point
	Orientation Direction
}

// Create a new agent in the starting point
func NewAgent(jungle Jungle) Agent {
	// Infer the starting position
	originX := math.MaxInt
	for point := range jungle.Map {
		if point.Y == 0 && point.X < originX {
			originX = point.X
		}
	}

	return Agent{
		Position:    Point{originX, 0},
		Orientation: East,
	}
}

// Turn the orientation of an agent in a given direction
func (a *Agent) Turn(d Direction) {
	a.Orientation = dsa.Mod(a.Orientation+d, 4)
}

// Get the neighbor of a point in a grid for a given direction
func (p Point) getNeighbor(d Direction) Point {
	deltas := map[Direction]Point{
		North: {0, -1},
		East:  {1, 0},
		South: {0, 1},
		West:  {-1, 0},
	}

	delta, ok := deltas[d]
	if !ok {
		panic(fmt.Sprintf("Unknown direction: %d", d))
	}

	return Point{
		X: p.X + delta.X,
		Y: p.Y + delta.Y,
	}
}

// Get the Terrain value of a point in a given Jungle
func (j *Jungle) getTerrain(p Point) Terrain {
	value, ok := j.Map[p]
	if ok {
		return value
	} else {
		return Nothing
	}
}

// Get the next neighbor of an Agent in a jungle
func (jungle Jungle) getNeighbor(a Agent) Point {
	current := a.Position
	target := current.getNeighbor(a.Orientation)

	if _, ok := jungle.Map[target]; ok {
		// Move within the map
		return target
	} else {
		// Wrap around the edges of the map
		back := Backwards(a.Orientation)
		target = current.getNeighbor(back)
		_, withinBounds := jungle.Map[target]

		for withinBounds {
			current = target
			target = current.getNeighbor(back)
			_, withinBounds = jungle.Map[target]
		}

		return current
	}
}

// Move the agent forward a given number of steps
func (a *Agent) MoveForward(steps int, jungle Jungle) {
	for i := 1; i <= steps; i++ {
		target := jungle.getNeighbor(*a)

		switch jungle.getTerrain(target) {
		case Empty:
			a.Position = target
		case Wall:
			return
		default:
			panic("Invalid terrain moving forward")
		}
	}
}

// Perform a given instruction
func (a *Agent) Perform(instruction Instruction, jungle Jungle) {
	if instruction.Direction != 0 {
		a.Turn(instruction.Direction)
	} else {
		a.MoveForward(instruction.Amount, jungle)
	}
}

// Get the password value of a given Agent state
func (a *Agent) GetPassword() int {
	return 1000*(a.Position.Y+1) + 4*(a.Position.X+1) + int(a.Orientation)
}

// Solve the first part by running the instructions
func solveFirstPart(path string) int {
	jungle, instructions := readInput(path)
	agent := NewAgent(jungle)

	// Run the instructions
	for _, instruction := range instructions {
		agent.Perform(instruction, jungle)
	}

	return agent.GetPassword()
}
