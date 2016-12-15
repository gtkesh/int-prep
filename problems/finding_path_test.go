package main

import "testing"

func TestContainsNaive(t *testing.T) {
	testMatrix := [][]int{
		{1, 5, 9, 13},
		{2, 6, 10, 14},
		{3, 7, 11, 15},
		{4, 8, 12, 16},
	}

	test := func(matrix [][]int, target int) bool {
		return containsFast(testMatrix, target)
	}

	if result := test(testMatrix, 0); result {
		t.Error(
			"expected", false,
			"got", result,
		)
	}

	if result := test(testMatrix, 10); !result {
		t.Error(
			"expected", true,
			"got", result,
		)
	}

	if result := test(testMatrix, 4); !result {
		t.Error(
			"expected", true,
			"got", result,
		)
	}

	if result := test(testMatrix, 16); !result {
		t.Error(
			"expected", true,
			"got", result,
		)
	}

	if result := test(testMatrix, 17); result {
		t.Error(
			"expected", false,
			"got", result,
		)
	}
}

func TestContainsFast(t *testing.T) {
	testMatrix := [][]int{
		{10, 15, 27, 32},
		{20, 25, 29, 33},
		{30, 35, 37, 39},
		{40, 45, 48, 50},
	}

	test := func(matrix [][]int, target int) bool {
		return containsFast(testMatrix, target)
	}

	if result := test(testMatrix, 0); result {
		t.Error(
			"expected", false,
			"got", result,
		)
	}

	if result := test(testMatrix, 10); !result {
		t.Error(
			"expected", true,
			"got", result,
		)
	}

	if result := test(testMatrix, 40); !result {
		t.Error(
			"expected", true,
			"got", result,
		)
	}

	if result := test(testMatrix, 29); !result {
		t.Error(
			"expected", true,
			"got", result,
		)
	}

	if result := test(testMatrix, 100); result {
		t.Error(
			"expected", false,
			"got", result,
		)
	}
}
