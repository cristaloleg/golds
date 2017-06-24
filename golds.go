package golds

const version = "0.0.0"

// Container basic interface for all collections
type Container interface {
	// Size returns amount of elements inside a collection
	Size() int

	// IsEmpty returns true if collection is empty
	IsEmpty() bool

	// Clear removes all elements from the collection
	Clear()
}

// Comparable is an interface for comparable items
// contains only one method Less, which return true
// if an element is less than given
type Comparable interface {
	Less(Comparable) bool
}
