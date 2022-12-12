package day12

import (
	"container/heap"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A Route includes a set of ordered points as Path
type Route struct {
	Path []Point
}

// A expand function returns all neighbors of a point
type ExpandFunc = func(Point, HillMap) []Point

// A cost function returns the cost of moving from one point to another
type CostFunc = func(Point, Point, HillMap) int

// A heuristic function returns an admissible estimate from one point to reach the end
type HeuristicFunc = func(Point, HillMap) int

// getValidNeighbors is an ExpandFunc that returns all valid neighbors of a given point
func getValidNeighbors(p Point, hill HillMap) []Point {
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
		diff := height - hill.Height[p]
		if diff <= 1 {
			neighbors = append(neighbors, candidate)
		}
	}

	return neighbors
}

// Having a cameFrom and an end point (from an A*), reconstruct the route
func reconstructRoute(cameFrom map[Point]Point, end Point) Route {
	route := Route{Path: []Point{}}
	current := end

	// Unwind the explored nodes
	for {
		from, ok := cameFrom[current]
		if !ok {
			break
		}

		route.Path = append(route.Path, from)
		current = from
	}

	// Reverse the slice to get the original order
	for i, j := 0, len(route.Path)-1; i < j; i, j = i+1, j-1 {
		route.Path[i], route.Path[j] = route.Path[j], route.Path[i]
	}

	return route
}

// A* finds a path from start to goal.
//
// hill is the map to be traversed, expand is an expansion function that returns
// all valid neighbors, cost is the cost function from point to another and the
// heuristic must be an admissible heuristic function in order to ensure the
// optimality of the results.
func AStar(hill HillMap, expand ExpandFunc, cost CostFunc, heuristic HeuristicFunc) (Route, bool) {
	route := Route{Path: []Point{}}

	// The open set of nodes to explore, sorted by its expected cost.
	openSet := dsa.PriorityQueue[Point]{}
	startNode := dsa.Item[Point]{Value: hill.Start, Priority: heuristic(hill.Start, hill)}
	heap.Push(&openSet, startNode)

	// The map of points and their antecesor in the route, will be updated
	// during execution as nodes are explored
	cameFrom := map[Point]Point{}

	// For each point, minCost stores the cost of the cheapest path from start to it.
	minCost := map[Point]int{}
	minCost[hill.Start] = 0

	// expectedCost stores the current expected cost to reach the end from
	// start. It can be understood as how promising a point is.
	expectedCost := map[Point]int{}
	expectedCost[hill.Start] = heuristic(hill.Start, hill)

	for openSet.Len() > 0 {
		// Pop the most promising node to explore and check if its the goal
		current := heap.Pop(&openSet).(*dsa.Item[Point]).Value
		if current == hill.End {
			return reconstructRoute(cameFrom, current), true
		}

		// Expand the node and check each neighbor
		for _, neighbor := range expand(current, hill) {
			score := minCost[current] + cost(current, neighbor, hill)

			// If going through current is better option than previously known,
			// update each variable accordingly
			if known, ok := minCost[neighbor]; !ok || score < known {
				cameFrom[neighbor] = current
				minCost[neighbor] = score
				expectedCost[neighbor] = score + heuristic(neighbor, hill)

				heap.Push(
					&openSet,
					dsa.Item[Point]{
						Value:    neighbor,
						Priority: score + heuristic(neighbor, hill),
					},
				)
			}
		}
	}

	// If the open set is consumed completely, no valid route exists
	return route, false
}

// Return the same cost regarding the movement
func constantCost(_, _ Point, _ HillMap) int {
	return 1
}

// Return the Manhattan distance to the end point
func ManhattanDistance(p Point, hill HillMap) int {
	return dsa.Abs(p.X-hill.End.X) + dsa.Abs(p.Y-hill.End.Y)
}

// Solve the first part reaching the end point
func solveFirstPart(path string) int {
	hill := readInput(path)

	route, ok := AStar(hill, getValidNeighbors, constantCost, ManhattanDistance)
	if !ok {
		panic("No route was found for the input")
	}

	return len(route.Path)
}
