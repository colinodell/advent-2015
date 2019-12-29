package main

import (
	"advent-2015/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := utils.ReadLines("./day07/input.txt")
	c := NewCircuit(input)

	fmt.Println("------- Part 1 -------")
	a := c.GetValue("a")
	fmt.Printf("After running the circuit, wire a has signal value %d\n\n", a)

	fmt.Println("------- Part 2 -------")
	c.Reset()
	c.Override("b", a)
	a = c.GetValue("a")
	fmt.Printf("After overriding 'b' and re-running the circuit, wire a has signal value %d\n\n", a)
}

type operation int

const (
	set operation = iota
	and
	or
	lshift
	rshift
	not
)

type Circuit map[string]*Wire

type Wire struct {
	Inputs    []string
	Operation operation
	resolved  bool
	output    uint16
}

func (c Circuit) GetValue(wire string) uint16 {
	// wire might actually be a number instead
	if n, err := strconv.Atoi(wire); err == nil {
		return uint16(n)
	}

	w := c[wire]
	if !w.resolved {
		switch w.Operation {
		case set:
			w.output = c.GetValue(w.Inputs[0])
		case and:
			w.output = c.GetValue(w.Inputs[0]) & c.GetValue(w.Inputs[1])
		case or:
			w.output = c.GetValue(w.Inputs[0]) | c.GetValue(w.Inputs[1])
		case lshift:
			w.output = c.GetValue(w.Inputs[0]) << c.GetValue(w.Inputs[1])
		case rshift:
			w.output = c.GetValue(w.Inputs[0]) >> c.GetValue(w.Inputs[1])
		case not:
			w.output = ^c.GetValue(w.Inputs[0])
		}

		w.resolved = true
	}

	return w.output
}

func (c Circuit) Reset() {
	for _, wire := range c {
		wire.output = 0
		wire.resolved = false
	}
}

func (c Circuit) Override(wire string, value uint16) {
	c[wire].resolved = true
	c[wire].output = value
}

func NewCircuit(instructions []string) Circuit {
	c := make(Circuit)

	reSet := regexp.MustCompile(`^(\w+) -> (\w+)$`)
	reAnd := regexp.MustCompile(`^(\w+) AND (\w+) -> (\w+)$`)
	reOr := regexp.MustCompile(`^(\w+) OR (\w+) -> (\w+)$`)
	reLshift := regexp.MustCompile(`^(\w+) LSHIFT (\d+) -> (\w+)$`)
	reRshift := regexp.MustCompile(`^(\w+) RSHIFT (\d+) -> (\w+)$`)
	reNot := regexp.MustCompile(`^NOT (\w+) -> (\w+)$`)

	for _, line := range instructions {
		if matches := reSet.FindStringSubmatch(line); len(matches) > 0 {
			c[matches[2]] = &Wire{Operation: set, Inputs: matches[1:2]}
		} else if matches := reAnd.FindStringSubmatch(line); len(matches) > 0 {
			c[matches[3]] = &Wire{Operation: and, Inputs: matches[1:3]}
		} else if matches := reOr.FindStringSubmatch(line); len(matches) > 0 {
			c[matches[3]] = &Wire{Operation: or, Inputs: matches[1:3]}
		} else if matches := reLshift.FindStringSubmatch(line); len(matches) > 0 {
			c[matches[3]] = &Wire{Operation: lshift, Inputs: matches[1:3]}
		} else if matches := reRshift.FindStringSubmatch(line); len(matches) > 0 {
			c[matches[3]] = &Wire{Operation: rshift, Inputs: matches[1:3]}
		} else if matches := reNot.FindStringSubmatch(line); len(matches) > 0 {
			c[matches[2]] = &Wire{Operation: not, Inputs: matches[1:2]}
		} else {
			panic(fmt.Sprintf("Nothing matched '%s'", line))
		}
	}

	return c
}
