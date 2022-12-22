package day22

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A Face is represented by their 2D coordinates on the input
type Face = Point

// Get the next neighbor of an Agent in a the cube
func (jungle Jungle) get3DNeighbor(a Agent) (Point, Direction) {
	current := a.Position
	target := current.getNeighbor(a.Orientation)

	if _, ok := jungle.Map[target]; ok {
		// Move within the map
		return target, a.Orientation
	} else {
		// Warp to a different face of the cube
		s := jungle.Side
		normalized := Point{dsa.Mod(current.X, s), dsa.Mod(current.Y, s)}
		face := Face{
			X: current.X / jungle.Side,
			Y: current.Y / jungle.Side,
		}

		var finalOrientation Direction
		var finalPosition Point

		switch a.Orientation {
		case North:
			switch face.X {
			case 0: // Face{0, 2}
				finalOrientation = East
				finalPosition = Point{1 * s, 1*s + normalized.X}
			case 1: // Face{1, 0}
				finalOrientation = East
				finalPosition = Point{0 * s, 3*s + normalized.X}
			case 2: // Face{2, 0}
				finalOrientation = North
				finalPosition = Point{0*s + normalized.X, 3*s + (s - 1)}
			}
		case East:
			switch face.Y {
			case 0: // Face{2, 0}
				finalOrientation = West
				finalPosition = Point{1*s + (s - 1), 2*s + (s - 1 - normalized.Y)}
			case 1: // Face{1, 1}
				finalOrientation = North
				finalPosition = Point{2*s + normalized.Y, 0*s + (s - 1)}
			case 2: // Face{1, 2}
				finalOrientation = West
				finalPosition = Point{2*s + (s - 1), 0*s + (s - 1 - normalized.Y)}
			case 3: // Face{0, 3}
				finalOrientation = North
				finalPosition = Point{1*s + normalized.Y, 2*s + (s - 1)}
			}
		case South:
			switch face.X {
			case 0: // Face{0, 3}
				finalOrientation = South
				finalPosition = Point{2*s + normalized.X, 0 * s}
			case 1: // Face{1, 2}
				finalOrientation = West
				finalPosition = Point{0*s + (s - 1), 3*s + normalized.X}
			case 2: // Face{2, 0}
				finalOrientation = West
				finalPosition = Point{1*s + (s - 1), 1*s + normalized.X}
			}
		case West:
			switch face.Y {
			case 0: // Face{1, 0}
				finalOrientation = East
				finalPosition = Point{0 * s, 2*s + (s - 1 - normalized.Y)}
			case 1: // Face{1, 1}
				finalOrientation = South
				finalPosition = Point{0*s + normalized.Y, 2 * s}
			case 2: // Face{0, 2}
				finalOrientation = East
				finalPosition = Point{1 * s, 0*s + (s - 1 - normalized.Y)}
			case 3: // Face{0, 3}
				finalOrientation = South
				finalPosition = Point{1*s + normalized.Y, 0 * s}
			}
		}

		return finalPosition, finalOrientation
	}
}

// Move the agent forward on the cube surface a given number of steps
func (a *Agent) MoveAlongCube(steps int, jungle Jungle) {
	for i := 1; i <= steps; i++ {
		target, orientation := jungle.get3DNeighbor(*a)

		switch jungle.getTerrain(target) {
		case Empty:
			a.Position = target
			a.Orientation = orientation
		case Wall:
			return
		default:
			panic("Invalid terrain moving forward")
		}
	}
}

// Perform a given instruction in the second part
func (a *Agent) PerformInCube(instruction Instruction, jungle Jungle) {
	if instruction.Direction != 0 {
		a.Turn(instruction.Direction)
	} else {
		a.MoveAlongCube(instruction.Amount, jungle)
	}
}

// Solve the second part using the cube map
func solveSecondPart(path string) int {
	jungle, instructions := readInput(path)
	agent := NewAgent(jungle)

	// Run the instructions
	for _, instruction := range instructions {
		agent.PerformInCube(instruction, jungle)
	}

	return agent.GetPassword()
}
