package misc_test

import "testing"
import . "github.com/cristaloleg/golds/misc"

var xSparseUnionFind bool

func TestSparseUnionFind(t *testing.T) {
	t.Parallel()

	uf := NewSparseUnionFind()
	if uf == nil {
		t.Error("cannot instantiate SparseUnionFind")
	}

	for i := 0; i < 10; i += 2 {
		uf.Union(i, i+1)
	}

	if value := uf.Count(); value != 5 {
		t.Errorf("expected 5 got %v", value)
	}

	for i := 1; i < 10; i += 2 {
		uf.Union(i, i-1)
	}

	if value := uf.Count(); value != 5 {
		t.Errorf("expected 5 got %v", value)
	}

	for i := 2; i < 10; i += 2 {
		uf.Union(i, i-1)
		if !uf.IsUnited(i, i-1) {
			t.Errorf("%v and %v must be united", i, i-1)
		}
	}

	if value := uf.Count(); value != 1 {
		t.Errorf("expected 1 got %v", value)
	}
}

func BenchmarkSparseUnionFind(t *testing.B) {
	uf := NewSparseUnionFind()
	for i := 0; i < 1000; i += 2 {
		uf.Union(i, i+1)
	}
}

func BenchmarkSparseUnionFindIsUnited(t *testing.B) {
	uf := NewSparseUnionFind()
	for i := 0; i < 1000; i += 2 {
		uf.Union(i, i+1)
	}
	t.ResetTimer()
	for i := 0; i < 1000; i += 2 {
		xSparseUnionFind = uf.IsUnited(i, i+1)
	}

	t.StartTimer()
}
