package golds

import "testing"

func TestBITRanged(t *testing.T) {
	tree := NewBITRanged(10)

	tree.Set(5, 10)
	tree.Set(7, 20)

	if res := tree.QueryRange(4, 9); res != 30 {
		t.Fatalf("want %#v, got %#v", 30, res)
	}
}
