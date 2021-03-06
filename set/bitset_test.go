package set_test

import "testing"
import . "github.com/cristaloleg/golds/set"

func TestNewBitSet(t *testing.T) {
	t.Parallel()

	s := NewBitSet(50)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	if s.Count() != 0 {
		t.Error("must be 0")
	}

	s.Set(25)
	if s.Count() != 1 {
		t.Error("must be 1")
	}
	if !s.Get(25) {
		t.Error("must be true")
	}

	s.Toggle(25)
	if s.Get(25) {
		t.Error("must be false")
	}

	s.Toggle(25)
	if !s.Get(25) {
		t.Error("must be true")
	}

	s2 := s.Clone()
	if value1, value2 := s.Count(), s2.Count(); value1 != value2 {
		t.Errorf("must be equal %v and %v", value1, value2)
	}
}

func TestAnyAndNone(t *testing.T) {
	t.Parallel()

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

func TestMany(t *testing.T) {
	t.Parallel()

	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	s.SetMany(2, 3, 5, 70)
	if count := s.Count(); count != 4 {
		t.Errorf("must be 4, but was %v", count)
	}

	bools := s.GetMany(2, 3, 5, 70)
	for _, b := range bools {
		if !b {
			t.Error("must be true")
		}
	}

	if s.NoneMany(2, 3, 5, 70) {
		t.Error("must be all true")
	}

	s.UnsetMany(2, 3, 5, 70)
	if count := s.Count(); count != 0 {
		t.Errorf("must be 0, but was %v", count)
	}

	if !s.NoneMany(2, 3, 5, 70) {
		t.Error("must be all false")
	}

	if s.AnyMany(2, 3, 5, 70) {
		t.Error("must be all false")
	}

	s.ToggleMany(2, 3, 5, 70)
	if count := s.Count(); count != 4 {
		t.Errorf("must be 4, but was %v", count)
	}

	if !s.AnyMany(2, 3, 5, 70) {
		t.Error("must be all true")
	}
}

func TestRange(t *testing.T) {
	t.Parallel()

	s := NewBitSet(10)
	if s == nil {
		t.Error("cannot instantiate BitSet")
	}

	s.SetRange(2, 70)
	if count := s.Count(); count != 69 {
		t.Errorf("must be 6, but was %v", count)
	}

	bools := s.GetRange(2, 70)
	for _, b := range bools {
		if !b {
			t.Error("must be true")
		}
	}

	if s.NoneRange(2, 70) {
		t.Error("must be all true")
	}

	s.UnsetRange(2, 70)
	if count := s.Count(); count != 0 {
		t.Errorf("must be 0, but was %v", count)
	}

	if !s.NoneRange(2, 70) {
		t.Error("must be all false")
	}

	if s.AnyRange(2, 70) {
		t.Error("must be all false")
	}

	s.ToggleRange(2, 70)
	if count := s.Count(); count != 69 {
		t.Errorf("must be 6, but was %v", count)
	}

	if !s.AnyRange(2, 70) {
		t.Error("must be all true")
	}
}
