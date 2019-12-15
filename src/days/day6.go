package days

import (
	"fmt"
	"strings"

	"../helpers"
)

// Day6a solves day 6 :)
func Day6a() {
	childOrbiters := make(map[string]string)
	var orbits = helpers.GetFileAsStringArray("inputs/day6.txt")
	for _, o := range orbits {
		relation := strings.Split(o, ")")
		childOrbiters[relation[1]] = relation[0]
	}
	total := 0
	for k := range childOrbiters {
		childCount := 0
		current := k
		for current != "COM" {
			current = childOrbiters[current]
			childCount++
		}
		total += childCount
	}
	fmt.Println(total)
}

// Day6b solves day 6 b :)
func Day6b() {
	childOrbiters := make(map[string]string)
	var orbits = helpers.GetFileAsStringArray("inputs/day6test.txt")
	for _, o := range orbits {
		relation := strings.Split(o, ")")
		childOrbiters[relation[1]] = relation[0]
	}
	youPath := traceOrbit("YOU", childOrbiters)
	sanPath := traceOrbit("SAN", childOrbiters)
	youPathIndex := len(youPath) - 1
	sanPathIndex := len(sanPath) - 1
	for youPath[youPathIndex] == sanPath[sanPathIndex] {
		sanPathIndex--
		youPathIndex--
	}
	// the path indices are reduced one extra time to account for the SAN and YOU orbits
	youPathDistanceToFork := len(youPath[:youPathIndex])
	sanPathDistanceToFork := len(sanPath[:sanPathIndex])
	fmt.Println(youPathDistanceToFork + sanPathDistanceToFork)
}

func traceOrbit(start string, orbitMap map[string]string) []string {
	var path []string
	current := start
	for current != "COM" {
		path = append(path, current)
		current = orbitMap[current]
	}
	return path
}
