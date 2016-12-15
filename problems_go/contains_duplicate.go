package main

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, num := range nums {
		val, ok := m[num]
		if ok {
			m[num] = val + 1
		} else {
			m[num] = 1
		}
	}

	for _, v := range m {
		if v > 1 {
			return true
		}
	}

	return false
}
