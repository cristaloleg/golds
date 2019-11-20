package golds

import (
	"testing"
)

func TestDisjointSet(t *testing.T) {
	ds := NewDisjointSet(10)
	if ds == nil {
		t.Error("cannot instantiate DisjointSet")
	}

	for i := 0; i < 10; i += 2 {
		ds.Union(i, i+1)
	}

	if value := ds.Count(); value != 5 {
		t.Errorf("expected 5 got %v", value)
	}
	if value := ds.Size(1); value != 2 {
		t.Errorf("expected 2 got %v", value)
	}

	for i := 1; i < 10; i += 2 {
		ds.Union(i-1, i)
	}

	if value := ds.Count(); value != 5 {
		t.Errorf("expected 5 got %v", value)
	}

	for i := 2; i < 10; i += 2 {
		ds.Union(i-1, i)
		if !ds.IsUnited(i-1, i) || !ds.IsUnited(i-1, i) {
			t.Errorf("should %v and %v be Uniond", i, i-1)
		}
	}

	if value := ds.Count(); value != 1 {
		t.Errorf("expected 1 got %v", value)
	}
}

func BenchmarkDisjointSet(b *testing.B) {
	ds := NewDisjointSet(1000)

	b.ResetTimer()
	for i := 0; i < 1000; i += 2 {
		ds.Union(i, i+1)
	}
}
