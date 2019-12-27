package main

import (
	"advent-2015/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := utils.ReadLines("./day02/input.txt")

	var wrappingPaperNeeded int
	var ribbonNeeded int
	for _, p := range parsePresents(input) {
		wrappingPaperNeeded += p.CalculateWrappingPaper()
		ribbonNeeded += p.CalculateRibbon()
	}

	fmt.Println("------- Part 1 -------")
	fmt.Printf("The total wrapping paper needed is %d square feet\n\n", wrappingPaperNeeded)

	fmt.Println("------- Part 2 -------")
	fmt.Printf("The total ribbon needed is %d feet\n\n", ribbonNeeded)
}

func parsePresents(input []string) []Present {
	var presents []Present

	regex := regexp.MustCompile(`^(\d+)x(\d+)x(\d+)$`)

	for _, line := range input {
		matches := regex.FindStringSubmatch(line)

		l, _ := strconv.Atoi(matches[1])
		w, _ := strconv.Atoi(matches[2])
		h, _ := strconv.Atoi(matches[3])

		presents = append(presents, Present{l, w, h})
	}

	return presents
}

type Present struct {
	length, width, height int
}

func (p Present) CalculateWrappingPaper() int {
	a := p.length * p.width
	b := p.width * p.height
	c := p.height * p.length

	slack := utils.Min(a, b, c)

	return 2*a + 2*b + 2*c + slack
}

func (p Present) CalculateRibbon() int {
	smallestPerimeter := 2 * (p.length + p.width + p.height - utils.Max(p.length, p.width, p.height))

	return smallestPerimeter + p.length*p.width*p.height
}
