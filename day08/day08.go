package main

import (
	"advent-2015/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ReadLines("./day08/input.txt")

	fmt.Println("------- Part 1 -------")
	fmt.Printf("The difference after decoding is %d\n\n", Decode(input...))

	fmt.Println("------- Part 2 -------")
	fmt.Printf("The difference after encoding is %d\n\n", Encode(input...))
}

func Decode(input ...string) int {
	literals, values := 0, 0

	for _, line := range input {
		literals += len(line)

		parsed, _ := strconv.Unquote(line)

		values += len(parsed)
	}

	return literals - values
}

func Encode(input ...string) int {
	literals, values := 0, 0

	for _, line := range input {
		literals += len(line)

		parsed := strconv.Quote(line)

		values += len(parsed)
	}

	return values - literals
}
