package set

// BitSet data structure
type BitSet struct {
	data []uint64
}

// NewBitSet returns a pointer to new BitSet
func NewBitSet(size int) *BitSet {
	b := &BitSet{
		data: make([]uint64, size/8+1),
	}
	return b
}

// Clone returns a copy of a BitSet
func (b *BitSet) Clone() *BitSet {
	t := &BitSet{}
	t.data = append(t.data, b.data...)
	return t
}

// Set sets given bit to true
func (b *BitSet) Set(i int) {
	x, y := b.getIndex(i)
	b.data[x] |= b.getMask(y)
}

// SetMany sets given bits to true
func (b *BitSet) SetMany(indexes ...int) {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		b.data[x] |= b.getMask(y)
	}
}

// SetRange sets bits in range [i,j] to true
func (b *BitSet) SetRange(i, j int) {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		b.data[x] |= b.getMask(y)
	}
}

// Unset sets given bit to false
func (b *BitSet) Unset(i int) {
	x, y := b.getIndex(i)
	b.data[x] &^= b.getMask(y)
}

// UnsetMany sets given bits to false
func (b *BitSet) UnsetMany(indexes ...int) {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		b.data[x] &^= b.getMask(y)
	}
}

// UnsetRange sets bits in range [i,j] to false
func (b *BitSet) UnsetRange(i, j int) {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		b.data[x] &^= b.getMask(y)
	}
}

// Get return true if bit is true, false otherwise
func (b *BitSet) Get(i int) bool {
	x, y := b.getIndex(i)
	return (b.data[x] & b.getMask(y)) != 0
}

// GetMany returns bit status for indexes
func (b *BitSet) GetMany(indexes ...int) []bool {
	res := make([]bool, len(indexes))
	for i, idx := range indexes {
		x, y := b.getIndex(idx)
		res[i] = (b.data[x] & b.getMask(y)) != 0
	}
	return res
}

// GetRange returns bits statuses from range [i,j]
func (b *BitSet) GetRange(i, j int) []bool {
	res := make([]bool, j-i+1)
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		res[idx-i] = (b.data[x] & b.getMask(y)) != 0
	}
	return res
}

// Toggle flips bit value
func (b *BitSet) Toggle(i int) {
	x, y := b.getIndex(i)
	b.data[x] ^= b.getMask(y)
}

// ToggleMany flips bits values
func (b *BitSet) ToggleMany(indexes ...int) {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		b.data[x] ^= b.getMask(y)
	}
}

// ToggleRange flips bits in range [i,j] to false
func (b *BitSet) ToggleRange(i, j int) {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		b.data[x] ^= b.getMask(y)
	}
}

// Count returns number of true bits
func (b *BitSet) Count() int {
	res := 0
	for _, v := range b.data {
		res += getBits(v)
	}
	return res
}

// Any returns true if at least 1 bit is true
func (b *BitSet) Any() bool {
	return b.Count() > 0
}

// AnyMany returns true if at least 1 bit from indexes is true
func (b *BitSet) AnyMany(indexes ...int) bool {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return true
		}
	}
	return false
}

// AnyRange returns true if at least 1 bit from range is true
func (b *BitSet) AnyRange(i, j int) bool {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return true
		}
	}
	return false
}

// None returns true if no bits is true
func (b *BitSet) None() bool {
	return b.Count() == 0
}

// NoneMany returns true if no bits from indexes are true
func (b *BitSet) NoneMany(indexes ...int) bool {
	for _, idx := range indexes {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return false
		}
	}
	return true
}

// NoneRange returns true if no bits in range are true
func (b *BitSet) NoneRange(i, j int) bool {
	for idx := i; idx <= j; idx++ {
		x, y := b.getIndex(idx)
		if (b.data[x] & b.getMask(y)) != 0 {
			return false
		}
	}
	return true
}

func (b *BitSet) getIndex(i int) (int, int) {
	return i >> 6, i & 63
}

func (b *BitSet) getMask(i int) uint64 {
	return uint64(1 << uint(i))
}
