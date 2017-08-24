package tree

import "testing"

func TestAATree(t *testing.T) {
	aa := NewAATree()
	if aa == nil {
		t.Error("fail")
	}
}
