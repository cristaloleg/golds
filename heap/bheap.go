package heap

// BHeap is a B-Heap structure
type BHeap struct {
	data   []interface{}
	comp   func(interface{}, interface{}) bool
	size   int
	mask   int
	shift  int
	rows   int
	length int
	next   int
}

const rowShift int = 16
const rowWidth int = 1 << uint(rowShift)

// NewBHeapSized instantiates a new BHeap
func NewBHeapSized(size int, comp func(interface{}, interface{}) bool) *BHeap {
	u := 1
	for ; (1 << uint(u)) != size; u++ {
	}

	r := 16
	h := &BHeap{
		data:   make([]interface{}, r),
		comp:   comp,
		size:   size,
		mask:   size - 1,
		shift:  u,
		rows:   r,
		length: 0,
		next:   1,
	}
	h.addRow()
	h.set(1, nil)
	return h
}

// Size return amount of keys in the heap
func (h *BHeap) Size() int {
	return h.next - 1
}

// IsEmpty returns true if heap is empty
func (h *BHeap) IsEmpty() bool {
	return h.next == 1
}

// Clear removes all elements from the heap
// func (h *BHeap) Clear() {
// 	h.size = 0
// 	h.data = nil
// }

// Build pushes all items from values to the heap
// func (h *BHeap) Build(values []interface{}) {
// 	size := len(values)
// 	if len(h.data) < size {
// 		h.data = make([]interface{}, size)
// 	}
// 	copy(h.data, values)
// 	h.size = size
// 	for i := h.size/2 - 1; i >= 0; i-- {
// 		h.down(i)
// 	}
// }

// Push adds element to the heap
func (h *BHeap) Push(value interface{}) {
	if h.length == h.next {
		h.addRow()
	}
	u := h.next
	h.next++
	h.set(u, value)
	h.up(u)
}

// Pop removes and returns top element of the heap
func (h *BHeap) Pop() (value interface{}, ok bool) {
	if h.IsEmpty() {
		return nil, false
	}
	value = h.get(1)
	h.swap(1, h.next-1)
	h.next--
	h.down(1)
	return value, true
}

// Top returns top element of the heap
func (h *BHeap) Top() (value interface{}, ok bool) {
	if h.IsEmpty() {
		return nil, false
	}
	return h.get(1), true
}

// Values returns values presented in heap
// func (h *BHeap) Values() []interface{} {
// 	return h.data[:h.next]
// }

func (h *BHeap) down(u int) {
	for {
		v1, v2 := h.childs(u)
		if v1 >= h.next {
			return
		}
		if v1 != v2 && v2 < h.next && h.comp(h.get(v2), h.get(v1)) {
			v1 = v2
		}
		if h.comp(h.get(u), h.get(v1)) {
			return
		}
		h.swap(u, v1)
		u = v1
	}
}

func (h *BHeap) up(u int) {
	for u > 1 {
		v := h.parent(u)
		if !h.comp(h.get(u), h.get(v)) {
			break
		}
		h.swap(u, v)
		u = v
	}
}

func (h *BHeap) get(i int) interface{} {
	x := i & (rowWidth - 1)
	y := i >> uint(rowShift)
	return h.data[y*rowWidth+x]
}

func (h *BHeap) set(i int, value interface{}) {
	x := i & (rowWidth - 1)
	y := i >> uint(rowShift)
	h.data[y*rowWidth+x] = value
}

// swap swaps elements on the given indexes
func (h *BHeap) swap(i, j int) {
	ii, jj := h.get(i), h.get(j)
	h.set(i, jj)
	h.set(j, ii)
}

func (h *BHeap) parent(u int) int {
	po := u & h.mask
	if u < h.size || po > 3 {
		return (u & (^h.mask)) | (po >> 1)
	}
	if po > 1 {
		return u - 2
	}
	v := (u - h.size) >> uint(h.shift)
	v += v & (^(h.mask >> 1))
	v |= h.size / 2
	return v
}

func (h *BHeap) childs(u int) (a, b int) {
	if u > h.mask && (u&(h.mask-1)) == 0 {
		return u + 2, u + 2
	}
	if u&(h.size>>1) == 0 {
		tmp := u + (u & h.mask)
		return tmp, tmp + 1
	}
	a = (u & (^h.mask)) >> 1
	a |= u & (h.mask >> 1)
	a++
	uu := a << uint(h.shift)
	return uu, uu + 1
}

func (h *BHeap) addRow() {
	h.data = append(h.data, make([]interface{}, rowWidth))
	h.length += rowWidth
}
