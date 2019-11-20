package golds

// BIT represents Binary indexed tree (Fenwick tree).
//
type BIT struct {
	data []int
}

// NewBIT instantiates a new BIT.
//
func NewBIT(size int) *BIT {
	t := &BIT{
		data: make([]int, size+1),
	}
	return t
}

// Update increases element at index by value.
//
func (t *BIT) Update(index int, value int) {
	t.update(index, value)
}

// Query returns sum on [0, index)
//
func (t *BIT) Query(index int) int {
	return t.query(index)
}

// QueryRange returns sum in a range.
//
func (t *BIT) QueryRange(i, j int) int {
	return t.Query(j) - t.Query(i-1)
}

// Get returns value of element at index.
//
func (t *BIT) Get(index int) int {
	return t.QueryRange(index, index)
}

// Set sets a value for element at index.
//
func (t *BIT) Set(index int, value int) {
	t.Update(index, value-t.Get(index))
}

func (t *BIT) update(index int, value int) {
	index++
	for ; index < len(t.data); index += index & -index {
		t.data[index] += value
	}
}

func (t *BIT) query(index int) int {
	index++
	res := 0
	for ; index > 0; index -= index & -index {
		res += t.data[index]
	}
	return res
}
