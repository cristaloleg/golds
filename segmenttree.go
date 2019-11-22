package golds

// SegmentTree represents a segment tree.
type SegmentTree struct {
	size int
	data []int
}

// NewSegmentTree instantiates a new SegmentTree.
func NewSegmentTree(size int) *SegmentTree {
	t := &SegmentTree{
		size: size,
		data: make([]int, 2*size),
	}
	return t
}

// Build will create a tree from a given values.
func (t *SegmentTree) Build(values []int) {
	t.size = len(values)
	t.data = make([]int, 2*t.size)

	for i := 0; i < t.size; i++ {
		t.data[t.size+i] = values[i]
	}
	for i := t.size - 1; i > 0; i-- {
		t.data[i] = t.data[i<<1] + t.data[i<<1|1]
	}
}

// Get returns value of element at index.
func (t *SegmentTree) Get(i int) int {
	if i >= t.size {
		panic("index is too big")
	}
	return t.data[i+t.size]
}

// Set sets a value for element at index.
func (t *SegmentTree) Set(i int, value int) {
	if i >= t.size {
		panic("index is too big")
	}

	i += t.size
	t.data[i] = value
	for ; i > 1; i >>= 1 {
		t.data[i>>1] = t.data[i] + t.data[i^1]
	}
}

// QueryRange returns sum in a range.
func (t *SegmentTree) QueryRange(i, j int) int {
	if i > j {
		panic("i should be less or equal than j")
	}
	if i >= t.size || j >= t.size {
		panic("index is too big")
	}

	var res int
	i += t.size
	j += t.size
	for i < j {
		if (i & 1) == 1 {
			res = res + t.data[i]
			i++
		}
		if (j & 1) == 1 {
			j--
			res = res + t.data[j]

		}
		i >>= 1
		j >>= 1
	}
	return res
}
