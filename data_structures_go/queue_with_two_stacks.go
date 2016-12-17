package main

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

type Queue struct {
	stack1 Stack
	stack2 Stack
}

func (q *Queue) enqueue(num int) {
	n := q.stack2.size()
	if n > 0 {
		for i := 0; i < n; i++ {
			q.stack1.push(q.stack2.pop())
		}
	}

	q.stack1.push(num)
}

func (q *Queue) dequeue() (int, bool) {
	if q.stack2.size() > 0 {
		return q.stack2.pop(), true
	}

	n := q.stack1.size()
	for i := 0; i < n; i++ {
		q.stack2.push(q.stack1.pop())
	}

	if q.stack2.size() > 0 {
		return q.stack2.pop(), true
	}

	return -1, false
}

func (q *Queue) size() int {
	return q.stack1.size() + q.stack2.size()

}
