package heap

import "testing"

func TestNewFibHeap(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewFibHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	for i := 0; i < 10; i++ {
		t.Log(h.Top())
		h.Push(10 - i)
	}
	h.Push(100)
	t.Log(h.Top())

	for i := 0; i < 10; i++ {
		value, ok := h.Pop()
		t.Log(value)
		if !ok || value != i+1 {
			t.Errorf("incorrect value, expected %v got %v", i+1, value)
		}
	}

	if value, ok := h.Top(); !ok || value != 100 {
		t.Errorf("expected 100, got %v", value)
	}
}
