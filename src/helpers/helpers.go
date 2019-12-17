package helpers

import (
	"bufio"
	"os"
	"strings"
)

// Check if error is nil app will panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetFileAsStringArray will open a file and return the content as a string array
func GetFileAsStringArray(path string) []string {
	f, err := os.Open(path)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

// GetIntCodeInstructions turns the puzzle input into an IntCode compter program
func GetIntCodeInstructions(path string) []string {
	return strings.Split(GetFileAsStringArray(path)[0], ",")
}
