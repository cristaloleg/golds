package set

// Set interface that all set-like structures implement
type Set interface {
	Put(value interface{})
	Has(value interface{}) bool
	Del(value interface{})
	Values() []interface{}

	// Iter() <-chan interface{}
	// Union(s Set)
	// Intersect(s Set)
	// Subtract(s Set)
	// IsSubset(s Set) bool
	// IsSuperset(s Set) bool
	// IsEqual(s Set) bool
	// RemoveIf(f func(interface{}) bool)
}

// Union returns set that contains all elements of a and b
func Union(a, b Set) Set {
	res := NewHashSet()
	for _, value := range a.Values() {
		res.Put(value)
	}
	for _, value := range b.Values() {
		res.Put(value)
	}
	return res
}

// Intersection returns all elements from a and b
func Intersection(a, b Set) Set {
	res := NewHashSet()
	for _, value := range a.Values() {
		if b.Has(value) {
			res.Put(value)
		}
	}
	return res
}

// Difference returns all elements from a but not b
func Difference(a, b Set) Set {
	res := NewHashSet()
	for _, value := range a.Values() {
		if !b.Has(value) {
			res.Put(value)
		}
	}
	return res
}
