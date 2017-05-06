package queue

import "github.com/cristaloleg/golds/heap"

// TopK XXX
type TopK struct {
	k    int
	data heap.BinaryHeap
}

// NewTopK XXX
func NewTopK(k int, comp func(a, b interface{}) bool) *TopK {
	invComp := func(a, b interface{}) bool {
		return comp(a, b)
	}
	q := &TopK{
		k:    k,
		data: *heap.NewBinaryHeap(invComp),
	}
	return q
}

// Push XXX
func (q *TopK) Push(value interface{}) {
	q.data.Push(value)
	if q.data.Size() > q.k {
		q.data.Pop()
	}
}

// Pop XXX
func (q *TopK) Pop() (value interface{}, ok bool) {
	return q.data.Pop()
}

// Top XXX
func (q *TopK) Top() (value interface{}, ok bool) {
	return q.data.Top()
}

// Values XXX
func (q *TopK) Values() []interface{} {
	return q.data.Values()
}
