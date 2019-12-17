package days

import (
	"fmt"

	"../helpers"
)

// Day7a solves first puzzle day7
func Day7a() {
	instructions := helpers.GetIntCodeInstructions("inputs/day7test.txt")

	inputs := []int{0, 0}
	res := 0
	for i := 0; i < 5; i++ {
		output := helpers.RunProgram(instructions, inputs, false)
		inputs[0]++
		inputs[1] = output[0]
		res = output[0]
	}

	fmt.Println(res)
}
