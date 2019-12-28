package main

import (
	"advent-2015/utils"
	"fmt"
	"regexp"
)

func main() {
	input := utils.ReadLines("./day06/input.txt")
	g := NewGridFromInstructions(input)

	fmt.Println("------- Part 1 -------")
	fmt.Printf("There are %d lights on\n\n", g.CountLit())

	fmt.Println("------- Part 2 -------")
	fmt.Printf("The total brightness is %d\n\n", g.TotalBrightness())
}

type Grid struct {
	onOrOff    map[utils.Vector2]bool
	brightness map[utils.Vector2]int
}

func NewGrid() *Grid {
	g := new(Grid)
	g.onOrOff = make(map[utils.Vector2]bool)
	g.brightness = make(map[utils.Vector2]int)
	return g
}

func NewGridFromInstructions(input []string) *Grid {
	grid := NewGrid()

	regex := regexp.MustCompile(`(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)`)

	for _, str := range input {
		matches := regex.FindStringSubmatch(str)

		min := utils.Vector2{utils.ToInt(matches[2]), utils.ToInt(matches[3])}
		max := utils.Vector2{utils.ToInt(matches[4]), utils.ToInt(matches[5])}

		switch matches[1] {
		case "turn on":
			grid.TurnOn(min, max)
		case "toggle":
			grid.Toggle(min, max)
		case "turn off":
			grid.TurnOff(min, max)
		}
	}

	return grid
}

func (g Grid) TurnOn(min, max utils.Vector2) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			pos := utils.Vector2{x, y}
			g.onOrOff[pos] = true
			g.brightness[pos]++
		}
	}
}

func (g Grid) TurnOff(min, max utils.Vector2) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			pos := utils.Vector2{x, y}
			g.onOrOff[pos] = false
			g.brightness[pos] = utils.Max(0, g.brightness[pos]-1)
		}
	}
}

func (g Grid) Toggle(min, max utils.Vector2) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			pos := utils.Vector2{x, y}
			g.onOrOff[pos] = !g.onOrOff[pos]
			g.brightness[pos] += 2
		}
	}
}

func (g Grid) CountLit() int {
	var count int

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			pos := utils.Vector2{x, y}
			if g.onOrOff[pos] {
				count++
			}
		}
	}

	return count
}

func (g Grid) TotalBrightness() int {
	var brightness int

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			pos := utils.Vector2{x, y}
			brightness += g.brightness[pos]
		}
	}

	return brightness
}
