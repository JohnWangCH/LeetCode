package main

func main() {}

func reverse(x int) int {
	maxIntValue := 2147483647
	minIntValue := -2147483648

	result := 0
	for x != 0 {
		isOverResultBoundary := (result > (maxIntValue / 10)) || (result < (minIntValue / 10)) ||
			((result == (maxIntValue / 10)) && (x%10 > 7)) || ((result == (minIntValue / 10)) && (x%10 < -8))

		if isOverResultBoundary {
			return 0
		}
		result = result*10 + x%10
		x = x / 10
	}
	return result
}
