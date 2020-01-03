package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	seed := "1113122113"

	fmt.Println("------- Part 1 -------")
	after40 := LookAndSayMultipleTimes(seed, 40)
	fmt.Printf("After 40 rounds, the result has length %d\n\n", len(after40))

	fmt.Println("------- Part 2 -------")
	after50 := LookAndSayMultipleTimes(after40, 10)
	fmt.Printf("After 50 rounds, the result has length %d\n\n", len(after50))
}

func LookAndSay(input string) string {
	lastChar := '0'
	lastCount := 0

	var result strings.Builder

	for i := 0; i < len(input); i++ {
		if rune(input[i]) == lastChar {
			lastCount++
		} else {
			if lastChar != '0' {
				result.WriteString(strconv.Itoa(lastCount))
				result.WriteRune(lastChar)
			}
			lastChar = rune(input[i])
			lastCount = 1
		}
	}

	result.WriteString(strconv.Itoa(lastCount))
	result.WriteRune(lastChar)

	return result.String()
}

func LookAndSayMultipleTimes(str string, rounds int) string {
	for i := 0; i < rounds; i++ {
		str = LookAndSay(str)
	}

	return str
}
