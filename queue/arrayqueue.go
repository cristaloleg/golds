package queue

// ArrayQueue is a queue based on slice
type ArrayQueue struct {
	in   int
	out  int
	data []interface{}
}

// NewArrayQueue returns a pointer to the ArrayQueue
func NewArrayQueue() *ArrayQueue {
	s := &ArrayQueue{
		in:   0,
		out:  0,
		data: make([]interface{}, 0),
	}
	return s
}

// Size return amount of keys in the queue
func (q *ArrayQueue) Size() int {
	return q.in - q.out
}

// IsEmpty returns true if queue is empty
func (q *ArrayQueue) IsEmpty() bool {
	return q.in == q.out
}

// Clear removes all elements from the queue
func (q *ArrayQueue) Clear() {
	q.in, q.out = 0, 0
	q.data = make([]interface{}, 0)
}

// Push adds element to the top of the queue
func (q *ArrayQueue) Push(value interface{}) {
	q.in++
	q.data = append(q.data, value)
}

// PushBulk adds elements to the top of the queue
func (q *ArrayQueue) PushBulk(values ...interface{}) {
	for _, v := range values {
		q.Push(v)
	}
}

// Pop removes and returns top element of the queue
func (q *ArrayQueue) Pop() (value interface{}, ok bool) {
	if q.IsEmpty() {
		return nil, false
	}
	value = q.data[q.out]
	q.out++
	return value, true
}

// Top returns top element of the queue
func (q *ArrayQueue) Top() (value interface{}, ok bool) {
	if q.IsEmpty() {
		return nil, false
	}
	return q.data[q.out], true
}

// Values returns values presented in queue
func (q *ArrayQueue) Values() []interface{} {
	return q.data[:]
}
