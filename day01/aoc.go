package day00

// Given instructions:
//   '(' --> up a floor
//   ')' --> down a floor
//
// Which floor will santa end up on?
func part1(input string) int {
	floor := 0
	for _, ch := range input {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
	}
	return floor
}

// Same instructions... which step is the first to take Santa underground?
func part2(input string) int {
	floor := 0
	for i, ch := range input {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}
