package list

import (
	"testing"

	"github.com/cristaloleg/golds"
)

var _ golds.Container = (*ArrayStack)(nil)

func TestNewArrayStack(t *testing.T) {
	s := NewArrayStack()
	if s == nil {
		t.Error("cannot instantiate ArrayStack")
	}

	if !s.IsEmpty() {
		t.Error("expected to be empty")
	}
	if value, ok := s.Top(); ok || value != nil {
		t.Error("expected to be empty")
	}
	if value, ok := s.Pop(); ok || value != nil {
		t.Error("expected to be empty")
	}

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	if s.Size() != 10 {
		t.Error("size have 10 elements")
	}

	for i := 10 - 1; i >= 0; i-- {
		value, ok := s.Top()
		if !ok || value != i {
			t.Errorf("expected %v but was %v", i, value)
		}

		value, ok = s.Pop()
		if !ok || value != i {
			t.Errorf("expected %v but was %v", i, value)
		}
	}
}
