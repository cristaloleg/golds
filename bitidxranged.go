package golds

// BITRanged represents Binary indexed tree (Fenwick tree) over ranges.
// See https://petr-mitrichev.blogspot.com/2013/05/fenwick-tree-range-updates.html
//
type BITRanged struct {
	data []int
	inv  []int
}

// NewBITRanged instantiates a new BITRanged
func NewBITRanged(size int) *BITRanged {
	t := &BITRanged{
		data: make([]int, size+1),
		inv:  make([]int, size+1),
	}
	return t
}

// Get returns value of element at index
func (t *BITRanged) Get(index int) int {
	return t.QueryRange(index, index)
}

// Set sets a value for element at index
func (t *BITRanged) Set(index int, value int) {
	t.Update(index, value-t.Get(index))
}

// Update increases element at index by value
func (t *BITRanged) Update(index int, value int) {
	t.UpdateRange(index, index, value)
}

// UpdateRange increases all elements in a range by value
func (t *BITRanged) UpdateRange(i, j int, value int) {
	t.update(&t.data, i, value)
	t.update(&t.data, j+1, -value)
	t.update(&t.inv, i, value*(i-1))
	t.update(&t.inv, j+1, -value*j)
}

// Query returns sum on [0, index)
func (t *BITRanged) Query(index int) int {
	return t.query(&t.data, index)*index - t.query(&t.inv, index)
}

// QueryRange returns sum in a range
func (t *BITRanged) QueryRange(i, j int) int {
	return t.Query(j) - t.Query(i-1)
}

func (t *BITRanged) update(array *[]int, index int, value int) {
	index++
	a := *array
	for ; index < len(a); index += index & -index {
		a[index] += value
	}
}

func (t *BITRanged) query(array *[]int, index int) int {
	res := 0
	index++
	a := *array
	for ; index > 0; index -= index & -index {
		res += a[index]
	}
	return res
}
