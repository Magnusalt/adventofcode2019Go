package days

import (
	"fmt"

	"../helpers"
)

func generatePermutations(k int, a []int, out chan []int) {
	if k == 1 {
		return
	}

	generatePermutations(k-1, a, out)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			ai := a[i]
			a[i] = a[k-1]
			a[k-1] = ai
			out <- a
		} else {
			a0 := a[0]
			a[0] = a[k-1]
			a[k-1] = a0
			out <- a
		}
		generatePermutations(k-1, a, out)
	}
}

// Day7a solves first puzzle day7
func Day7a() {
	a := []int{0, 1, 2, 3, 4}
	nbrOfPosibilities := 120 // 5!
	reporter := make(chan []int)
	go generatePermutations(len(a), a, reporter)

	instructions := helpers.GetIntCodeInstructions("inputs/day7.txt")
	max := 0
	for i := 0; i < nbrOfPosibilities-1; i++ {
		combination := <-reporter
		fmt.Println(combination)
		inputs := []int{0, 0}
		output := []int{0}
		for i := 0; i < 5; i++ {
			inputs[0] = combination[i]
			inputs[1] = output[0]
			output = helpers.RunProgram(instructions, inputs, false)
			fmt.Println(output)
		}
		fmt.Println("====")
		if output[0] > max {
			max = output[0]
		}
	}
	fmt.Println(max)
}
