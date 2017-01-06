package convertRoman

import "fmt"

// ConvertRoman converts a Roman Numeral into an integer and
// prints to the screen. There is no return value
func ConvertRoman(input string) {
	current, previous, total := 0, 0, 0
	for index, value := range input {
		current = convertChar(value)
		total += current
		if index > 0 && current > previous {
			total -= (previous * 2)
		}
		previous = current
	}
	fmt.Printf("%s is %d\n", input, total)
}

func convertChar(val rune) int {
	switch string(val) {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return -1
}
