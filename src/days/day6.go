package days

import (
	"fmt"
	"strings"

	"../helpers"
)

type orbiter struct {
	name  string
	toCOM *orbiter
}

func (o orbiter) hasName(name string) bool {
	return o.name == name
}

func orbitsContains(orbiters []orbiter, name string) (bool, *orbiter) {
	for _, o := range orbiters {
		if o.hasName(name) {
			return true, &o
		}
	}
	return false, nil
}

// Day6a solves day 6 :)
func Day6a() {
	var uniqueOrbiters []orbiter
	var orbits = helpers.GetFileAsStringArray("day6test.txt")

	for _, o := range orbits {
		rel := strings.Split(o, ")")
		containsParent, _ := orbitsContains(uniqueOrbiters, rel[0])
		if !containsParent {
			uniqueOrbiters = append(uniqueOrbiters, orbiter{name: rel[0], toCOM: nil})
		} else {

		}
		containsChild, child := orbitsContains(uniqueOrbiters, rel[1])
		if !containsChild {
			_, parent := orbitsContains(uniqueOrbiters, rel[0])
			uniqueOrbiters = append(uniqueOrbiters, orbiter{name: rel[1], toCOM: parent})
		} else {
			if child.toCOM == nil {
				p := orbiter{name: rel[0], toCOM: nil}
				uniqueOrbiters = append(uniqueOrbiters, p)
				child.toCOM = p
			}
		}
	}

	for _, uo := range uniqueOrbiters {
		if uo.toCOM != nil {
			fmt.Println(uo.name, (*uo.toCOM).name)
		} else {
			fmt.Println(uo.name)
		}

	}
}
