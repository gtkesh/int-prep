package main

import (
	"fmt"
	"testing"
)

func TestTwoSumNaive(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	// expectedIndices := []int{0, 1}
	// indices := twoSumNaive(nums, target)

	fmt.Println(twoSum(nums, target))

	// if indices != expectedIndices {
	// 	t.Error(
	// 		"expected", expectedIndices,
	// 		"got", indices,
	// 	)
	// }
}
