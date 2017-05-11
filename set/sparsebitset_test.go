package set

import "testing"

func TestNewSparseBitSet(t *testing.T) {
	s := NewSparseBitSet(5)
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
	s := NewSparseBitSet(10)
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
