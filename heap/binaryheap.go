package heap

// BinaryHeap ...
type BinaryHeap struct {
	size int
	data []interface{}
	comp func(interface{}, interface{}) bool
}

// NewBinaryHeap instantiates a new BinaryHeap
func NewBinaryHeap(comp func(interface{}, interface{}) bool) *BinaryHeap {
	h := &BinaryHeap{
		size: 0,
		data: make([]interface{}, 0),
		comp: comp,
	}
	return h
}

// NewBinaryHeapSized instantiates a new BinaryHeap
func NewBinaryHeapSized(size int, comp func(interface{}, interface{}) bool) *BinaryHeap {
	h := &BinaryHeap{
		size: 0,
		data: make([]interface{}, size),
		comp: comp,
	}
	return h
}

// Size return amount of keys in the heap
func (h *BinaryHeap) Size() int {
	return h.size
}

// IsEmpty returns true if heap is empty
func (h *BinaryHeap) IsEmpty() bool {
	return h.size == 0
}

// Clear removes all elements from the heap
func (h *BinaryHeap) Clear() {
	h.size = 0
	h.data = nil
}

// Build pushes all items from values to the heap
func (h *BinaryHeap) Build(values []interface{}) {
	size := len(values)
	if len(h.data) < size {
		h.data = make([]interface{}, size)
	}
	copy(h.data, values)
	h.size = size
	for i := h.size/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

// Push adds element to the heap
func (h *BinaryHeap) Push(value interface{}) {
	if len(h.data) >= h.size {
		h.data = append(h.data, nil)
	}
	h.data[h.size] = value
	h.up(h.size)
	h.size++
}

// PushMany adds elements to the heap
func (h *BinaryHeap) PushMany(values ...interface{}) {
	for _, v := range values {
		h.Push(v)
	}
}

// Pop removes and returns top element of the heap
func (h *BinaryHeap) Pop() (value interface{}, ok bool) {
	if h.size == 0 {
		return nil, false
	}
	h.size--
	h.swap(0, h.size)
	h.down(0)
	value = h.data[h.size]
	h.data = h.data[:h.size]
	return value, true
}

// PopMany removes and returns top k elements of the heap
func (h *BinaryHeap) PopMany(k int) (values []interface{}, ok bool) {
	if h.size == 0 {
		return nil, false
	}
	k = min(k, h.size)
	values = make([]interface{}, k)
	for i := 0; i < k; i++ {
		value, _ := h.Pop()
		values[i] = value
	}
	return values, true
}

// Top returns top element of the heap
func (h *BinaryHeap) Top() (value interface{}, ok bool) {
	if h.size == 0 {
		return nil, false
	}
	return h.data[0], true
}

// Values returns values presented in heap
func (h *BinaryHeap) Values() []interface{} {
	return h.data[:h.size]
}

// down pushes element down in the heap-tree
func (h *BinaryHeap) down(i int) {
	for {
		j := 2*i + 1
		if j >= h.size {
			break
		}
		if j2 := j + 1; j2 < h.size && h.comp(h.data[j2], h.data[j]) {
			j = j2
		}
		if h.comp(h.data[i], h.data[j]) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

// up pushes element up in the heap-tree
func (h *BinaryHeap) up(i int) {
	for {
		j := (i - 1) / 2
		if !h.comp(h.data[i], h.data[j]) {
			break
		}
		h.swap(j, i)
		i = j
	}
}

// swap swaps elements on the given indexes
func (h *BinaryHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
