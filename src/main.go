package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileAsStringArray(path string) []string {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func getNextLine(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func main() {
	firstDay4()
}
