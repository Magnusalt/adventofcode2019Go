package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func first() {
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

func recSum(currentSum int64, moduleWeight int64) int64 {
	nextWeight := ((moduleWeight / 3) - 2)

	if nextWeight <= 0 {
		return currentSum
	}

	return recSum(currentSum+nextWeight, nextWeight)
}

func second() {
	f, err := os.Open("day1.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, parseError := strconv.ParseInt(line, 10, 32)
		check(parseError)
		sum += recSum(0, parsedLine)
	}
	fmt.Println(sum)
}

func main() {
	//first()
	second()
}
