package heap

import "testing"

func TestNewMinMaxHeap(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewMinMaxHeap(comp)
	if h == nil {
		t.Error("cannot instantiate MinMaxHeap")
	}

	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	h.Push(-90)
	h.Push(100)

	if value, ok := h.Min(); !ok || value != -90 {
		t.Errorf("expected min value -90, but was %v", value)
	}
	if value, ok := h.Max(); !ok || value != 100 {
		t.Errorf("expected max value 100, but was %v", value)
	}

	h.PopMin()
	h.PopMax()

	for i := 0; i < 10; i++ {
		if value, ok := h.PopMin(); !ok || value != i {
			t.Errorf("incorrect min value, expected %v got %v", i, value)
		}

		if value, ok := h.PopMax(); !ok || value != 9-i {
			t.Errorf("incorrect max value, expected %v got %v", i, value)
		}
	}
}
