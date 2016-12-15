package main

/*
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

Note:
Each element in the result must be unique.
The result can be in any order.

*/
func intersection(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0)
	m := make(map[int]bool, 0)
	intersectionMap := make(map[int]bool)
	for _, num := range nums1 {
		m[num] = true
	}

	for _, num := range nums2 {
		if m[num] {
			intersectionMap[num] = true
		}
	}

	for k, _ := range intersectionMap {
		intersection = append(intersection, k)
	}

	return intersection
}
