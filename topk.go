package golds

// TopK structure to store k-biggest values.
type TopK struct {
	k    int
	data *BinaryHeap
}

// NewTopK instantiates a new TopK
func NewTopK(k int, comp CmpFunc) *TopK {
	q := &TopK{
		k:    k,
		data: NewBinaryHeap(comp),
	}
	return q
}

// Push adds new valueб remains only if it's bigger.
func (q *TopK) Push(value interface{}) {
	q.data.Push(value)
	if q.data.Size() > q.k {
		q.data.Pop()
	}
}

// Pop removes top element of top elements
func (q *TopK) Pop() (value interface{}, ok bool) {
	return q.data.Pop()
}

// Top returns top element of top elements
func (q *TopK) Top() (value interface{}, ok bool) {
	return q.data.Top()
}

// Values returns all values, unsorted
func (q *TopK) Values() []interface{} {
	return q.data.Values()
}
