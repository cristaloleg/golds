package maps

// Map interface that all key-value structures implement
type Map interface {
	Put(key interface{}, value interface{})
	Get(ket interface{}) (value interface{}, ok bool)
	Has(key interface{}) bool
	Del(key interface{})
	Keys() []interface{}
}
