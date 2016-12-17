package data_structures

type Stack struct {
	stack []int
}

func (s *Stack) push(num int) {
	s.stack = append(s.stack, num)
}

func (s *Stack) pop() int {
	num := -1
	if len(s.stack) > 0 {
		num = s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
	}
	return num
}

func (s *Stack) size() int {
	return len(s.stack)
}
