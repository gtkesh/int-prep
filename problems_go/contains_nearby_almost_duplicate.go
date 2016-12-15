package main

import "math"

/*
Given an array of integers, find out whether there are two distinct indices i and j in the array
such that the difference between nums[i] and nums[j] is at most t and the difference between i and j is at most k.
*/

// TODO:
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	return false
}

func containsNearbyAlmostDuplicateNaive(nums []int, k int, t int) bool {
	for i, _ := range nums {
		for j, _ := range nums {
			diffIndices := i - j
			diffNums := nums[i] - nums[j]
			if i != j && math.Abs(float64(diffNums)) <= float64(t) && math.Abs(float64(diffIndices)) <= float64(k) {
				return true
			}
		}
	}

	return false
}
