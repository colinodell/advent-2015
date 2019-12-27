package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	check(err)

	return strings.TrimSpace(string(buf))
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
