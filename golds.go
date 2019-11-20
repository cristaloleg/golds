package golds

// Container basic interface for all collections
type Container interface {
	// Size returns amount of elements inside a collection
	Size() int

	// IsEmpty returns true if collection is empty
	IsEmpty() bool

	// Clear removes all elements from the collection
	Clear()
}

// CmpFunc returns true if a is greater then b.
type CmpFunc func(a, b interface{}) bool
