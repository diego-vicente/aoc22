package day12

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// getReverseNeighbors checks height going backwards
func getReverseNeighbors(p Point, hill HillMap) []Point {
	directions := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	neighbors := []Point{}

	for _, delta := range directions {
		candidate := Point{p.X + delta.X, p.Y + delta.Y}

		// Check the point exists in the map
		height, ok := hill.Height[candidate]
		if !ok {
			continue
		}

		// Check that the height difference is admissible
		diff := hill.Height[p] - height
		if diff <= 1 {
			neighbors = append(neighbors, candidate)
		}
	}

	return neighbors
}

// Breadth-First Search traverse the map from the end to get to the first 'a'
func BreadthFirstSearch(hill HillMap, expand ExpandFunc) (Route, bool) {
	route := Route{Path: []Point{}}

	// Set of all the nodes that have already been explored
	explored := dsa.NewSet[Point]()
	explored.Add(hill.End)

	// The map of points and their antecesor in the route, will be updated
	// during execution as nodes are explored
	cameFrom := map[Point]Point{}

	// The queue of nodes to explore next, as a FIFO structure
	nodeQueue := dsa.Queue[Point]{}
	nodeQueue.Add(hill.End)

	for nodeQueue.Size() > 0 {
		// Pop the most first node to explore and check if its 'a'
		current := nodeQueue.Pop()
		if hill.Height[current] == 0 {
			return reconstructRoute(cameFrom, current), true
		}

		// Expand the node and check each neighbor
		for _, neighbor := range expand(current, hill) {
			if !explored.Contains(neighbor) {
				cameFrom[neighbor] = current
				explored.Add(neighbor)
				nodeQueue.Add(neighbor)
			}
		}

	}

	// If the queue is consumed completely, no valid route exists
	return route, false
}

// Solve the first part reaching the end point
func solveSecondPart(path string) int {
	hill := readInput(path)

	route, ok := BreadthFirstSearch(hill, getReverseNeighbors)
	if !ok {
		panic("No route was found for the input")
	}

	return len(route.Path)
}
