package filter

import "github.com/cristaloleg/golds/set"

// BloomFilter XXX
type BloomFilter struct {
	data     set.BitSet
	size     int
	elements int
	hashes   int
}

// NewBloomFilter XXX
func NewBloomFilter() *BloomFilter {
	f := &BloomFilter{}
	return f
}

// Add XXX
func (f *BloomFilter) Add(value interface{}) {
	f.elements++
	a, b := 0, 0
	for i := 0; i < f.hashes; i++ {
		index := (a + b*i) % f.size
		f.data.Set(index)
	}
}

// Has XXX
func (f *BloomFilter) Has(value interface{}) bool {
	a, b := 0, 0
	for i := 0; i < f.hashes; i++ {
		index := (a + b*i) % f.size
		if !f.data.Get(index) {
			return false
		}
	}
	return true
}

// Count XXX
func (f *BloomFilter) Count() int {
	return f.elements
}
