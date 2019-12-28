package main

import (
	"advent-2015/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadLines("./day05/input.txt")

	fmt.Println("------- Part 1 -------")
	fmt.Printf("There are %d nice strings in the input\n\n", CountNiceStrings(input, IsNiceV1))

	fmt.Println("------- Part 2 -------")
	fmt.Printf("There are %d nice strings in the input\n\n", CountNiceStrings(input, IsNiceV2))
}

func CountNiceStrings(input []string, test func(string) bool) int {
	var total int
	for _, str := range input {
		if test(str) {
			total++
		}
	}

	return total
}

func IsNiceV1(input string) bool {
	var vowelCount int
	var hasDoubleLetter bool

	for i, char := range input {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowelCount++
		}

		if i != 0 {
			pair := string(input[i-1]) + string(char)

			if pair == "ab" || pair == "cd" || pair == "pq" || pair == "xy" {
				return false
			}

			if pair[0] == pair[1] {
				hasDoubleLetter = true
			}
		}
	}

	return vowelCount >= 3 && hasDoubleLetter
}

func IsNiceV2(input string) bool {
	return hasRepeatingLetterWithGap(input) && hasTwoPairs(input)
}

func hasRepeatingLetterWithGap(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}

	return false
}

func hasTwoPairs(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if strings.Count(input, input[i:i+2]) >= 2 {
			return true
		}
	}

	return false
}
