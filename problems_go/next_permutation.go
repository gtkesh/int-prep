package main

func nextPermutation(nums []int) {
	n := len(nums)

	p := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			p = i - 1
			break
		}
	}

	q := 0
	for i := n - 1; i >= 0; i-- {
		if nums[i] > nums[p] {
			q = i
			break
		}
	}

	if p == 0 && q == 0 {
		reverseSlice(nums, 0, n-1)
		return
	}

	swap(nums, p, q)
	reverseSlice(nums, p+1, n-1)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverseSlice(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
