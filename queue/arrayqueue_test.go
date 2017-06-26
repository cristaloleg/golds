package queue_test

import (
	"testing"

	"github.com/cristaloleg/golds"
	. "github.com/cristaloleg/golds/queue"
)

var _ golds.Container = (*ArrayQueue)(nil)

func TestNewArrayQueue(t *testing.T) {
	t.Parallel()

	s := NewArrayQueue()
	if s == nil {
		t.Error("cannot instantiate ArrayQueue")
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

	for i := 0; i < 10; i++ {
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
