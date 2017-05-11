package set

// SparseBitSet XXX
type SparseBitSet struct {
	data map[int]uint64
}

// NewSparseBitSet returns a pointer to new BitSet
func NewSparseBitSet(size int) *SparseBitSet {
	b := &SparseBitSet{
		data: make(map[int]uint64),
	}
	return b
}

// Set sets given bit to true
func (b *SparseBitSet) Set(i int) {
	x, y := b.getIndex(i)
	value, ok := b.data[x]
	if ok {
		value |= b.getMask(y)
	} else {
		value = b.getMask(y)
	}
	b.data[x] = value
}

// Unset sets given bit to false
func (b *SparseBitSet) Unset(i int) {
	x, y := b.getIndex(i)
	value, ok := b.data[x]
	if !ok {
		return
	}
	value &= b.getMaskInv(y)
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
	if ok {
		value ^= b.getMask(y)
	} else {
		value = b.getMask(y)
	}
	value ^= b.getMaskInv(y)
	if value == 0 {
		delete(b.data, x)
	} else {
		b.data[x] = value
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

// Count returns number of true bits
func (b *SparseBitSet) Count() int {
	res := 0
	for _, v := range b.data {
		res += b.getBits(v)
	}
	return res
}

// Any returns true if at least 1 bit is true
func (b *SparseBitSet) Any() bool {
	return b.Count() > 0
}

// None returns true if no bits is true
func (b *SparseBitSet) None() bool {
	return b.Count() == 0
}
func (b *SparseBitSet) getIndex(i int) (int, int) {
	return i >> 3, i & 7
}

func (b *SparseBitSet) getMask(i int) uint64 {
	return uint64(1 << uint(i))
}

func (b *SparseBitSet) getMaskInv(i int) uint64 {
	return ^uint64(1 << uint(i))
}

func (b *SparseBitSet) getBits(i uint64) int {
	res := 0
	for ; i != 0; res++ {
		i &= i - 1
	}
	return res
}
