package heap

import "testing"
import "sort"

func TestNewHeap(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewBinaryHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	_, ok1 := h.Top()
	_, ok2 := h.Pop()
	if ok1 || ok2 {
		t.Errorf("expected to be nil")
	}

	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	h.Push(100)

	if value := h.Size(); value != 11 {
		t.Errorf("want size 11, got %v", value)
	}

	values := h.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i].(int) < values[j].(int)
	})

	for i := 0; i < 10; i++ {
		if values[i] != i {
			t.Errorf("want %v, got %v", i, values[i])
		}
	}

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
