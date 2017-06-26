package set_test

import "testing"
import . "github.com/cristaloleg/golds/set"

var _ Set = (*HashSet)(nil)

func TestHashSet(t *testing.T) {
	t.Parallel()

	s := NewHashSet()
	if s == nil {
		t.Error("cannot instantiate a HashSet")
	}

	if s.Size() != 0 {
		t.Error("must be empty")
	}

	s.Put(1)

	if s.Size() != 1 {
		t.Error("must containt 1 element")
	}

	if !s.Has(1) {
		t.Error("must containt element")
	}

	s.Put(1)

	if !s.Has(1) {
		t.Error("must containt element")
	}

	s.Del(1)

	if s.Size() != 0 {
		t.Error("must be empty")
	}
}
