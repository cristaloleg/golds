package interval

// FenwickTreeSimple represents Fenwick tree, aka BIT - Binary indexed tree
type FenwickTreeSimple struct {
	data []int
}

// NewFenwickTreeSimple returns a pointer to the FenwickTree
func NewFenwickTreeSimple(size int) *FenwickTreeSimple {
	t := &FenwickTreeSimple{
		data: make([]int, size+1),
	}
	return t
}

// Update increases element at index by value
func (t *FenwickTreeSimple) Update(index int, value int) {
	t.update(&t.data, index, value)
}

// Query returns sum on [0, index)
func (t *FenwickTreeSimple) Query(index int) int {
	return t.query(&t.data, index)
}

// QueryRange returns sum in a range
func (t *FenwickTreeSimple) QueryRange(i, j int) int {
	return t.Query(j) - t.Query(i-1)
}

// Get returns value of element at index
func (t *FenwickTreeSimple) Get(index int) int {
	return t.QueryRange(index, index)
}

// Set sets a value for element at index
func (t *FenwickTreeSimple) Set(index int, value int) {
	t.Update(index, value-t.Get(index))
}

func (t *FenwickTreeSimple) update(array *[]int, index int, value int) {
	index++
	for ; index < len(*array); index += index & -index {
		(*array)[index] += value
	}
}

func (t *FenwickTreeSimple) query(array *[]int, index int) int {
	res := 0
	index++
	for ; index > 0; index -= index & -index {
		res += (*array)[index]
	}
	return res
}
