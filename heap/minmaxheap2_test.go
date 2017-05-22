package heap

import "testing"

func TestNewMinMaxHeap2(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewMinMaxHeap2(comp)
	if h == nil {
		t.Error("cannot instantiate MinMaxHeap2")
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
}
