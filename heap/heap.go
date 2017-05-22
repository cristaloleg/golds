package heap

// Heap interface that all heal-like collections implement
type Heap interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Top() (value interface{}, ok bool)
	Build(values []interface{})
}
