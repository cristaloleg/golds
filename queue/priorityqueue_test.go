package queue_test

import "testing"
import . "github.com/cristaloleg/golds/queue"

func TestPriorityQueue(t *testing.T) {
	t.Parallel()

	q := NewPriorityQueue()
	if q == nil {
		t.Error("cannot instantiate TopK")
	}

	if _, _, ok := q.Top(); ok {
		t.Errorf("expected to be empty")
	}
	if _, _, ok := q.Pop(); ok {
		t.Errorf("expected to be empty")
	}

	values := []int{1, 7, 2, 5, 17, 0, 4, 5, 23}
	for i, value := range values {
		q.Push(i, value)
	}

	if value, prio, ok := q.Top(); !ok || value.(int) != 23 || prio != len(values)-1 {
		t.Errorf("want greater than %v with priority %v, got %v %v", 7, len(values)-1, value, prio)
	}

	if value := len(q.Values()); value != 9 {
		t.Errorf("want size %v, got %v", 9, value)
	}

	expected := []int{23, 5, 4}
	for i := 0; i < 3; i++ {
		if value, prio, ok := q.Pop(); !ok || value != expected[i] {
			t.Errorf("incorrect value, expected %v got %v with priority %v", expected[i], value, prio)
		}
	}
}
