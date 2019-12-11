package days

import (
	"fmt"

	"../helpers"
)

// Day8a solves first puzzle of day 8
func Day8a() {
	imgData := helpers.GetFileAsStringArray("inputs/day8.txt")[0]
	var imgDataInt []int

	for _, p := range imgData {
		imgDataInt = append(imgDataInt, int(p)-48)
	}

	width := 25
	height := 6
	var images []map[int]int
	for i := 0; i < len(imgDataInt); i += width * height {
		image := make(map[int]int)
		for _, p := range imgDataInt[i : i+width*height] {
			image[p]++
		}
		images = append(images, image)
	}

	leastZeroesIndex := 0
	for i := 1; i < len(images); i++ {
		if images[leastZeroesIndex][0] > images[i][0] {
			leastZeroesIndex = i
		}
	}
	ans := images[leastZeroesIndex][1] * images[leastZeroesIndex][2]
	fmt.Println(ans)
}

// Day8b solves second puzzle of day 8
func Day8b() {
	imgData := helpers.GetFileAsStringArray("inputs/day8.txt")[0]
	var imgDataInt []int

	for _, p := range imgData {
		imgDataInt = append(imgDataInt, int(p)-48)
	}

	width := 25
	height := 6
	var images [][]int
	for i := 0; i < len(imgDataInt); i += width * height {
		image := make([]int, width*height)
		for i, p := range imgDataInt[i : i+width*height] {
			image[i] = p
		}
		images = append(images, image)
	}

	rows := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		for x := 0; x < width; x++ {
			for z := 0; z < len(images); z++ {
				v := images[z][y*width+x]
				if v == 2 {
					continue
				}
				if v == 1 || v == 0 {
					row[x] = v
					break
				}
			}
		}
		rows[y] = row
	}
	for _, x := range rows {
		for _, y := range x {
			if y == 0 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
