package set

import "testing"

func TestSparseBitSet(t *testing.T) {
	s := NewSparseBitSet()
	if s == nil {
		t.Error("cannot instantiate SparseBitSet")
	}

	if s.Count() != 0 {
		t.Error("must be 0")
	}

	s.Set(0)

	if s.Count() != 1 {
		t.Error("must be 1")
	}

	if !s.Get(0) {
		t.Error("must be true")
	}

	s.Toggle(0)

	if s.Get(0) {
		t.Error("must be false")
	}

	s.Toggle(0)

	if !s.Get(0) {
		t.Error("must be true")
	}
}

func TestSparseAnyAndNone(t *testing.T) {
	s := NewSparseBitSet()
	if s == nil {
		t.Error("cannot instantiate SparseBitSet")
	}

	if s.Count() != 0 {
		t.Errorf("must be empty, but have %v", s.Count())
	}

	if s.Any() {
		t.Error("Any must be false")
	}
	if !s.None() {
		t.Error("None must be true")
	}

	s.Set(1)

	if !s.Any() {
		t.Error("Any must be true")
	}
	if s.None() {
		t.Error("None must be false")
	}

	s.Unset(1)
}

func TestSparseBitSetEmpty(t *testing.T) {
	s := NewSparseBitSet()
	if s == nil {
		t.Error("cannot instantiate SparseBitSet")
	}

	s.Unset(10)
	if s.Count() != 0 {
		t.Errorf("must be empty, but have %v", s.Count())
	}
	if s.Get(11) {
		t.Errorf("must be false")
	}

	s.Toggle(10)
	s.Set(11)
	if s.Count() != 2 {
		t.Errorf("must be 2, but have %v", s.Count())
	}

	s.Toggle(10)
	if s.Count() != 1 {
		t.Errorf("must be 1, but have %v", s.Count())
	}

	s.Unset(10)
	if s.Count() != 1 {
		t.Errorf("must be 1, but have %v", s.Count())
	}

	s.Unset(11)
	if s.Count() != 0 {
		t.Errorf("must be empty, but have %v", s.Count())
	}
}
