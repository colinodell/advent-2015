package utils

import "math"

func Min(numbers ...int) int {
	min := math.MaxInt64

	for _, x := range numbers {
		if x < min {
			min = x
		}
	}

	return min
}

func Max(numbers ...int) int {
	max := math.MinInt64

	for _, x := range numbers {
		if x > max {
			max = x
		}
	}

	return max
}
