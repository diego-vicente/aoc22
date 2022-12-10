package day10

import (
	"fmt"

	"github.com/diego-vicente/aoc22/aoc/dsa"
)

// A CRT display has its dimensions and the set of pixels to display
type CRT struct {
	Width  int
	Height int
	Pixels dsa.Set[int]
}

// Display the current CRT pixel rows and flush the pixels
func (crt *CRT) Display() {
	for i := 0; i < crt.Width; i++ {
		if crt.Pixels.Contains(i) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")

	// Flush the current pixels
	crt.Pixels = dsa.NewSet[int]()
}

// Run a set of instructions and check signal strength at given checkpoints
func (crt *CRT) Run(instructions []Executable) {
	cpu := CPU{Register: 1, CycleCount: 1}
	pixel := 0

	for _, instruction := range instructions {
		if pixel >= cpu.Register-1 && pixel <= cpu.Register+1 {
			crt.Pixels.Add(pixel)
		}

		instruction.RunOn(&cpu)

		if pixel == crt.Width-1 {
			crt.Display()
			pixel = 0
		} else {
			pixel++
		}

		if cpu.CycleCount >= crt.Height*crt.Width {
			break
		}
	}
}

// Solve the second part by using the CRT display instructions
func solveSecondPart(path string) {
	input := readInput(path)

	crt := CRT{Height: 6, Width: 40, Pixels: dsa.NewSet[int]()}
	crt.Run(input)

	crt.Display()
}
