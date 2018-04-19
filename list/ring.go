package list

// Ring is a ring-buffer, aka circular list
type Ring struct {
	size int
	in   int
	out  int
	data []interface{}
}

// NewRing instantiates a new Ring
func NewRing(size int) *Ring {
	r := &Ring{
		size: size,
		data: make([]interface{}, size),
	}
	return r
}

// Size return amount of keys in the ring
func (r *Ring) Size() int {
	return (r.in - r.out + r.size) % r.size
}

// IsEmpty returns true if ring is empty
func (r *Ring) IsEmpty() bool {
	return r.Size() == 0
}

// Clear removes all elements from the ring
func (r *Ring) Clear() {
	r.size, r.in, r.out = 0, 0, 0
}

// Push adds element to the ring
func (r *Ring) Push(value interface{}) {
	r.data[r.in] = value
	r.in = (r.in + 1) % r.size
	if r.in == r.out {
		r.out = (r.out + 1) % r.size
	}
}

// Pop removes and returns top element of the ring
func (r *Ring) Pop() (interface{}, bool) {
	if r.IsEmpty() {
		return nil, false
	}
	value := r.data[r.out]
	r.out = (r.out + 1) % r.size
	return value, true
}

// Top returns top element of the ring
func (r *Ring) Top() (interface{}, bool) {
	if r.IsEmpty() {
		return nil, false
	}
	return r.data[r.out], true
}
