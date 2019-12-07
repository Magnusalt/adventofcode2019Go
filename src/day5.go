package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day5a() {
	instructions := strings.Split(getFileAsStringArray("day5.txt")[0], ",")

	instructionPointer := 0
	reader := bufio.NewReader(os.Stdin)

	opCode := getOpcode(instructions[instructionPointer])

	for opCode != "99" {
		jumpDelta := 0
		fmt.Println(opCode)
		currentInstruction := opCode[4]
		switch currentInstruction {
		case '1':
			v1 := readValue(instructions, instructionPointer+1, opCode[2])
			v2 := readValue(instructions, instructionPointer+2, opCode[1])
			addRes := v1 + v2
			instructions = writeValue(addRes, instructions, instructionPointer+3)
			jumpDelta += 4
		case '2':
			v1 := readValue(instructions, instructionPointer+1, opCode[2])
			v2 := readValue(instructions, instructionPointer+2, opCode[1])
			multRes := v1 * v2
			instructions = writeValue(multRes, instructions, instructionPointer+3)
			jumpDelta += 4
		case '3':
			input, err := reader.ReadString('\r')
			check(err)
			intInput := parseInt(input[:1])
			instructions = writeValue(intInput, instructions, instructionPointer+1)
			jumpDelta += 2
		case '4':
			vToOutput := readValue(instructions, instructionPointer+1, opCode[2])
			fmt.Println(vToOutput)
			jumpDelta += 2
		case '5':
			v1 := readValue(instructions, instructionPointer+1, opCode[2])
			v2 := readValue(instructions, instructionPointer+2, opCode[1])
			if v1 != 0 {
				instructionPointer = v2
			} else {
				jumpDelta += 3
			}
		case '6':
			v1 := readValue(instructions, instructionPointer+1, opCode[2])
			v2 := readValue(instructions, instructionPointer+2, opCode[1])
			if v1 == 0 {
				instructionPointer = v2
			} else {
				jumpDelta += 3
			}
		case '7':
			v1 := readValue(instructions, instructionPointer+1, opCode[2])
			v2 := readValue(instructions, instructionPointer+2, opCode[1])
			if v1 < v2 {
				instructions = writeValue(1, instructions, instructionPointer+3)
			} else {
				instructions = writeValue(0, instructions, instructionPointer+3)
			}
			jumpDelta += 4
		case '8':
			v1 := readValue(instructions, instructionPointer+1, opCode[2])
			v2 := readValue(instructions, instructionPointer+2, opCode[1])
			if v1 == v2 {
				instructions = writeValue(1, instructions, instructionPointer+3)
			} else {
				instructions = writeValue(0, instructions, instructionPointer+3)
			}
			jumpDelta += 4
		}

		instructionPointer += jumpDelta

		opCode = getOpcode(instructions[instructionPointer])
	}
}

func readValue(instructions []string, pos int, mode byte) int {
	if mode == '1' {
		vImmidiate := parseInt(instructions[pos])
		return vImmidiate
	}
	address := parseInt(instructions[pos])
	vPosition := parseInt(instructions[address])
	return vPosition
}

func parseInt(n string) int {
	v, er := strconv.Atoi(n)
	check(er)
	return v
}

func writeValue(v int, instructions []string, pos int) (updatedInstructions []string) {
	addressToWriteTo := parseInt(instructions[pos])
	instructions[addressToWriteTo] = strconv.Itoa(v)
	updatedInstructions = instructions
	return updatedInstructions
}

func getOpcode(ins string) string {
	if ins == "99" {
		return ins
	}
	paddingLength := 5 - len(ins)
	opCode := strings.Repeat("0", paddingLength) + ins
	return opCode
}
