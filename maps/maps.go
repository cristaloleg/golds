package maps

// Map interface that all key-value structures implement
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, ok bool)
	Has(key interface{}) bool
	Del(key interface{})
	Keys() []interface{}
}

// OrderedMap interface that all ordered key-value
type OrderedMap interface {
	Succ(interface{}) interface{}
	Prec(interface{}) interface{}
	Min() interface{}
	Max() interface{}
}
