package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"../helpers"
)

type direction struct {
	length int
	dir    string
}

type point struct {
	x int
	y int
}

func parseCable(line []string) []direction {
	var cable []direction
	for _, strDir := range line {
		l := strDir[1:]
		length, e := strconv.Atoi(l)
		helpers.Check(e)
		d := strDir[:1]
		cable = append(cable, direction{length, d})
	}

	return cable
}

func traceCable(cable []direction) (points []point) {

	currentPos := point{0, 0}
	for _, v := range cable {
		for i := 0; i < v.length; i++ {
			if v.dir == "R" {
				currentPos.x++
				points = append(points, point{currentPos.x, currentPos.y})
			}
			if v.dir == "L" {
				currentPos.x--
				points = append(points, point{currentPos.x, currentPos.y})
			}
			if v.dir == "U" {
				currentPos.y++
				points = append(points, point{currentPos.x, currentPos.y})
			}
			if v.dir == "D" {
				currentPos.y--
				points = append(points, point{currentPos.x, currentPos.y})
			}
		}
	}
	return points
}

// ByValue sorts a slice of int in ascending order
type ByValue []int

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i] < a[j] }

func firstDay3() {
	lines := helpers.GetFileAsStringArray("inputs/day3.txt")

	cable1 := parseCable(strings.Split(lines[0], ","))
	cable2 := parseCable(strings.Split(lines[1], ","))

	cableTrace1 := traceCable(cable1)
	cableTrace2 := traceCable(cable2)

	var intersectionDistance []int
	for _, v1 := range cableTrace1 {
		for _, v2 := range cableTrace2 {
			if v1.x == v2.x && v1.y == v2.y {
				distance := math.Abs(float64(v1.x)) + math.Abs(float64(v1.y))
				intersectionDistance = append(intersectionDistance, int(distance))
			}
		}
	}
	sort.Sort(ByValue(intersectionDistance))
	fmt.Println(intersectionDistance[0])
}

func secondDay3() {
	lines := helpers.GetFileAsStringArray("inputs/day3.txt")

	cable1 := parseCable(strings.Split(lines[0], ","))
	cable2 := parseCable(strings.Split(lines[1], ","))

	cableTrace1 := traceCable(cable1)
	cableTrace2 := traceCable(cable2)

	var intersections []point
	for _, v1 := range cableTrace1 {
		for _, v2 := range cableTrace2 {
			if v1.x == v2.x && v1.y == v2.y {
				intersections = append(intersections, point{v1.x, v1.y})
			}
		}
	}

	var intersectionStepDistance []int
	for _, v := range intersections {

		fmt.Println(v)
		stepsCable1 := 0
		for cableTrace1[stepsCable1].x != v.x || cableTrace1[stepsCable1].y != v.y {
			stepsCable1++
		}

		stepsCable2 := 0
		for cableTrace2[stepsCable2].x != v.x || cableTrace2[stepsCable2].y != v.y {
			stepsCable2++
		}

		intersectionStepDistance = append(intersectionStepDistance, stepsCable1+stepsCable2+2)
	}

	sort.Sort(ByValue(intersectionStepDistance))
	for _, v := range intersectionStepDistance {
		fmt.Println(v)
	}
}
