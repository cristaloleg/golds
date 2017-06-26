package interval_test

import "testing"
import . "github.com/cristaloleg/golds/interval"

func TestSegmentTree(t *testing.T) {
	t.Parallel()

	s := NewSegmentTree(10)
	if s == nil {
		t.Error("cannot instantiate SegmentTree")
	}

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s.Build(values)

	if value := s.QueryRange(0, 4); value != 10 {
		t.Errorf("want %v, got %v", 10, value)
	}

	s.Modify(1, 12)

	if value := s.QueryRange(0, 4); value != 20 {
		t.Errorf("want %v, got %v", 20, value)
	}

	for i := 0; i < 4; i++ {
		s.Modify(i, 12)
	}

	if value := s.QueryRange(0, 4); value != 48 {
		t.Errorf("want %v, got %v", 48, value)
	}
}
