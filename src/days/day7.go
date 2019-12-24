package days

import (
	"fmt"

	"../helpers"
)

func generatePermutations(k int, a []int, out *[][]int) {
	if k == 1 {
		return
	}

	generatePermutations(k-1, a, out)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			ai := a[i]
			a[i] = a[k-1]
			a[k-1] = ai
		} else {
			a0 := a[0]
			a[0] = a[k-1]
			a[k-1] = a0
		}
		res := make([]int, 5)
		copy(res, a)
		*out = append(*out, res)
		generatePermutations(k-1, a, out)
	}
}

// Day7a solves first puzzle day7
func Day7a() {
	a := []int{0, 1, 2, 3, 4}
	var perm [][]int
	generatePermutations(len(a), a, &perm)

	instructions := helpers.GetIntCodeInstructions("inputs/day7.txt")
	max := 0
	for _, combination := range perm {
		fmt.Println(combination)
		inputs := []int{0, 0}
		output := []int{0}
		for i := 0; i < 5; i++ {
			inputs[0] = combination[i]
			inputs[1] = output[0]
			output, _, _ = helpers.RunProgram(instructions, inputs, false, 0)
			fmt.Println(output)
		}
		fmt.Println("====")
		if output[0] > max {
			max = output[0]
		}
	}
	fmt.Println(max)
}

// Day7b solves second part of day 7
func Day7b() {
	a := []int{5, 6, 7, 8, 9}
	var perm [][]int
	generatePermutations(len(a), a, &perm)

	max := 0

	for _, combination := range perm {

		fmt.Println(combination)

		inputs := []int{0, 0}
		outputs := make([][]int, 5)
		instructions := make([][]string, 5)
		pcs := []int{0, 0, 0, 0, 0}
		reachedEnd := false

		for i := 0; i < 5; i++ {
			inputs[0] = combination[i]
			if i == 0 {
				inputs[1] = 0
			} else {
				inputs[1] = outputs[i-1][0]
			}
			instructions[i] = helpers.GetIntCodeInstructions("inputs/day7.txt")
			outputs[i], reachedEnd, pcs[i] = helpers.RunProgram(instructions[i], inputs, false, pcs[i])
		}
		for !reachedEnd {
			for i := 0; i < 5; i++ {
				if i == 0 {
					inputs = outputs[4]
				} else {
					inputs = outputs[i-1]
				}
				halted := false
				outputs[i], halted, pcs[i] = helpers.RunProgram(instructions[i], inputs, false, pcs[i])
				reachedEnd = halted && i == 4
			}
		}
		if outputs[4][0] > max {
			max = outputs[4][0]
		}
	}
	fmt.Println(max)
}
