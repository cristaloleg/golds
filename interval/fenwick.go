package interval

// FenwickTree represents Fenwick tree, aka BIT - Binary indexed tree
type FenwickTree struct {
	data []int
	inv  []int
}

// NewFenwickTree instantiates a new FenwickTree
func NewFenwickTree(size int) *FenwickTree {
	t := &FenwickTree{
		data: make([]int, size+1),
		inv:  make([]int, size+1),
	}
	return t
}

// Update increases element at index by value
func (t *FenwickTree) Update(index int, value int) {
	t.UpdateRange(index, index, value)
}

// UpdateRange increases all elements in a range by value
func (t *FenwickTree) UpdateRange(i, j int, value int) {
	t.update(&t.data, i, value)
	t.update(&t.data, j+1, -value)
	t.update(&t.inv, i, value*(i-1))
	t.update(&t.inv, j+1, -value*j)
}

// Query returns sum on [0, index)
func (t *FenwickTree) Query(index int) int {
	return t.query(&t.data, index)*index - t.query(&t.inv, index)
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

func (t *FenwickTree) update(array *[]int, index int, value int) {
	index++
	for ; index < len(*array); index += index & -index {
		(*array)[index] += value
	}
}

func (t *FenwickTree) query(array *[]int, index int) int {
	res := 0
	index++
	for ; index > 0; index -= index & -index {
		res += (*array)[index]
	}
	return res
}
