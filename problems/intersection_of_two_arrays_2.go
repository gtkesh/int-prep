package main

func intersect(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0)

	nums1Map := make(map[int]int, 0)
	nums2Map := make(map[int]int, 0)

	for _, num := range nums1 {
		if _, ok := nums1Map[num]; ok {
			nums1Map[num] += 1
		} else {
			nums1Map[num] = 1
		}
	}

	for _, num := range nums2 {
		if _, ok := nums2Map[num]; ok {
			nums2Map[num] += 1
		} else {
			nums2Map[num] = 1
		}
	}

	for k, v := range nums1Map {
		if val, ok := nums2Map[k]; ok {
			freq := min(v, val)
			for i := 0; i < freq; i++ {
				intersection = append(intersection, k)
			}
		}
	}

	return intersection
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
