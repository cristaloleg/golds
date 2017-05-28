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

	if value, ok := h.Min(); ok || value != nil {
		t.Errorf("must be empty")
	}
	if value, ok := h.Max(); ok || value != nil {
		t.Errorf("must be empty")
	}

	for i := 0; i < 10; i++ {
		h.Push(i)
		if value, ok := h.Max(); !ok || value != i {
			t.Errorf("want %v, got %v", i, value)
		}

	}
	for i := 0; i < 10; i++ {
		h.Push(-i)
	}

	h.Push(-90)
	h.Push(100)

	if value := h.Size(); value != 22 {
		t.Errorf("want size %v, got %v", 22, value)
	}

	if value, ok := h.Min(); !ok || value != -90 {
		t.Errorf("expected min value -90, but was %v", value)
	}
	if value, ok := h.Max(); !ok || value != 100 {
		t.Errorf("expected max value 100, but was %v", value)
	}
	if value, ok := h.Pop(); !ok || value != -90 {
		t.Errorf("expected max value 10, but was %v", value)
	}
	if value, ok := h.PopMax(); !ok || value != 100 {
		t.Errorf("expected max value 10, but was %v", value)
	}

	for i := 0; i < 10; i++ {
		value, ok := h.Pop()
		if !ok || value != i-9 {
			t.Errorf("incorrect value, expected %v got %v", i-9, value)
		}
	}

	h.Clear()
	if !h.IsEmpty() {
		t.Error("expected to be empty")
	}

	h.PushMany(10, 20, 30)
	if value := h.Size(); value != 3 {
		t.Errorf("want size %v, got %v", 3, value)
	}

	if value, ok := h.Pop(); !ok || value != 10 {
		t.Errorf("expected max value 10, but was %v", value)
	}
	if value, ok := h.PopMax(); !ok || value != 30 {
		t.Errorf("expected max value 20, but was %v", value)
	}
	if value, ok := h.Pop(); !ok || value != 20 {
		t.Errorf("expected max value 20, but was %v", value)
	}
}
