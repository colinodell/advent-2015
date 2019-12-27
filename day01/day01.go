package main

import (
	"advent-2015/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("./day01/input.txt")

	fmt.Println("------- Part 1 -------")
	fmt.Printf("The directions take Santa to floor %d\n\n", Travel(input))

	fmt.Println("------- Part 2 -------")
	fmt.Printf("Santa reached the basement at character %d\n\n", TravelUntilBasementReached(input))
}

func Travel(directions string) int {
	var level int
	for _, x := range directions {
		if x == '(' {
			level++
		} else if x == ')' {
			level--
		}
	}

	return level
}

func TravelUntilBasementReached(directions string) int {
	var level int
	for i, x := range directions {
		if x == '(' {
			level++
		} else if x == ')' {
			level--
			if level == -1 {
				return i + 1
			}
		}
	}

	panic("Santa never reached the basement!")
}
