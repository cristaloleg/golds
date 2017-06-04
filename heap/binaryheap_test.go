package heap

import (
	"sort"
	"testing"

	"github.com/cristaloleg/golds"
)

var xBinaryHeap interface{}

var _ Heap = (*BinaryHeap)(nil)
var _ golds.Container = (*BinaryHeap)(nil)

func TestBinaryHeap(t *testing.T) {
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
	if values, ok := h.PopMany(100); ok || values != nil {
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

	h.PushMany(10, 20, 30)
	for i := 1; i <= 3; i++ {
		value, ok := h.Pop()
		if !ok || value != i*10 {
			t.Errorf("incorrect value, expected %v got %v", i*10, value)
		}
	}

	h.PushMany(10, 20)
	values2, ok := h.PopMany(4)
	if !ok || len(values2) != 3 {
		t.Errorf("want size %v, got %v", 3, len(values2))
	}
}

func TestBinaryHeapBuild(t *testing.T) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewBinaryHeapSized(10, comp)
	if h == nil {
		t.Error("cannot instantiate BinaryHeap")
	}

	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		if i <= 50 {
			values[i] = i
		} else {
			values[i] = 100 - i
		}
	}
	h.Build(values)

	if value := h.Size(); value != 100 {
		t.Errorf("expected size %v, got %v", 100, value)
	}

	h.Pop()

	for i := 1; i < 50; i++ {
		value1, ok1 := h.Pop()
		value2, ok2 := h.Pop()
		if !ok1 || !ok2 || value1 != i || value2 != i {
			t.Errorf("incorrect values, expected %v, got %v and %v", i, value1, value2)
		}
	}

	h.Clear()
	if !h.IsEmpty() {
		t.Errorf("should be empty")
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
