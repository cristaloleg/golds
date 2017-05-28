package heap

import "math"

// MinMaxHeap2 new implementation of min-max heap
type MinMaxHeap2 struct {
	size int
	data []interface{}
	comp func(interface{}, interface{}) bool
}

// NewMinMaxHeap2 return a pointer to the MinMaxHeap2
func NewMinMaxHeap2(comp func(a, b interface{}) bool) *MinMaxHeap2 {
	h := &MinMaxHeap2{
		size: 0,
		data: make([]interface{}, 0),
		comp: comp,
	}
	return h
}

// Size return amount of keys in the heap
func (h *MinMaxHeap2) Size() int {
	return h.size
}

// IsEmpty returns true if heap is empty
func (h *MinMaxHeap2) IsEmpty() bool {
	return h.size == 0
}

// Clear removes all elements from the heap
func (h *MinMaxHeap2) Clear() {
	h.size = 0
	h.data = nil
}

// Push adds element to the heap
func (h *MinMaxHeap2) Push(value interface{}) {
	h.size++
	h.data = append(h.data, value)
	h.up(h.size - 1)
}

// PushMany adds elements to the heap
func (h *MinMaxHeap2) PushMany(values ...interface{}) {
	for _, v := range values {
		h.Push(v)
	}
}

// PopMin removes and returns top element of the heap
func (h *MinMaxHeap2) PopMin() (value interface{}, ok bool) {
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

// Pop removes and returns top element of the heap
func (h *MinMaxHeap2) Pop() (value interface{}, ok bool) {
	return h.PopMin()
}

// Top returns top element of the heap
func (h *MinMaxHeap2) Top() (value interface{}, ok bool) {
	return h.Min()
}

// Min returns min element of the heap
func (h *MinMaxHeap2) Min() (value interface{}, ok bool) {
	if h.size == 0 {
		return nil, false
	}
	return h.data[0], true
}

// Max returns max element of the heap
func (h *MinMaxHeap2) Max() (value interface{}, ok bool) {
	size := len(h.data)
	switch {
	case size == 0:
		return nil, false
	case size == 1:
		return h.data[0], true
	case size == 2:
		return h.data[1], true
	default:
		if h.comp(h.data[1], h.data[2]) {
			return h.data[2], true
		}
		return h.data[1], true
	}
}

func (h *MinMaxHeap2) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinMaxHeap2) up(index int) {
	isMin := h.level(index)%2 == 0
	parentIndex := h.parentIndex(index)

	if parentIndex == -1 {
		return
	}
	if isMin != h.comp(h.data[index], h.data[parentIndex]) {
		h.swap(index, parentIndex)
		h.upAux(parentIndex, !isMin)
	} else {
		h.upAux(index, isMin)
	}
}

func (h *MinMaxHeap2) upAux(index int, isMin bool) {
	gpIndex := h.gpIndex(index)
	for gpIndex != -1 {
		if isMin != h.comp(h.data[index], h.data[gpIndex]) {
			break
		}
		h.swap(index, gpIndex)
		index, gpIndex = gpIndex, h.gpIndex(gpIndex)
	}
}

func (h *MinMaxHeap2) down(index int) {
	isMin := h.level(index)%2 == 0
	h.downAux(index, isMin)
}

func (h *MinMaxHeap2) downAux(index int, isMin bool) {
	// FIXME
	m := 0

	if h.parentIndex(h.parentIndex(index)) == m {

		if isMin == h.comp(h.data[m], h.data[index]) {
			h.swap(index, m)

			if isMin != h.comp(h.data[m], h.parentIndex(m)) {
				h.swap(m, h.parentIndex(m))
			}
			h.downAux(m, isMin)
		}
	} else if isMin == h.comp(h.data[m], h.data[index]) {
		h.swap(index, m)
	}
}

func (h *MinMaxHeap2) level(index int) int {
	if index == 0 {
		return 0
	}
	return int(math.Log(float64(index+1.0)) / math.Log(2.0))
}

func (h *MinMaxHeap2) parentIndex(index int) int {
	if index <= 0 {
		return -1
	}
	return (index - 1) / 2
}

func (h *MinMaxHeap2) gpIndex(index int) int {
	return h.parentIndex(h.parentIndex(index))
}
