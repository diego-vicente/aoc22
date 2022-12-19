package day19

import (
	"fmt"
	"math"
)

// A local status of a search is the inventory and number of robots
type LocalStatus struct {
	Inventory map[Material]int
	Robots    map[Material]int
}

// Create a new LocalStatus as the initial status of a search
func NewLocalStatus() LocalStatus {
	state := LocalStatus{
		Inventory: map[Material]int{Ore: 0, Clay: 0, Obsidian: 0, Geode: 0},
		Robots:    map[Material]int{Ore: 1, Clay: 0, Obsidian: 0, Geode: 0},
	}
	return state
}

// Create a new LocalStatus from the values of another one
func LocalStatusFrom(from LocalStatus) LocalStatus {
	state := LocalStatus{
		Inventory: map[Material]int{},
		Robots:    map[Material]int{},
	}

	for k, v := range from.Inventory {
		state.Inventory[k] = v
	}

	for k, v := range from.Robots {
		state.Robots[k] = v
	}

	return state
}

// Return the maximum number of geodes possible the current status
func (ls LocalStatus) MostPromising(timeLeft int) int {
	nRobots := ls.Robots[Geode]
	inventory := ls.Inventory[Geode]

	for i := timeLeft; i > 0; i-- {
		nRobots++
		inventory += nRobots
	}

	return inventory
}

// Return if a given requirements are met in the local status
func canBuild(ls LocalStatus, requirements Requirements) bool {
	for material, required := range requirements {
		// Substract the number of robots, since those units are not available
		if ls.Inventory[material]-ls.Robots[material] < required {
			return false
		}
	}

	return true
}

// Get the signature that identifies a given local status
//
// Generates a uniquely identifying string, used for memoization in the global
// status of a search
func getSignature(ls LocalStatus, timeLeft int) string {
	signature := fmt.Sprintf(".%d", timeLeft)

	for _, material := range []Material{Ore, Clay, Obsidian, Geode} {
		signature += fmt.Sprintf(
			".%d.%d.", ls.Inventory[material], ls.Robots[material],
		)
	}

	return signature
}

// The global status of a search includes the input memoization and other
// heuristics
type GlobalStatus struct {
	Blueprint Blueprint
	Limit     map[Material]int
	BestKnown int
	memo      map[string]int
}

// Create a new GlobalStatus from a given input
func NewGlobalStatus(bp Blueprint) GlobalStatus {
	limits := map[Material]int{}

	// Find the maximum number of extracting robots worth per material
	for _, material := range []Material{Ore, Clay, Obsidian} {
		limits[material] = math.MinInt

		for _, requirements := range bp.Robot {
			if required, ok := requirements[material]; ok {
				if required > limits[material] {
					limits[material] = required
				}
			}
		}
	}

	return GlobalStatus{
		Limit:     limits,
		Blueprint: bp,
		BestKnown: 0,
		memo:      map[string]int{},
	}
}

// Recursively run a Depth-First Search to find the optimum
func DepthFirstSearch(ls LocalStatus, timeLeft int, gs *GlobalStatus) int {
	// Return if the time is up
	if timeLeft <= 0 {
		result := ls.Inventory[Geode]

		if gs.BestKnown < result {
			gs.BestKnown = result
		}

		return result
	}

	// Check if it is a known branching
	signature := getSignature(ls, timeLeft)
	if result, ok := gs.memo[signature]; ok {
		return result
	}

	// Return if the branching is not promising anymore
	if gs.BestKnown > ls.MostPromising(timeLeft) {
		return -1
	}

	// Update the inventory
	for _, material := range []Material{Ore, Clay, Obsidian, Geode} {
		ls.Inventory[material] += ls.Robots[material]
	}

	// Review each of the possible robots to build
	futures := []int{}
	for _, material := range []Material{Geode, Obsidian, Clay, Ore} {
		// Check if it possible to build it
		possible := canBuild(ls, gs.Blueprint.Robot[material])
		// Check if it is needed
		needed := true
		if l, ok := gs.Limit[material]; ok && ls.Robots[material] >= l {
			needed = false
		}

		// If not possible or not needed, skip this option
		if !possible || !needed {
			continue
		}

		// Make copies of the current structures
		futureState := LocalStatusFrom(ls)

		// Consume the materials required
		for m, required := range gs.Blueprint.Robot[material] {
			futureState.Inventory[m] -= required
		}

		// Update the robot counts
		futureState.Robots[material] += 1

		// Branch to the future where a robot is built
		futures = append(
			futures,
			DepthFirstSearch(futureState, timeLeft-1, gs),
		)
	}

	// Branch to the action where a robot is built
	futures = append(futures, DepthFirstSearch(LocalStatusFrom(ls), timeLeft-1, gs))

	// Find the best case out of this status
	bestCase := 0
	for _, future := range futures {
		if future > bestCase {
			bestCase = future
		}
	}

	// Memoize before returning
	gs.memo[signature] = bestCase
	return bestCase
}

// Find the maximum number of geodes that can be extracted with a Blueprint
func FindOptimum(bp Blueprint, time int) int {
	global := NewGlobalStatus(bp)
	initialState := NewLocalStatus()
	return DepthFirstSearch(initialState, time, &global)
}

// Solve the first part by finding the quality level
func solveFirstPart(path string) int {
	result := 0

	for _, blueprint := range readInput(path) {
		result += blueprint.Id * FindOptimum(blueprint, 24)
	}

	return result
}
