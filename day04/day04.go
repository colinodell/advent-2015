package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("------- Part 1 -------")
	fmt.Printf("Lowest number with 5 zeros is %d\n\n", MineAdventCoin5("bgvyzdsv"))

	fmt.Println("------- Part 2 -------")
	fmt.Printf("Lowest number with 6 zeros is %d\n\n", MineAdventCoin6("bgvyzdsv"))
}

func MineAdventCoin5(secretKey string) int {
	for number := 0; ; number++ {
		input := secretKey + strconv.Itoa(number)
		sum := md5.Sum([]byte(input))

		if sum[0] == 0 && sum[1] == 0 && sum[2] < 16 {
			return number
		}
	}
}

func MineAdventCoin6(secretKey string) int {
	for number := 0; ; number++ {
		input := secretKey + strconv.Itoa(number)
		sum := md5.Sum([]byte(input))

		if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			return number
		}
	}
}
