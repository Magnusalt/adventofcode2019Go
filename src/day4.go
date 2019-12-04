package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func checkNotDecreasing(pin string) bool {
	for i := 1; i < len(pin); i++ {
		if pin[i] < pin[i-1] {
			return false
		}
	}
	return true
}

func firstDay4() {
	count := 0
	for i := 387638; i <= 919123; i++ {
		s := strconv.Itoa(i)
		exp := regexp.MustCompile(`(00|11|22|33|44|55|66|77|88|99)`)
		if exp.MatchString(s) && checkNotDecreasing(s) {
			count++
		}
	}
	fmt.Println(count)
}
