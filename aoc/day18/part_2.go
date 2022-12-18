package day18

import (
	"math"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// Compute the Bounding Box of a set of points
func ComputeBoundingBox(droplet dsa.Set[Point]) BoundingBox {
	bounds := BoundingBox{
		Min: Point{math.MaxInt, math.MaxInt, math.MaxInt},
		Max: Point{math.MinInt, math.MinInt, math.MinInt},
	}

	for _, p := range droplet.Values() {
		if p.X < bounds.Min.X {
			bounds.Min.X = p.X
		}
		if p.X > bounds.Max.X {
			bounds.Max.X = p.X
		}

		if p.Y < bounds.Min.Y {
			bounds.Min.Y = p.Y
		}
		if p.Y > bounds.Max.Y {
			bounds.Max.Y = p.Y
		}

		if p.Z < bounds.Min.Z {
			bounds.Min.Z = p.Z
		}
		if p.Z > bounds.Max.Z {
			bounds.Max.Z = p.Z
		}
	}

	return bounds
}

// Compute if a Point is within the bounds taking into account a tolerance
func (bb BoundingBox) WithinBounds(p Point, tolerance int) bool {
	return (p.X+tolerance) >= bb.Min.X && (p.X-tolerance) <= bb.Max.X &&
		(p.Y+tolerance) >= bb.Min.Y && (p.Y-tolerance) <= bb.Max.Y &&
		(p.Z+tolerance) >= bb.Min.Z && (p.Z-tolerance) <= bb.Max.Z
}

// Compute the total surface of the Bounding Box
//
// If padding is not zero, it will be incremented to each of the box's side
// lengths on each end.
func (bb BoundingBox) Surface(padding int) int {
	xLength := bb.Max.X - bb.Min.X + 2*padding + 1
	yLength := bb.Max.Y - bb.Min.Y + 2*padding + 1
	zLength := bb.Max.Z - bb.Min.Z + 2*padding + 1

	baseArea := xLength * yLength
	basePerimeter := xLength*2 + yLength*2

	return baseArea*2 + basePerimeter*zLength
}

// Return the set of Points of water covering a Droplet
//
// Recursively compute all exterior points of a Droplet that are contained
// within a bounding box of padding=1 from the original lava droplet bounds (so
// that the water can wrap around the edges and cover completely the shape).
func ComputeWater(droplet Droplet) dsa.Set[Point] {
	water := dsa.NewSet[Point]()

	// Choose the origin out of the droplet
	origin := Point{
		X: droplet.Bounds.Min.X - 1,
		Y: droplet.Bounds.Min.Y - 1,
		Z: droplet.Bounds.Min.Z - 1,
	}

	// Keep track of the next points to be flooded
	candidates := dsa.NewSet[Point]()
	candidates.Add(origin)

	// Consume each candidate adding the to the water set
	for candidates.Size() > 0 {
		p := candidates.Values()[0]
		candidates.Remove(p)
		water.Add(p)

		// Add the neighbors to the set of candidate points if they are not
		// water, lava or out of bounds
		for _, neighbor := range p.Neighbors() {
			if !droplet.Points.Contains(neighbor) &&
				!water.Contains(neighbor) &&
				droplet.Bounds.WithinBounds(neighbor, 1) {
				candidates.Add(neighbor)
			}
		}
	}

	return water
}

// Solve the second part using water to wrap the droplet
func solveSecondPart(path string) int {
	input := readInput(path)

	// Compute the body of water surrounding the droplet
	water := ComputeWater(input)

	// Its exterior surface is the difference between the water exposed surface
	// and the outside of the water's bounding box
	exposed := ApproximateSurface(water)
	outside := input.Bounds.Surface(1)
	return exposed - outside
}
