package queue

import "testing"

func TestNewTopK(t *testing.T) {
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

	expected := []int{7, 17, 23}
	for i := 0; i < 3; i++ {
		if value, ok := q.Pop(); !ok || value != expected[i] {
			t.Errorf("incorrect value, expected %v got %v", expected[i], value)
		}
	}
}
