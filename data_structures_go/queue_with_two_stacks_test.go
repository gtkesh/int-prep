package data_structures

import "testing"

func TestQueueEnqueue(t *testing.T) {
	q := new(Queue)

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	if q.size() != 3 {
		t.Error(
			"expected", 3,
			"got", q.size(),
		)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := new(Queue)

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	if q.size() != 3 {
		t.Error(
			"expected", 3,
			"got", q.size(),
		)
	}

	item, _ := q.dequeue()
	if item != 1 {
		t.Error(
			"expected", 1,
			"got", item,
		)
	}
	if q.size() != 2 {
		t.Error(
			"expected", 2,
			"got", q.size(),
		)
	}

	q.enqueue(4)
	if q.size() != 3 {
		t.Error(
			"expected", 3,
			"got", q.size(),
		)
	}

	q.enqueue(5)
	if q.size() != 4 {
		t.Error(
			"expected", 4,
			"got", q.size(),
		)
	}

	item, _ = q.dequeue()
	if item != 2 {
		t.Error(
			"expected", 2,
			"got", item,
		)
	}
	if q.size() != 3 {
		t.Error(
			"expected", 3,
			"got", q.size(),
		)
	}

	item, _ = q.dequeue()
	if item != 3 {
		t.Error(
			"expected", 3,
			"got", item,
		)
	}
	if q.size() != 2 {
		t.Error(
			"expected", 2,
			"got", q.size(),
		)
	}

	item, _ = q.dequeue()
	if item != 4 {
		t.Error(
			"expected", 4,
			"got", item,
		)
	}
	if q.size() != 1 {
		t.Error(
			"expected", 1,
			"got", q.size(),
		)
	}

	item, _ = q.dequeue()
	if item != 5 {
		t.Error(
			"expected", 5,
			"got", item,
		)
	}
	if q.size() != 0 {
		t.Error(
			"expected", 0,
			"got", q.size(),
		)
	}
}

func TestStackPush(t *testing.T) {
	s := new(Stack)

	s.push(1)
	s.push(2)
	s.push(3)

	if len(s.stack) != 3 {
		t.Error(
			"expected", 3,
			"got", len(s.stack),
		)
	}

}

func TestStackPop(t *testing.T) {
	s := new(Stack)

	s.push(1)
	s.push(2)
	s.push(3)

	if len(s.stack) != 3 {
		t.Error(
			"expected", 3,
			"got", len(s.stack),
		)
	}

	item := s.pop()
	if item != 3 {
		t.Error(
			"expected", 3,
			"got", item,
		)
	}
	if len(s.stack) != 2 {
		t.Error(
			"expected", 2,
			"got", len(s.stack),
		)
	}

	item = s.pop()
	if item != 2 {
		t.Error(
			"expected", 2,
			"got", item,
		)
	}
	if len(s.stack) != 1 {
		t.Error(
			"expected", 1,
			"got", len(s.stack),
		)
	}

	item = s.pop()
	if item != 1 {
		t.Error(
			"expected", 1,
			"got", item,
		)
	}
	if len(s.stack) != 0 {
		t.Error(
			"expected", 0,
			"got", len(s.stack),
		)
	}

}
