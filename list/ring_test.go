package list

import "testing"

func TestRing(t *testing.T) {
	r := NewRing(10)
	if r == nil {
		t.Error("cannot instantiate Ring")
	}

	if !r.IsEmpty() || r.Size() != 0 {
		t.Error("want empty")
	}
	if value, ok := r.Top(); ok || value != nil {
		t.Errorf("want nil, got %v", value)
	}
	if value, ok := r.Pop(); ok || value != nil {
		t.Errorf("want nil, got %v", value)
	}

	r.Push(10)
	if value, ok := r.Top(); !ok || value != 10 {
		t.Errorf("want %v, got %v", 10, value)
	}

	r.Push(20)
	r.Push(30)
	if value, ok := r.Top(); !ok || value != 10 {
		t.Errorf("want %v, got %v", 10, value)
	}
	if value, ok := r.Pop(); !ok || value != 10 {
		t.Errorf("want %v, got %v", 10, value)
	}
	if value, ok := r.Top(); !ok || value != 20 {
		t.Errorf("want %v, got %v", 10, value)
	}

	r.Clear()
}
