package heap

import "testing"

func TestNewHeap(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewBinaryHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	h.Push(100)

	for i := 0; i < 10; i++ {
		value, ok := h.Pop()
		if !ok || value != i {
			t.Errorf("incorrect value, expected %v got %v", i, value)
		}
	}

	if value, ok := h.Top(); !ok || value != 100 {
		t.Errorf("expected 100, got %v", value)
	}
}
