package heap

import (
	"sort"
	"testing"
)

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

	if value := h.Size(); value != 12 {
		t.Errorf("want size 12, got %v", value)
	}

	if value, ok := h.Min(); !ok || value != -90 {
		t.Errorf("expected min value -90, but was %v", value)
	}
	if value, ok := h.Max(); !ok || value != 100 {
		t.Errorf("expected max value 100, but was %v", value)
	}

	h.PopMin()
	h.PopMax()

	tmp := h.Values()
	values := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		values[i] = tmp[i].(int)
	}
	sort.Sort(sort.IntSlice(values))

	for i := 0; i < 10; i++ {
		if values[i] != i {
			t.Errorf("want %v, got %v", i, values[i])
		}
	}

	for i := 0; i < 10; i++ {
		if value, ok := h.PopMin(); !ok || value != i {
			t.Errorf("incorrect min value, expected %v got %v", i, value)
		}

		if value, ok := h.PopMax(); !ok || value != 9-i {
			t.Errorf("incorrect max value, expected %v got %v", i, value)
		}
	}
}
