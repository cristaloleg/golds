package interval

// FenwickTree XXX
type FenwickTree struct {
	data []int
}

// NewFenwickTree XXX
func NewFenwickTree() *FenwickTree {
	t := &FenwickTree{}
	return t
}

// Update XXX
func (t *FenwickTree) Update(index int, value int) {
	size := len(t.data)
	for ; index <= size; index |= index + 1 {
		t.data[index] += value
	}
}

// Query XXX
func (t *FenwickTree) Query(index int) int {
	res := 0
	for ; index >= 0; index -= index & -index {
		res += t.data[index]
	}
	return res
}

// QueryRange XXX
func (t *FenwickTree) QueryRange(i, j int) int {
	return t.Query(j) - t.Query(i-1)
}

// Get XXX
func (t *FenwickTree) Get(index int) int {
	return t.QueryRange(index, index)
}

// Set XXX
func (t *FenwickTree) Set(index int, value int) {
	t.Update(index, value-t.Get(index))
}
