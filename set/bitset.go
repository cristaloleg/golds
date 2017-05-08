package set

type BitSet struct {
	data []uint64
}

func NewBitSet(size int) *BitSet {
	b := &BitSet{
		data: make([]uint64, size/8),
	}
	return b
}

func (b *BitSet) Set(i int) {
	x, y := b.getIndex(i)
	b.data[x] |= b.getMask(y)
	// f.bits[i>>shift] |= 1 << (i & (word - 1))
}

func (b *BitSet) Unset(i int) {
	x, y := b.getIndex(i)
	b.data[x] &= b.getMaskInv(y)
}

func (b *BitSet) Get(i int) bool {
	x, y := b.getIndex(i)
	return (b.data[x] & b.getMask(y)) != 0
	//return ((f.bits[i>>shift] >> (i & (word - 1))) & 1) != 0
}

func (b *BitSet) Toggle(i int) {
	x, y := b.getIndex(i)
	b.data[x] ^= b.getMask(y)
}

func (b *BitSet) Count() int {
	res := 0
	for _, v := range b.data {
		res += b.getBits(v)
	}
	return res
}

func (b *BitSet) getIndex(i int) (int, int) {
	return i / 8, i % 8
}

func (b *BitSet) getMask(i int) uint64 {
	return uint64(1 << uint(i))
}

func (b *BitSet) getMaskInv(i int) uint64 {
	return ^uint64(1 << uint(i))
}

func (b *BitSet) getBits(i uint64) int {
	res := 0
	for ; i != 0; res++ {
		i &= i - 1
	}
	return res
}
