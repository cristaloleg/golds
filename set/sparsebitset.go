package set

// SparseBitSet ...
type SparseBitSet struct {
	data map[int]uint64
}

// NewSparseBitSet returns a pointer to new BitSet
func NewSparseBitSet() *SparseBitSet {
	b := &SparseBitSet{
		data: make(map[int]uint64),
	}
	return b
}

// Set sets given bit to true
func (b *SparseBitSet) Set(i int) {
	x, y := b.getIndex(i)
	value := b.data[x]
	b.data[x] = value | b.getMask(y)
}

// SetMany sets given bits to true
func (b *SparseBitSet) SetMany(indexes ...int) {
	for _, idx := range indexes {
		b.Set(idx)
	}
}

// SetRange sets bits in range [i,j] to true
func (b *SparseBitSet) SetRange(i, j int) {
	for idx := i; idx <= j; idx++ {
		b.Set(idx)
	}
}

// UnsetMany sets given bits to false
func (b *SparseBitSet) UnsetMany(indexes ...int) {
	for _, idx := range indexes {
		b.Unset(idx)
	}
}

// UnsetRange sets bits in range [i,j] to false
func (b *SparseBitSet) UnsetRange(i, j int) {
	for idx := i; idx <= j; idx++ {
		b.Unset(idx)
	}
}

// Unset sets given bit to false
func (b *SparseBitSet) Unset(i int) {
	x, y := b.getIndex(i)
	value, ok := b.data[x]
	if !ok {
		return
	}
	value &^= b.getMask(y)
	if value == 0 {
		delete(b.data, x)
	} else {
		b.data[x] = value
	}
}

// Toggle flips bit value
func (b *SparseBitSet) Toggle(i int) {
	x, y := b.getIndex(i)
	value, ok := b.data[x]
	mask := b.getMask(y)
	if !ok {
		b.data[x] = mask
		return
	}
	value ^= mask
	if value == 0 {
		delete(b.data, x)
	} else {
		b.data[x] = value
	}
}

// ToggleMany flips bits values
func (b *SparseBitSet) ToggleMany(indexes ...int) {
	for _, idx := range indexes {
		b.Toggle(idx)
	}
}

// ToggleRange flips bits in range [i,j] to false
func (b *SparseBitSet) ToggleRange(i, j int) {
	for idx := i; idx <= j; idx++ {
		b.Toggle(idx)
	}
}

// Get return true if bit is true, false otherwise
func (b *SparseBitSet) Get(i int) bool {
	x, y := b.getIndex(i)
	value, ok := b.data[x]
	if !ok {
		return false
	}
	return (value & b.getMask(y)) != 0
}

// GetMany returns bit status for indexes
func (b *SparseBitSet) GetMany(indexes ...int) []bool {
	res := make([]bool, len(indexes))
	for i, idx := range indexes {
		res[i] = b.Get(idx)
	}
	return res
}

// GetRange returns bits statuses from range [i,j]
func (b *SparseBitSet) GetRange(i, j int) []bool {
	res := make([]bool, j-i+1)
	for idx := i; idx <= j; idx++ {
		res[idx-i] = b.Get(idx)
	}
	return res
}

// Count returns number of true bits
func (b *SparseBitSet) Count() int {
	res := 0
	for _, v := range b.data {
		res += getBits(v)
	}
	return res
}

// Any returns true if at least 1 bit is true
func (b *SparseBitSet) Any() bool {
	return b.Count() > 0
}

//AnyMany returns true if at least 1 bit from indexes is true
func (b *SparseBitSet) AnyMany(indexes ...int) bool {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return true
		}
	}
	return false
}

//AnyRange returns true if at least 1 bit from range is true
func (b *SparseBitSet) AnyRange(i, j int) bool {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return true
		}
	}
	return false
}

//NoneMany returns true if no bits from indexes are true
func (b *SparseBitSet) NoneMany(indexes ...int) bool {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return false
		}
	}
	return true
}

//NoneRange returns true if no bits in range are true
func (b *SparseBitSet) NoneRange(i, j int) bool {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return false
		}
	}
	return true
}

// None returns true if no bits is true
func (b *SparseBitSet) None() bool {
	return b.Count() == 0
}

func (b *SparseBitSet) getIndex(i int) (int, int) {
	return i >> 6, i & 63
}

func (b *SparseBitSet) getMask(i int) uint64 {
	return uint64(1 << uint(i))
}
