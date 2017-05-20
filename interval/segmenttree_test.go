package interval

import "testing"

func TestSegmentTree(t *testing.T) {
	s := NewSegmentTree(10)
	if s == nil {
		t.Error("cannot instantiate SegmentTree")
	}

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s.Build(values)
	t.Log(s.Values())

	if value := s.QueryRange(0, 4); value != 10 {
		t.Errorf("want %v, got %v", 10, value)
	}

	s.Modify(1, 12)
	t.Log(s.Values())

	if value := s.QueryRange(0, 4); value != 20 {
		t.Errorf("want %v, got %v", 20, value)
	}

	for i := 0; i < 4; i++ {
		s.Modify(i, 12)
	}
	s.ModifyRange(0, 4, 12)
	t.Error(s.Values())

	if value := s.QueryRange(0, 4); value != 48 {
		t.Errorf("want %v, got %v", 48, value)
	}
}
