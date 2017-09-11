package queue

import (
	"github.com/cristaloleg/golds/set"
)

// UniqueQueue is a queue which allows only unique values in queue
type UniqueQueue struct {
	set   *set.HashSet
	queue *ArrayQueue
}

// NewUniqueQueue returns a pointer to the UniqueQueue
func NewUniqueQueue() *UniqueQueue {
	s := &UniqueQueue{
		set:   set.NewHashSet(),
		queue: NewArrayQueue(),
	}
	return s
}

// Size return amount of keys in the queue
func (s *UniqueQueue) Size() int {
	return s.set.Size()
}

// IsEmpty returns true if queue is empty
func (s *UniqueQueue) IsEmpty() bool {
	return s.set.IsEmpty()
}

// Clear removes all elements from the queue
func (s *UniqueQueue) Clear() {
	s.set.Clear()
	s.queue.Clear()
}

// Flush removes viewed elements from the queue
func (s *UniqueQueue) Flush() {
	s.set.Clear()
}

// Push adds element to the top of the queue
func (s *UniqueQueue) Push(value interface{}) {
	if s.set.Has(value) {
		return
	}
	s.set.Put(value)
	s.queue.Push(value)
}

// PushBulk adds elements to the top of the queue
func (s *UniqueQueue) PushBulk(values ...interface{}) {
	for _, v := range values {
		if !s.set.Has(v) {
			s.set.Put(v)
			s.queue.Push(v)
		}
	}
}

// Pop removes and returns top element of the queue
func (s *UniqueQueue) Pop() (value interface{}, ok bool) {
	if s.IsEmpty() {
		return nil, false
	}
	value, ok = s.queue.Pop()
	if ok {
		s.set.Del(value)
	}
	return value, ok
}

// Top returns top element of the queue
func (s *UniqueQueue) Top() (value interface{}, ok bool) {
	return s.queue.Top()
}

// Values returns values presented in queue
func (s *UniqueQueue) Values() []interface{} {
	return s.queue.Values()
}
