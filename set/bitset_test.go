package set

import "testing"

func TestNewBitSet(t *testing.T) {
	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	s.Set(0)

	if s.Count() != 1 {
		t.Error("must be 1")
	}
}
