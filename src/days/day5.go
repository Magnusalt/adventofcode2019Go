package days

import (
	"../helpers"
)

// Day5 spins up an IntCode computer and runs the program
func Day5() {
	instructions := helpers.GetIntCodeInstructions("inputs/day5.txt")

	helpers.RunProgram(instructions, make([]int, 0), true, 0)
}
