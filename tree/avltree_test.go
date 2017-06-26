package tree_test

import "testing"
import . "github.com/cristaloleg/golds/tree"

var xAvlTest interface{}

func TestAvlTree(t *testing.T) {
	t.Parallel()

	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	avl := NewAvlTree(comp)
	if avl == nil {
		t.Error("cannot instantiate AvlTree")
	}

	for i := 0; i < 10; i++ {
		expected := i

		avl.Put(expected)

		if value, ok := avl.Max(); !ok || value != expected {
			t.Errorf("expected max %v got %v", expected, value)
		}

		if !avl.Has(expected) {
			t.Errorf("expected to have %v", expected)
		}

		avl.Put(-expected)
		if value, ok := avl.Min(); !ok || value != -expected {
			t.Errorf("expected min %v got %v", -expected, value)
		}

		if !avl.Has(-expected) {
			t.Errorf("expected to have %v", -expected)
		}
	}
}

func TestAvlTreeEmpty(t *testing.T) {
	t.Parallel()

	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	avl := NewAvlTree(comp)
	if avl == nil {
		t.Error("cannot instantiate AvlTree")
	}

	if value := avl.Size(); value != 0 {
		t.Errorf("expected to be empty, but was %v", value)
	}
	if !avl.IsEmpty() {
		t.Error("expected to be empty")
	}
	if value, ok := avl.Min(); ok || value != nil {
		t.Error("expected to be empty")
	}
	if value, ok := avl.Max(); ok || value != nil {
		t.Error("expected to be empty")
	}
	if avl.Has(1) {
		t.Errorf("expected to not have 1")
	}

	avl.Put(1)

	if !avl.Has(1) {
		t.Errorf("expected to have 1")
	}
	if value := avl.Size(); value != 1 {
		t.Errorf("expected to be non-empty, but was %v", value)
	}
	if avl.IsEmpty() {
		t.Error("expected to be non-empty")
	}
	if value, ok := avl.Min(); !ok || value != 1 {
		t.Error("expected to be non-empty")
	}
	if value, ok := avl.Max(); !ok || value != 1 {
		t.Error("expected to be non-empty")
	}

	avl.Clear()

	if value := avl.Size(); value != 0 {
		t.Errorf("expected to be empty, but was %v", value)
	}

}

func BenchmarkAvlMin(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	avl := NewAvlTree(comp)
	for i := 0; i < 1000; i++ {
		avl.Put(i)
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		xAvlTest, _ = avl.Min()
	}
	t.StartTimer()
}

func BenchmarkAvlHas(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	avl := NewAvlTree(comp)
	for i := 0; i < 1000; i++ {
		avl.Put(i)
	}

	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		xAvlTest = avl.Has(i & 1024)
	}
	t.StartTimer()
}
