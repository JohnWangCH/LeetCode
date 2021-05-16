package main

import (
	"math"
	"strings"
	"unicode"
)

func main() {}

func myAtoi(s string) int {
	trimmedString := strings.TrimSpace(s)
	if len(trimmedString) == 0 {
		return 0
	}

	sign := 1
	number := 0

	if string(trimmedString[0]) == "+" || string(trimmedString[0]) == "-" {
		if string(trimmedString[0]) == "-" {
			sign = -1
		}
		trimmedString = trimmedString[1:]
	}

	for _, char := range trimmedString {
		if !unicode.IsDigit(char) {
			break
		}

		if (sign * number) >= math.MaxInt32 {
			return math.MaxInt32
		} else if (sign * number) <= math.MinInt32 {
			return math.MinInt32
		}

		number = number*10 + (int(char) - int('0'))
	}

	resultValue := sign * number
	if resultValue >= math.MaxInt32 {
		return math.MaxInt32
	} else if resultValue <= math.MinInt32 {
		return math.MinInt32
	}
	return resultValue
}
