package day15

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A Range1D is defined by its start and end
type Range1D struct {
	Start int
	End   int
}

// An IndexRange is, indeed, one-dimensional
type IndexRange = Range1D

// The size of a range is teh difference between its Start and End
func (r Range1D) Size() int {
	return dsa.Abs(r.End - r.Start)
}

// Two Range1D overalp if the former's End is larger than the latter's start
func (r1 Range1D) Overlaps(r2 Range1D) bool {
	if r1.Start < r2.Start {
		return r1.End >= r2.Start
	} else {
		return r2.End >= r1.Start
	}
}

// A MultiRange1D is a set of ranges that have gaps in the middle
type MultiRange1D struct {
	Segments []Range1D
}

// The size of a MultiRange1D is the sum of its segments
func (mr MultiRange1D) Size() int {
	size := 0

	for _, segment := range mr.Segments {
		size += segment.Size()
	}

	return size
}

// Return all segment values and indices that overlap a second different segment
func (mr MultiRange1D) Overlapping(other Range1D) (IndexRange, []Range1D) {
	overlaps := []Range1D{}
	indices := IndexRange{-1, -1}

	for i, segment := range mr.Segments {
		if segment.Overlaps(other) {
			// Update the overlapping indices
			if len(overlaps) == 0 {
				indices.Start = i
			}
			if indices.End < i {
				indices.End = i
			}

			// Add the segments to the overlap
			overlaps = append(overlaps, segment)
		}
	}

	return indices, overlaps
}

// Change a set of given segments by a new, consolidated one
func (mr *MultiRange1D) Override(newRange Range1D, dropIndices IndexRange) {
	tail := make([]Range1D, len(mr.Segments)-(dropIndices.End+1))
	copy(tail, mr.Segments[dropIndices.End+1:])

	mr.Segments = append(mr.Segments[:dropIndices.Start], newRange)
	mr.Segments = append(mr.Segments, tail...)
}

// Add a new 2DRange to an existing MultiRange2D and consolidate its segments
func (mr *MultiRange1D) Add(simple Range1D) {
	// If there are no segments, just add it
	if len(mr.Segments) == 0 {
		mr.Segments = []Range1D{simple}
		return
	}

	indices, overlaps := mr.Overlapping(simple)
	new := simple

	// Adjust the new segment to include all overlapping
	for _, overlap := range overlaps {
		if overlap.Start < new.Start {
			new.Start = overlap.Start
		}

		if overlap.End > new.End {
			new.End = overlap.End
		}
	}

	// If there are no overlaps, adjust the indices accordingly
	if indices.Start == -1 {
		cutoff := len(mr.Segments)

		// Find the correct position to keep it sorted
		for i, elem := range mr.Segments {
			if elem.Start > new.Start {
				cutoff = i
				break
			}
		}

		// Not really a range, because we don't want to drop items
		indices = IndexRange{Start: cutoff, End: cutoff - 1}
	}

	// Override all overlapping elements with the new one
	mr.Override(new, indices)
}

// Compute the Manhattan distance between two points
func manhattanDistance(p1, p2 Point) int {
	return dsa.Abs(p1.X-p2.X) + dsa.Abs(p1.Y-p2.Y)
}

// Compute the known knownCoverage of a sensor-beacon pair for a given row
func knownCoverage(sensor Sensor, beacon Beacon, y int) (Range1D, bool) {
	distance := manhattanDistance(sensor.Position, beacon.Position)
	slack := distance - dsa.Abs(y-sensor.Position.Y)

	// Return that this pair does not cover such row
	if slack < 0 {
		return Range1D{0, 0}, false
	}

	// Return the correct coverage range for the row
	coverage := Range1D{
		sensor.Position.X - slack,
		sensor.Position.X + slack,
	}

	return coverage, true
}

// Compute the known coverage for all sensor-beacon pairs
func coverageOnRow(input map[Sensor]Beacon, y int) MultiRange1D {
	totalCoverage := MultiRange1D{[]Range1D{}}

	for sensor, beacon := range input {
		sensorCoverage, ok := knownCoverage(sensor, beacon, y)
		if ok {
			totalCoverage.Add(sensorCoverage)
		}
	}

	return totalCoverage
}

// Solve the first pair by computing coverage for row 2000000
func solveFirstPart(path string) int {
	input := readInput(path)
	return coverageOnRow(input, 2000000).Size()
}
