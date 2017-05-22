package heap

import "testing"
import "github.com/cristaloleg/golds"

var xFibHeap interface{}

var _ Heap = (*FibHeap)(nil)
var _ golds.Container = (*FibHeap)(nil)

func TestFibHeap(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewFibHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	if value, ok := h.Top(); ok || value != nil {
		t.Error("expected to be empty")
	}
	if value, ok := h.Pop(); ok || value != nil {
		t.Error("expected to be empty")
	}

	if value := h.Size(); value != 0 {
		t.Errorf("want size 0 got %v", value)
	}
	if !h.IsEmpty() {
		t.Errorf("expected to be empty")
	}

	h.Push(1)
	if value := h.Size(); value != 1 {
		t.Errorf("want size 1 got %v", value)
	}
	if h.IsEmpty() {
		t.Errorf("expected to be non-empty")
	}
	h.Pop()

	if value := h.Size(); value != 0 {
		t.Errorf("want size 0 got %v", value)
	}
	if !h.IsEmpty() {
		t.Errorf("expected to be empty")
	}
}

func TestFibHeapMultiple(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewFibHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	for i := 0; i < 10; i++ {
		h.Push(10 - i)
	}
	h.Push(100)

	if value := h.Size(); value != 11 {
		t.Errorf("want size 11 got %v", value)
	}
	if h.IsEmpty() {
		t.Errorf("expected to be non-empty")
	}
	for i := 0; i < 10; i++ {
		value, ok := h.Pop()
		if !ok || value != i+1 {
			t.Errorf("incorrect value, expected %v got %v", i+1, value)
		}
	}

	if value, ok := h.Top(); !ok || value != 100 {
		t.Errorf("expected 100, got %v", value)
	}

	h.Clear()
	if h.Size() != 0 {
		t.Error("expected to be empty")
	}
}

func BenchmarkFibHeapPush(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewFibHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		h.Clear()
		for i := 0; i < 1000; i++ {
			h.Push(i)
		}
	}
	t.StartTimer()
}

func BenchmarkFibHeapPushPop(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewFibHeap(comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 1000; i++ {
			h.Push(i)
		}
		for i := 0; i < 1000; i++ {
			xFibHeap, xFibHeap = h.Pop()
		}

	}
	t.StartTimer()
}
