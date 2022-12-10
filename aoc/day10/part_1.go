package day10

import (
	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A CPU has a register to modify and a cycle count
type CPU struct {
	Register   int
	CycleCount int
}

// A executable is any struct that defines how to be executed
type Executable interface {
	RunOn(*CPU)
}

// To run a NoOp, simply increase the cycle count
func (op NoOp) RunOn(cpu *CPU) {
	cpu.CycleCount++
}

// To run an AddX, increase the counter and then add the value to the register
func (op AddX) RunOn(cpu *CPU) {
	cpu.CycleCount++
	cpu.Register += op.X
}

// Run a set of instructions and check signal strength at given checkpoints
func (cpu *CPU) Run(instructions []Executable, checkpoints dsa.Set[int]) int {
	signalStrength := 0

	for _, instruction := range instructions {
		instruction.RunOn(cpu)

		if checkpoints.Contains(cpu.CycleCount) {
			signalStrength += cpu.Register * cpu.CycleCount
		}
	}

	return signalStrength
}

// Solve the first part by checking the signal strength
func solveFirstPart(path string) int {
	input := readInput(path)
	checkpoints := dsa.NewSetFrom([]int{20, 60, 100, 140, 180, 220})
	cpu := CPU{Register: 1, CycleCount: 1}
	return cpu.Run(input, checkpoints)
}
