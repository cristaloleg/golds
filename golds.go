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

// IntCmp to compare two int params.
func IntCmp(a, b interface{}) bool {
	return a.(int) < b.(int)
}

// FloatCmp to compare two float params.
func FloatCmp(a, b interface{}) bool {
	return a.(float64) < b.(float64)
}

// StringCmp to compare two string params.
func StringCmp(a, b interface{}) bool {
	return a.(string) < b.(string)
}
