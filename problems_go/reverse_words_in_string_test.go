package main

import "testing"

func TestReverseWords(t *testing.T) {

	if r := reverseWords("the sky is blue"); r != "blue is sky the" {
		t.Error(
			"expected", "blue is sky the",
			"got", r,
		)
	}

	if r := reverseWords(""); r != "" {
		t.Error(
			"expected", "",
			"got", r,
		)
	}

	if r := reverseWords("dog"); r != "dog" {
		t.Error(
			"expected", "dog",
			"got", r,
		)
	}

}
