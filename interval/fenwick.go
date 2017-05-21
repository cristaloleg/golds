package interval

// FenwickTree represents Fenwick tree, aka BIT - Binary indexed tree
type FenwickTree struct {
	size int
	data []int
}

// NewFenwickTree returns a pointer to the FenwickTree
func NewFenwickTree(size int) *FenwickTree {
	t := &FenwickTree{
		size: size,
		data: make([]int, size+1),
	}
	return t
}

// Update increases element at index by value
func (t *FenwickTree) Update(index int, value int) {
	index++
	for ; index <= t.size; index += index & -index {
		t.data[index] += value
	}
}

// Query returns sum on [0, index)
func (t *FenwickTree) Query(index int) int {
	res := 0
	index++
	for ; index > 0; index -= index & -index {
		res += t.data[index]
	}
	return res
}

// QueryRange returns sum in a range
func (t *FenwickTree) QueryRange(i, j int) int {
	return t.Query(j) - t.Query(i-1)
}

// Get returns value of element at index
func (t *FenwickTree) Get(index int) int {
	return t.QueryRange(index, index)
}

// Set sets a value for element at index
func (t *FenwickTree) Set(index int, value int) {
	t.Update(index, value-t.Get(index))
}
