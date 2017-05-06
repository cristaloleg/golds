package heap

// BinaryHeap XXX
type BinaryHeap struct {
	data []interface{}
	comp func(interface{}, interface{}) bool
}

// NewBinaryHeap XXX
func NewBinaryHeap(comp func(interface{}, interface{}) bool) *BinaryHeap {
	h := &BinaryHeap{
		data: make([]interface{}, 0),
		comp: comp,
	}
	return h
}

// Size return amount of keys in BinaryHeap
func (h *BinaryHeap) Size() int {
	return len(h.data)
}

// Push XXX
func (h *BinaryHeap) Push(value interface{}) {
	h.data = append(h.data, value)
	if len(h.data) > 1 {
		h.siftUp(len(h.data) - 1)
	}
}

// Pop XXX
func (h *BinaryHeap) Pop() (value interface{}, ok bool) {
	size := len(h.data)
	if size == 0 {
		return nil, false
	}
	value, h.data[0] = h.data[0], h.data[size-1]
	h.data = h.data[:size-1]
	h.siftDown(0)
	return value, true
}

// Top XXX
func (h *BinaryHeap) Top() (value interface{}, ok bool) {
	if len(h.data) == 0 {
		return nil, false
	}
	return h.data[0], true
}

// Values XXX
func (h *BinaryHeap) Values() []interface{} {
	return h.data
}

func (h *BinaryHeap) siftDown(i int) {
	for 2*i+1 < len(h.data) {
		l := 2*i + 1
		r := 2*i + 2
		j := l
		if r < len(h.data) && h.comp(h.data[r], h.data[l]) {
			j = r
		}
		if h.comp(h.data[i], h.data[j]) {
			break
		}
		h.data[j], h.data[i] = h.data[i], h.data[j]
		i = j
	}
}

func (h *BinaryHeap) siftUp(i int) {
	for h.comp(h.data[i], h.data[(i-1)/2]) {
		h.data[(i-1)/2], h.data[i] = h.data[i], h.data[(i-1)/2]
		i = (i - 1) / 2
	}
}
