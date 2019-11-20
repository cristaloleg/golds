package golds

import (
	"testing"
)

func TestFenwickTree(t *testing.T) {
	size := 10
	f := NewBIT(size)
	if f == nil {
		t.Error("cannot instantiate FenwickTree")
	}

	for i := 0; i < size; i++ {
		f.Update(i, size-i)
	}

	for i := 0; i < size; i++ {
		if value := f.Get(i); value != size-i {
			t.Errorf("want %v at %v, got %v", size-i, i, value)
		}
	}

	f.Set(3, 10)
	if value := f.Get(3); value != 10 {
		t.Errorf("want %v at %v, got %v", 10, 3, value)
	}
}
