package main

func rob(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	memo := make([]int, n)
	memo[0] = nums[0]
	memo[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		memo[i] = max(memo[i-2]+nums[i], memo[i-1])
	}

	return memo[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
