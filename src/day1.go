package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func firstDay1() {
	f, err := os.Open("day1.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, parseError := strconv.ParseInt(line, 10, 32)
		check(parseError)
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

func secondDay1() {
	f, err := os.Open("day1.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, parseError := strconv.Atoi(line)
		check(parseError)
		sum += recSum(0, parsedLine)
	}
	fmt.Println(sum)
}

func recSumConcurrent(currentSum int, mass int, out chan int) {
	sum := recSum(currentSum, mass)
	out <- sum
}

func day1bConcurrent() {
	masses := getFileAsStringArray("day1.txt")
	sum := 0
	sumChan := make(chan int)
	massesToCount := len(masses)
	for _, m := range masses {
		parsed, err := strconv.Atoi(m)
		check(err)
		go recSumConcurrent(0, parsed, sumChan)
	}

	for i := 0; i < massesToCount; i++ {
		sum += <-sumChan
	}

	fmt.Println(sum)
}
