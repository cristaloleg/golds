package set

// Set interface that all set-like structures implement
type Set interface {
	Put(value interface{})
	Has(value interface{}) bool
	Del(value interface{})
	Values() []interface{}
}
