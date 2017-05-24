package heap

import "math"

// MinMaxHeap2 new implementation of min-max heap
type MinMaxHeap2 struct {
	data []interface{}
	comp func(interface{}, interface{}) bool
}

// NewMinMaxHeap2 return a pointer to the MinMaxHeap2
func NewMinMaxHeap2(comp func(a, b interface{}) bool) *MinMaxHeap2 {
	h := &MinMaxHeap2{
		data: make([]interface{}, 0),
		comp: comp,
	}
	return h
}

// Size return amount of keys in the heap
func (h *MinMaxHeap2) Size() int {
	return len(h.data)
}

// IsEmpty returns true if heap is empty
func (h *MinMaxHeap2) IsEmpty() bool {
	return len(h.data) == 0
}

// Clear removes all elements from the heap
func (h *MinMaxHeap2) Clear() {
	h.data = nil
}

// Push adds element to the heap
func (h *MinMaxHeap2) Push(value interface{}) {
	h.data = append(h.data, value)
	h.bubbleUp(len(h.data) - 1)
}

// PushBulk adds elements to the heap
func (h *MinMaxHeap2) PushBulk(values ...interface{}) {
	for _, v := range values {
		h.Push(v)
	}
}

// Min returns min element of the heap
func (h *MinMaxHeap2) Min() (value interface{}, ok bool) {
	if len(h.data) == 0 {
		return nil, false
	}
	return h.data[0], true
}

// Max returns max element of the heap
func (h *MinMaxHeap2) Max() (value interface{}, ok bool) {
	size := len(h.data)
	if size == 0 {
		return nil, false
	}
	if size == 1 {
		return h.data[0], true
	}
	if size == 2 {
		return h.data[1], true
	}
	if h.comp(h.data[1], h.data[2]) {
		return h.data[2], true
	}
	return h.data[1], true
}

func (h *MinMaxHeap2) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinMaxHeap2) bubbleUp(index int) {
	isMin := h.level(index)%2 == 0
	parentIndex := h.parentIndex(index)

	if parentIndex == -1 {
		return
	}
	if isMin != h.comp(h.data[index], h.data[parentIndex]) {
		h.swap(index, parentIndex)
		h.bubbleUpAuxIter(parentIndex, !isMin)
	} else {
		h.bubbleUpAuxIter(index, isMin)
	}
}

func (h *MinMaxHeap2) bubbleUpAuxIter(index int, isMin bool) {
	gpIndex := h.gpIndex(index)
	for gpIndex != -1 {
		if isMin != h.comp(h.data[index], h.data[gpIndex]) {
			break
		}
		h.swap(index, gpIndex)
		index, gpIndex = gpIndex, h.gpIndex(gpIndex)
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
