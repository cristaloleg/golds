package set

import "testing"

func TestNewBitSet(t *testing.T) {
	s := NewBitSet(5)
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

func TestAnyAndNone(t *testing.T) {
	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
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

func TestBulk(t *testing.T) {
	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	s.SetBulk(2, 3, 5, 7)
	if count := s.Count(); count != 4 {
		t.Errorf("must be 4, but was %v", count)
	}

	s.UnsetBulk(2, 3, 5, 7)
	if count := s.Count(); count != 0 {
		t.Errorf("must be 0, but was %v", count)
	}

	s.ToggleBulk(2, 3, 5, 7)
	if count := s.Count(); count != 4 {
		t.Errorf("must be 4, but was %v", count)
	}
}

func TestRange(t *testing.T) {
	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	s.SetRange(2, 7)
	if count := s.Count(); count != 6 {
		t.Errorf("must be 6, but was %v", count)
	}

	s.UnsetRange(2, 7)
	if count := s.Count(); count != 0 {
		t.Errorf("must be 0, but was %v", count)
	}

	s.ToggleRange(2, 7)
	if count := s.Count(); count != 6 {
		t.Errorf("must be 6, but was %v", count)
	}
}
