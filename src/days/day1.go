package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../helpers"
)

// Day1a solves first puzzle of day 1
func Day1a() {
	f, err := os.Open("inputs/day1.txt")
	helpers.Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, parseError := strconv.ParseInt(line, 10, 32)
		helpers.Check(parseError)
		sum += (parsedLine / 3) - 2
	}
	fmt.Println(sum)
}

func recSum(currentSum int, moduleWeight int) int {
	nextWeight := ((moduleWeight / 3) - 2)

	if nextWeight <= 0 {
		return currentSum
	}

	return recSum(currentSum+nextWeight, nextWeight)
}

// Day1b solves second puzzle of day 1
func Day1b() {
	f, err := os.Open("inputs/day1.txt")
	helpers.Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, parseError := strconv.Atoi(line)
		helpers.Check(parseError)
		sum += recSum(0, parsedLine)
	}
	fmt.Println(sum)
}

func recSumConcurrent(currentSum int, mass int, out chan int) {
	sum := recSum(currentSum, mass)
	out <- sum
}

// Day1bConcurrent solves second puzzle of day 1 using goroutines
func Day1bConcurrent() {
	masses := helpers.GetFileAsStringArray("inputs/day1.txt")
	sum := 0
	sumChan := make(chan int)
	massesToCount := len(masses)
	for _, m := range masses {
		parsed, err := strconv.Atoi(m)
		helpers.Check(err)
		go recSumConcurrent(0, parsed, sumChan)
	}

	for i := 0; i < massesToCount; i++ {
		sum += <-sumChan
	}

	fmt.Println(sum)
}
