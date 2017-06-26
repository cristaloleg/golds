package tree_test

import (
	"testing"

	. "github.com/cristaloleg/golds/tree"
)

var x interface{}

func TestRedBlackTree(t *testing.T) {
	t.Parallel()

	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	rb := NewRedBlackTree(comp)
	if rb == nil {
		t.Error("cannot instatiate RedBlackTree")
	}

	for i := 0; i < 10; i++ {
		expected := i

		rb.Put(expected)

		if value := rb.Max(); value != expected {
			t.Errorf("expected max %v got %v", expected, value)
		}

		if !rb.Has(expected) {
			t.Errorf("expected to have %v", expected)
		}

		rb.Put(-expected)
		if value := rb.Min(); value != -expected {
			t.Errorf("expected min %v got %v", -expected, value)
		}

		if !rb.Has(-expected) {
			t.Errorf("expected to have %v", -expected)
		}
	}
}

func BenchmarkRedBlackMin(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	rb := NewRedBlackTree(comp)
	for i := 0; i < 1000; i++ {
		rb.Put(i)
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		x = rb.Min()
	}
	t.StartTimer()
}

func BenchmarkRedBlackHas(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	rb := NewRedBlackTree(comp)
	for i := 0; i < 1000; i++ {
		rb.Put(i)
	}

	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		x = rb.Has(i & 1024)
	}
	t.StartTimer()
}
