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

func secondDay4() {
	var pincodes []string
	twoAdjacent := regexp.MustCompile(`(00|11|22|33|44|55|66|77|88|99)`)
	for i := 387638; i <= 919123; i++ {
		s := strconv.Itoa(i)
		if twoAdjacent.MatchString(s) && checkNotDecreasing(s) {
			pincodes = append(pincodes, s)
		}
	}
	count := 0
	twoOrMore := regexp.MustCompile(`4{2,5}|5{2,5}|6{2,5}|7{2,5}|8{2,5}|9{2,5}`)
	for _, p := range pincodes {
		matches := twoOrMore.FindAllString(p, 3)
		found := false
		for _, sub := range matches {
			if len(sub) == 2 && !found {
				count++
				found = true
			}
		}
	}
	fmt.Println(count)
}
