package main

import (
	"fmt"
)

func main() {
	password := "hepxcrrq"

	fmt.Println("------- Part 1 -------")
	password = GetNextValidPassword(password)
	fmt.Printf("The next password is %s\n\n", password)

	fmt.Println("------- Part 2 -------")
	password = GetNextValidPassword(password)
	fmt.Printf("The next password is %s\n\n", password)
}

func GetNextValidPassword(str string) string {
	for {
		str = nextPassword(str)
		if hasStraight(str) && hasOnlyValidCharacters(str) && hasTwoNonOverlappingPairs(str) {
			return str
		}
	}
}

func nextPassword(previous string) string {
	p := []byte(previous)

	for i := len(p) - 1; i >= 0; i-- {
		p[i]++
		if p[i] > 'z' {
			p[i] = 'a'
		} else {
			break
		}
	}

	return string(p)
}

func hasStraight(str string) bool {
	for i := 0; i < len(str)-3; i++ {
		if str[i]+1 == str[i+1] && str[i+1]+1 == str[i+2] {
			return true
		}
	}

	return false
}

func hasOnlyValidCharacters(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == 'i' || str[i] == 'o' || str[i] == 'l' {
			return false
		}
	}

	return true
}

func hasTwoNonOverlappingPairs(str string) bool {
	seen := make(map[uint8]bool)

	count := 0
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] && !seen[str[i]] {
			count++
			seen[str[i]] = true
		}
	}

	return count > 1
}
