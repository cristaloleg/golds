package heap

// MinMaxHeap XXX
type MinMaxHeap struct {
	min BinaryHeap
	max BinaryHeap
}

// NewMinMaxHeap XXX
func NewMinMaxHeap(comp func(a, b interface{}) bool) *MinMaxHeap {
	inv := func(a, b interface{}) bool {
		return comp(b, a)
	}
	h := &MinMaxHeap{
		min: *NewBinaryHeap(comp),
		max: *NewBinaryHeap(inv),
	}
	return h
}

// Size return amount of keys in BinaryHeap
func (h *MinMaxHeap) Size() int {
	return h.min.Size()
}

//Push XXX
func (h *MinMaxHeap) Push(value interface{}) {
	h.min.Push(value)
	h.max.Push(value)
}

// PopMin XXX
func (h *MinMaxHeap) PopMin() (value interface{}, ok bool) {
	return h.min.Pop()
}

// PopMax XXX
func (h *MinMaxHeap) PopMax() (value interface{}, ok bool) {
	return h.max.Pop()
}

// Min XXX
func (h *MinMaxHeap) Min() (value interface{}, ok bool) {
	return h.min.Top()
}

// Max XXX
func (h *MinMaxHeap) Max() (value interface{}, ok bool) {
	return h.max.Top()
}

// Values XXX
func (h *MinMaxHeap) Values() []interface{} {
	return h.min.Values()
}
