package tree

import "testing"

func TestAvlTree(t *testing.T) {
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

		if value := avl.Max(); value != expected {
			t.Errorf("expected max %v got %v", expected, value)
		}

		if !avl.Has(expected) {
			t.Errorf("expected to have %v", expected)
		}

		avl.Put(-expected)
		if value := avl.Min(); value != -expected {
			t.Errorf("expected min %v got %v", -expected, value)
		}

		if !avl.Has(-expected) {
			t.Errorf("expected to have %v", -expected)
		}
	}
}

func BenchmarkAvlMin(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	rb := NewAvlTree(comp)
	for i := 0; i < 1000; i++ {
		rb.Put(i)
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		x = rb.Min()
	}
	t.StartTimer()
}

func BenchmarkAvlHas(t *testing.B) {
	comp := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	rb := NewAvlTree(comp)
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
