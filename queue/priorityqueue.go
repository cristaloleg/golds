package queue

import "github.com/cristaloleg/golds/heap"

// PriorityQueue items sorted by a priority
type PriorityQueue struct {
	data heap.BinaryHeap
}

type pqItem struct {
	value    interface{}
	priority int
}

// NewPriorityQueue instantiates a new PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	comp := func(a, b interface{}) bool {
		return a.(*pqItem).priority > b.(*pqItem).priority
	}
	q := &PriorityQueue{
		data: *heap.NewBinaryHeap(comp),
	}
	return q
}

// Push adds value with a priority to the queue
func (q *PriorityQueue) Push(priority int, value interface{}) {
	item := &pqItem{
		priority: priority,
		value:    value,
	}
	q.data.Push(item)
}

// Pop removes and returns item with a biggest priority in the queue
func (q *PriorityQueue) Pop() (value interface{}, priority int, ok bool) {
	value, ok = q.data.Pop()
	if !ok {
		return nil, 0, false
	}
	return value.(*pqItem).value, value.(*pqItem).priority, true
}

// Top returns item with a biggest priority in the queue
func (q *PriorityQueue) Top() (value interface{}, priority int, ok bool) {
	value, ok = q.data.Top()
	if !ok {
		return nil, 0, false
	}
	return value.(*pqItem).value, value.(*pqItem).priority, true
}

// Values returns all values in the queue
func (q *PriorityQueue) Values() []interface{} {
	size := q.data.Size()
	values := q.data.Values()
	res := make([]interface{}, size)
	for i := 0; i < size; i++ {
		res[i] = values[i].(*pqItem).value
	}
	return res
}
