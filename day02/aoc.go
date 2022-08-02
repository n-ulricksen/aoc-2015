package day00

import (
	"math"
	"strconv"
	"strings"
)

type present struct {
	l int
	w int
	h int
}

// Find amount of wrapping paper needed for all presents. Each present requires
// its surface area in square feet, plus the area of the smallest side of the
// present for slack.
func part1(input string) int {
	presents := parseInput(input)

	totalSqft := 0
	for _, p := range presents {
		sqft := getSurfaceArea(p) + getSmallestSideArea(p)
		totalSqft += sqft
	}

	return totalSqft
}

// Find the amount of ribbon needed to wrap each present around its shortest
// perimeter, and tie a knot on top.
//   Wrap: smallest perimeter of any face
//   Bow:  cubic feet of volume of the present
func part2(input string) int {
	presents := parseInput(input)

	totalRibbon := 0
	for _, p := range presents {
		ribbon := getSmallestPerimeter(p) + getVolume(p)
		totalRibbon += ribbon
	}

	return totalRibbon
}

func parseInput(input string) []present {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	presents := make([]present, len(lines))

	for i, line := range lines {
		dimensions := strings.Split(line, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		p := present{l, w, h}
		presents[i] = p
	}
	return presents
}

// getSurfaceArea returns the calculated surface area of a given present
func getSurfaceArea(p present) int {
	l, w, h := p.l, p.w, p.h
	return 2*l*w + 2*w*h + 2*h*l
}

// getSmallestSideArea returns the area of the smallest side of the given
// present
func getSmallestSideArea(p present) int {
	area := math.MaxInt
	area = min(area, p.l*p.w)
	area = min(area, p.w*p.h)
	area = min(area, p.h*p.l)
	return area
}

// getSmallestPerimeter returns the length of the smallest perimeter of the
// given present
func getSmallestPerimeter(p present) int {
	perimeter := math.MaxInt
	perimeter = min(perimeter, 2*p.l+2*p.w)
	perimeter = min(perimeter, 2*p.w+2*p.h)
	perimeter = min(perimeter, 2*p.h+2*p.l)
	return perimeter
}

// getVolume returns the volume in cubic feet of the given present
func getVolume(p present) int {
	return p.l * p.w * p.h
}

// min returns the minimum of two given ints
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
