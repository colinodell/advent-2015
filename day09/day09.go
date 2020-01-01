package main

import (
	"advent-2015/utils"
	"fmt"
	"github.com/dbyio/heappermutations"
	"math"
	"regexp"
	"strings"
)

func main() {
	input := utils.ReadFile("./day09/input.txt")

	shortest, longest := FindShortestAndLongestRoute(input)

	fmt.Println("------- Part 1 -------")
	fmt.Printf("The shortest distance is %d\n\n", shortest)

	fmt.Println("------- Part 2 -------")
	fmt.Printf("The longest distance is %d\n\n", longest)
}

type route struct {
	from string
	to   string
}

func FindShortestAndLongestRoute(input string) (int, int) {
	routes, cities := parseInput(input)

	shortestDistance := math.MaxInt64
	longestDistance := math.MinInt64

	for _, permutation := range heappermutations.Strings(cities) {
		distance := 0

		for i := 0; i < len(permutation)-1; i++ {
			distance += routes[route{permutation[i], permutation[i+1]}]
		}

		if distance < shortestDistance {
			shortestDistance = distance
		}

		if distance > longestDistance {
			longestDistance = distance
		}
	}

	return shortestDistance, longestDistance
}

func parseInput(input string) (routes map[route]int, cities []string) {
	routes = make(map[route]int)
	foundCities := make(map[string]bool)

	for _, line := range strings.Split(input, "\n") {
		from, to, distance := parseLine(line)

		routes[route{from, to}] = distance
		routes[route{to, from}] = distance

		if !foundCities[from] {
			cities = append(cities, from)
			foundCities[from] = true
		}

		if !foundCities[to] {
			cities = append(cities, to)
			foundCities[to] = true
		}
	}

	return
}

var regex = regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)

func parseLine(line string) (string, string, int) {
	matches := regex.FindStringSubmatch(line)
	return matches[1], matches[2], utils.ToInt(matches[3])
}
