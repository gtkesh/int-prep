package main

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}
