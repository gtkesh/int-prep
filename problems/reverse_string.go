package main

func reverseString(s string) string {
	var reversed []byte

	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}

	return string(reversed)
}
