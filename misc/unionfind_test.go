package misc

import "testing"

var xUnionFind bool

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
		if !uf.IsUnited(i-1, i) || !uf.IsUnited(i-1, i) {
			t.Errorf("should %v and %v be Uniond", i, i-1)
		}
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

func BenchmarkUnionFindIsUnited(t *testing.B) {
	uf := NewSparseUnionFind()
	for i := 0; i < 1000; i += 2 {
		uf.Union(i, i+1)
	}
	t.ResetTimer()
	for i := 0; i < 1000; i += 2 {
		xUnionFind = uf.IsUnited(i, i+1)
	}

	t.StartTimer()
}
