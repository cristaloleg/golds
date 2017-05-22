package heap

import (
	"sort"
	"testing"

	"github.com/cristaloleg/golds"
)

var xBinaryHeap interface{}

var _ Heap = (*BinaryHeap)(nil)
var _ golds.Container = (*BinaryHeap)(nil)

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
		value, ok := h.Pop()
		if !ok || value != i {
			t.Errorf("incorrect value, expected %v got %v", i, value)
		}
	}

	if value, ok := h.Top(); !ok || value != 100 {
		t.Errorf("expected 100, got %v", value)
	}
}

func BenchmarkBinaryHeapPush(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewBinaryHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 1000; i++ {
			h.Push(i)
		}
	}
	t.StartTimer()
}

func BenchmarkBinaryHeapPushPop(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewBinaryHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 1000; i++ {
			h.Push(i)
		}
		for i := 0; i < 1000; i++ {
			xBinaryHeap, xBinaryHeap = h.Pop()
		}
	}
	t.StartTimer()
}
