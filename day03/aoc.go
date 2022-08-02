package day00

import (
	"fmt"
	"strings"
)

type position struct {
	x int
	y int
}

// Given a list of instructions (left, right, up, down), return the number of
// unique locations visited.
func part1(input string) int {
	directions := strings.Split(input, "")

	pos := position{0, 0} // Starting position
	houses := 1           // Starting position is visited and counted
	visited := map[string]bool{pos.String(): true}

	for _, dir := range directions {
		dx, dy := 0, 0
		switch dir {
		case ">":
			dx = 1
		case "v":
			dy = 1
		case "<":
			dx = -1
		case "^":
			dy = -1
		}
		pos.x += dx
		pos.y += dy

		posStr := pos.String()
		if !visited[posStr] {
			houses++
		}
		visited[posStr] = true
	}

	return houses
}

// Same as part1, but now there are two Santas. If each Santa takes turns
// reading from the list of instructions, how many houses will be visited?
func part2(input string) int {
	directions := strings.Split(input, "")

	posSanta := position{0, 0} // Santa starting position
	posRobo := position{0, 0}  // Robo-Santa starting position
	houses := 1                // Starting position is visited and counted
	visited := map[string]bool{posSanta.String(): true}

	for i, dir := range directions {
		dx, dy := 0, 0
		switch dir {
		case ">":
			dx = 1
		case "v":
			dy = 1
		case "<":
			dx = -1
		case "^":
			dy = -1
		}

		var posStr string
		if i%2 == 0 {
			posSanta.x += dx
			posSanta.y += dy
			posStr = posSanta.String()
		} else {
			posRobo.x += dx
			posRobo.y += dy
			posStr = posRobo.String()
		}

		if !visited[posStr] {
			houses++
		}
		visited[posStr] = true
	}

	return houses
}

func (p position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
