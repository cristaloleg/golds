package maps_test

import (
	"testing"

	"github.com/cristaloleg/golds"
	. "github.com/cristaloleg/golds/maps"
)

var _ Map = (*HashMap)(nil)
var _ golds.Container = (*HashMap)(nil)

func TestHashMap(t *testing.T) {
	t.Parallel()

	m := NewHashMap()
	if m == nil {
		t.Error("cannot instantiate a HashMap")
	}

	if m.Size() != 0 {
		t.Error("must be empty")
	}

	m.Put(1, 2)

	if m.Size() != 1 {
		t.Error("must containt 1 element")
	}

	if !m.Has(1) {
		t.Error("must containt element")
	}

	m.Put(1, 3)

	if value, ok := m.Get(1); value != 3 || !ok {
		t.Error("incorrect value")
	}

	m.Del(1)

	if m.Size() != 0 {
		t.Error("must be empty")
	}
}
