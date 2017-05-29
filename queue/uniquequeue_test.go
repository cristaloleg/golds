package queue

import "testing"

func TestUniqueQueue(t *testing.T) {
	q := NewUniqueQueue()
	if q == nil {
		t.Error("cannot instantiate UniqueQueue")
	}

	if _, ok := q.Top(); ok {
		t.Error("want empty")
	}
	if _, ok := q.Pop(); ok {
		t.Error("want empty")
	}

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	if value := q.Size(); value != 10 {
		t.Errorf("want size %v, got %v", 10, value)
	}

	for i := 0; i < 5; i++ {
		q.Push(i)
	}

	if value := q.Size(); value != 10 {
		t.Errorf("want size %v, got %v", 10, value)
	}

	for i := 0; i < 5; i++ {
		if value, ok := q.Pop(); !ok || value != i {
			t.Errorf("want %v, got %v", i, value)
		}
	}

	q.Clear()

	q.PushBulk(1, 2, 2, 1, 3)
	if value := q.Size(); value != 3 {
		t.Errorf("want size %v, got %v", 3, value)
	}

	values := q.Values()
	for i, v := range values {
		if i+1 != v {
			t.Errorf("want %v, got %v", i+1, v)
		}
	}
}
