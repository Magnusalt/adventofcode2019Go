package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../helpers"
)

func firstDay2() {
	f, err := os.Open("inputs/day2.txt")
	helpers.Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	stringInstructions := strings.Split(line, ",")
	intInstructions := make([]int, len(stringInstructions))

	for k, v := range stringInstructions {
		parsed, err := strconv.Atoi(v)
		helpers.Check(err)
		intInstructions[k] = parsed
	}

	intInstructions[1] = 12
	intInstructions[2] = 2

	currentPosition := 0
	currentInstruction := intInstructions[currentPosition]

	for currentInstruction != 99 {
		a := intInstructions[currentPosition+1]
		b := intInstructions[currentPosition+2]
		targetIndex := intInstructions[currentPosition+3]
		if currentInstruction == 1 {
			intInstructions[targetIndex] = intInstructions[a] + intInstructions[b]
		}
		if currentInstruction == 2 {
			intInstructions[targetIndex] = intInstructions[a] * intInstructions[b]
		}
		currentPosition += 4
		currentInstruction = intInstructions[currentPosition]
	}

	fmt.Println(intInstructions[0])

}

func attempt(program []int, noun int, verb int) int {
	program[1] = noun
	program[2] = verb

	currentPosition := 0
	currentInstruction := program[currentPosition]

	for currentInstruction != 99 {
		a := program[currentPosition+1]
		b := program[currentPosition+2]
		targetIndex := program[currentPosition+3]
		if currentInstruction == 1 {
			program[targetIndex] = program[a] + program[b]
		}
		if currentInstruction == 2 {
			program[targetIndex] = program[a] * program[b]
		}
		currentPosition += 4
		currentInstruction = program[currentPosition]
	}

	return program[0]
}

func getOriginalProgram(input []string) []int {
	intInstructions := make([]int, len(input))

	for k, v := range input {
		parsed, err := strconv.Atoi(v)
		helpers.Check(err)
		intInstructions[k] = parsed
	}
	return intInstructions
}

func secondDay2() {
	f, err := os.Open("inputs/day2.txt")
	helpers.Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	stringInstructions := strings.Split(line, ",")

	targetOutput := 19690720

	for noun := range make([]int, 100) {
		for verb := range make([]int, 100) {
			program := getOriginalProgram(stringInstructions)
			attemptResult := attempt(program, noun, verb)
			if attemptResult == targetOutput {
				fmt.Println("noun: ", noun, "verb: ", verb)
				fmt.Println("result: ", 100*noun+verb)
			}
		}
	}
}
