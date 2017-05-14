package misc

import "testing"

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	if uf == nil {
		t.Error("cannot instantiate UnionFind")
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

func BenchmarkUnionFind(t *testing.B) {
	uf := NewUnionFind(1000)
	for i := 0; i < 1000; i += 2 {
		uf.Union(i, i+1)
	}
}
