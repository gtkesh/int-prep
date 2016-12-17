package main

import "strings"

/*
Given an input string, reverse the string word by word.

For example:

Given s = "the sky is blue",
return "blue is sky the

*/

func reverseWords(s string) string {
	var reversed string
	words := strings.Split(reverseStr(s), " ")

	for _, w := range words {
		reversed += reverseStr(w) + " "
	}
	return reversed
}

func reverseStr(s string) string {
	reversed := []rune{}
	i := len(s)
	for _, c := range s {
		i--
		reversed[i] = c
	}
	return string(reversed)
}
