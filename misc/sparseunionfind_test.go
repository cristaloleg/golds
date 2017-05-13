package misc

import "testing"

func TestSparseUnionFind(t *testing.T) {
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
		uf.Union(i-1, i)
	}

	if value := uf.Count(); value != 5 {
		t.Errorf("expected 5 got %v", value)
	}

	for i := 2; i < 10; i += 2 {
		uf.Union(i-1, i)
	}

	if value := uf.Count(); value != 1 {
		t.Errorf("expected 1 got %v", value)
	}
}

func BenchmarkSparseUnionFind(t *testing.B) {
	uf := NewSparseUnionFind()
	for i := 0; i < 1000000; i += 2 {
		uf.Union(i, i+1)
	}
}
