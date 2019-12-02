package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func firstDay2() {
	f, err := os.Open("day2.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	instructions := strings.Split(line, ",")

	for k, v := range instructions {
		fmt.Println(k)
		fmt.Println(v)
	}

}
