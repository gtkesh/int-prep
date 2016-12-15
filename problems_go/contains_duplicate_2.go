package main

import "math"

/*

Given an array of integers and an integer k,
Find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the difference between i and j is at most k.

*/

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, _ := range nums {
		for j, _ := range nums {
			diff := i - j
			if i != j && nums[i] == nums[j] && math.Abs(float64(diff)) <= float64(k) {
				return true
			}
		}
	}

	return false
}
