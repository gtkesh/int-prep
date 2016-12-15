package main

/*

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

func twoSumNaive(nums []int, target int) []int {
	indices := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				indices[0], indices[1] = i, j
				return indices
			}
		}
	}

	return indices
}

func twoSum1(nums []int, target int) []int {

	table := map[int]int{}

	for i, item := range nums {
		table[item] = i
	}

	for i, item := range nums {
		complement := target - item
		if val, ok := table[complement]; ok && val != i {
			return []int{i, val}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {

	table := map[int]int{}

	for i, item := range nums {
		complement := target - item
		if val, ok := table[complement]; ok && val != i {
			return []int{val, i}
		}

		table[item] = i
	}

	return []int{}
}
