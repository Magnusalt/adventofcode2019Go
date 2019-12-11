package helpers

import (
	"bufio"
	"os"
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
