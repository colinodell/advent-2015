package main

import (
	"advent-2015/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("./day03/input.txt")

	fmt.Println("------- Part 1 -------")
	fmt.Printf("Santa delivers to %d houses\n\n", CountHouses(input))

	fmt.Println("------- Part 2 -------")
	fmt.Printf("Santa and Robo-Santa deliver to %d houses\n\n", CountHousesWithRoboSanta(input))
}

func CountHouses(input string) int {
	var position utils.Vector2
	houses := make(map[utils.Vector2]bool)

	houses[position] = true

	for _, char := range input {
		switch char {
		case '^':
			position = position.Add(utils.Vector2{0, 1})
		case '>':
			position = position.Add(utils.Vector2{1, 0})
		case 'v':
			position = position.Add(utils.Vector2{0, -1})
		case '<':
			position = position.Add(utils.Vector2{-1, 0})
		}
		houses[position] = true
	}

	return len(houses)
}

func CountHousesWithRoboSanta(input string) int {
	var santa utils.Vector2
	var roboSanta utils.Vector2

	houses := make(map[utils.Vector2]bool)

	houses[santa] = true

	for i, char := range input {
		var position utils.Vector2
		if i%2 == 0 {
			position = santa
		} else {
			position = roboSanta
		}

		switch char {
		case '^':
			position = position.Add(utils.Vector2{0, 1})
		case '>':
			position = position.Add(utils.Vector2{1, 0})
		case 'v':
			position = position.Add(utils.Vector2{0, -1})
		case '<':
			position = position.Add(utils.Vector2{-1, 0})
		}

		houses[position] = true

		if i%2 == 0 {
			santa = position
		} else {
			roboSanta = position
		}
	}

	return len(houses)
}
