package queue_test

import "testing"
import . "github.com/cristaloleg/golds/queue"

func TestNewTopK(t *testing.T) {
	t.Parallel()

	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	q := NewTopK(3, comp)
	if q == nil {
		t.Error("cannot instantiate TopK")
	}

	values := []int{1, 7, 2, 5, 17, 0, 4, 5, 23}
	for _, value := range values {
		q.Push(value)
	}

	if value, ok := q.Top(); !ok || value.(int) < 7 {
		t.Errorf("want greater than %v, got %v", 7, value)
	}

	if value := len(q.Values()); value != 3 {
		t.Errorf("want size %v, got %v", 3, value)
	}

	expected := []int{7, 17, 23}
	for i := 0; i < 3; i++ {
		if value, ok := q.Pop(); !ok || value != expected[i] {
			t.Errorf("incorrect value, expected %v got %v", expected[i], value)
		}
	}
}
