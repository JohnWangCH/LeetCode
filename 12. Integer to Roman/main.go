package main

import (
	"strings"
)

func main() {}

func intToRoman(num int) string {
	romanList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	integerList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	for i := 0; i < len(romanList); i++ {
		repeatCount := num / integerList[i]
		if repeatCount != 0 {
			result += strings.Repeat(romanList[i], repeatCount)
		}

		num = num % integerList[i]
		if num == 0 {
			return result
		}
	}
	return result
}
