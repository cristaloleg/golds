package heap_test

import (
	"testing"

	. "github.com/cristaloleg/golds/heap"
)

var xBHeap interface{}

// var _ Heap = (*BHeap)(nil)

// var _ golds.Container = (*BHeap)(nil)

func TestBHeap(t *testing.T) {
	t.Parallel()

	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	h := NewBHeapSized(8, comp)
	if h == nil {
		t.Error("cannot instantiate BHeap")
	}

	_, ok1 := h.Top()
	_, ok2 := h.Pop()
	if ok1 || ok2 || !h.IsEmpty() {
		t.Errorf("expected to be nil")
	}

	for i := 0; i < 15; i++ {
		h.Push(i)
	}
	h.Push(-90)
	// h.Push(100)

	if value := h.Size(); value != 16 {
		t.Errorf("want size 16, got %v", value)
	}

	// tmp := h.Values()
	// values := make([]int, len(tmp))
	// for i := 0; i < len(tmp); i++ {
	// 	values[i] = tmp[i].(int)
	// }
	// sort.Sort(sort.IntSlice(values))

	// for i := 0; i < 10; i++ {
	// 	if values[i] != i {
	// 		t.Errorf("want %v, got %v", i, values[i])
	// 	}
	// }

	if value, ok := h.Top(); !ok || value != -90 {
		t.Errorf("expected %v in top, got %v", -90, value)
	}
	if value, ok := h.Pop(); !ok || value != -90 {
		t.Errorf("expected %v in top, got %v", -90, value)
	}

	for i := 0; i < 10; i++ {
		value, ok := h.Pop()
		if !ok || value != i {
			t.Errorf("incorrect value, expected %v got %v", i, value)
		}
	}

	// if value, ok := h.Top(); !ok || value != 100 {
	// 	t.Errorf("expected 100, got %v", value)
	// }
}

// func TestBHeapBuild(t *testing.T) {
// 	comp := func(a, b interface{}) bool {
// 		return a.(int) < b.(int)
// 	}
// 	h := NewBHeapSized(10, comp)
// 	if h == nil {
// 		t.Error("cannot instantiate BHeap")
// 	}

// 	values := make([]interface{}, 100)
// 	for i := 0; i < 100; i++ {
// 		if i <= 50 {
// 			values[i] = i
// 		} else {
// 			values[i] = 100 - i
// 		}
// 	}
// 	h.Build(values)

// 	if value := h.Size(); value != 100 {
// 		t.Errorf("expected size %v, got %v", 100, value)
// 	}

// 	h.Pop()

// 	for i := 1; i < 50; i++ {
// 		value1, ok1 := h.Pop()
// 		value2, ok2 := h.Pop()
// 		if !ok1 || !ok2 || value1 != i || value2 != i {
// 			t.Errorf("incorrect values, expected %v, got %v and %v", i, value1, value2)
// 		}
// 	}

// 	h.Clear()
// 	if !h.IsEmpty() {
// 		t.Errorf("should be empty")
// 	}
// }

// func BenchmarkBHeapPush(t *testing.B) {
// 	comp := func(a, b interface{}) bool {
// 		return a.(int) < b.(int)
// 	}
// 	h := NewBHeap(comp)
// 	if h == nil {
// 		t.Error("cannot instantiate BHeap")
// 	}

// 	t.ReportAllocs()
// 	t.ResetTimer()
// 	for i := 0; i < t.N; i++ {
// 		for i := 0; i < 1000; i++ {
// 			h.Push(i)
// 		}
// 	}
// 	t.StartTimer()
// }

// func BenchmarkBHeapPushPop(t *testing.B) {
// 	comp := func(a, b interface{}) bool {
// 		return a.(int) < b.(int)
// 	}
// 	h := NewBHeap(comp)
// 	if h == nil {
// 		t.Error("cannot instantiate BHeap")
// 	}

// 	t.ReportAllocs()
// 	t.ResetTimer()
// 	for i := 0; i < t.N; i++ {
// 		for i := 0; i < 1000; i++ {
// 			h.Push(i)
// 		}
// 		for i := 0; i < 1000; i++ {
// 			xBHeap, xBHeap = h.Pop()
// 		}
// 	}
// 	t.StartTimer()
// }
