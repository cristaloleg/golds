package maps

// HashMap allows to store key-value pairs
type HashMap struct {
	data map[interface{}]interface{}
}

// NewHashMap returns a pointer to the HashMap
func NewHashMap() *HashMap {
	h := &HashMap{
		make(map[interface{}]interface{}),
	}
	return h
}

// Size return amount of keys in HashMap
func (h *HashMap) Size() int {
	return len(h.data)
}

// IsEmpty returns true if HashMap is empty
func (h *HashMap) IsEmpty() bool {
	return len(h.data) == 0
}

// Clear removes all elements from the HashMap
func (h *HashMap) Clear() {
	h.data = make(map[interface{}]interface{})
}

// Put put key-value pair in HashMap
func (h *HashMap) Put(key interface{}, value interface{}) {
	h.data[key] = value
}

// Get returns a value by given key
func (h *HashMap) Get(key interface{}) (value interface{}, ok bool) {
	value, ok = h.data[key]
	return value, ok
}

// Has return true if given key is inside a HashMap
func (h *HashMap) Has(key interface{}) bool {
	_, ok := h.data[key]
	return ok
}

// Del removes key-value by given key from a HashMap
func (h *HashMap) Del(key interface{}) {
	delete(h.data, key)
}

// Keys returns keys stored in HashMap
func (h *HashMap) Keys() []interface{} {
	res := make([]interface{}, len(h.data))
	i := 0
	for k := range h.data {
		res[i] = k
		i++
	}
	return res
}
