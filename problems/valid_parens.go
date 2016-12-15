package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, c := range s {
		if _, ok := m[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || c != m[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
