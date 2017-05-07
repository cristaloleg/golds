package set

import "testing"

func TestNewBitSet(t *testing.T) {
	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	if s.Count() != 0 {
		t.Error("must be 0")
	}

	s.Set(0)

	if s.Count() != 1 {
		t.Error("must be 1")
	}

	if s.Get(0) != true {
		t.Error("must be true")
	}

	s.Toggle(0)

	if s.Get(0) != false {
		t.Error("must be false")
	}

	s.Toggle(0)

	if s.Get(0) != true {
		t.Error("must be true")
	}
}
