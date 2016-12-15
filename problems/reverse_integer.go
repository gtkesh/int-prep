package main

import "strconv"

func reverse(x int) int {

	var num int
	isNeg := x < 0

	s := strconv.Itoa(x)
	digits := len(s)

	start := 0
	if isNeg {
		start = 1
	}

	divider := 1
	for i := start; i < digits-1; i++ {
		divider *= 10
	}

	multiplier := 1
	for i := start; i < digits; i++ {
		k := x / divider
		num += k * multiplier
		x = x % divider
		divider /= 10
		multiplier *= 10
	}

	if num > 2147483647 || num < -2147483648 {
		return 0
	}

	return num
}
