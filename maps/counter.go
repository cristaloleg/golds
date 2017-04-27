package maps

// Counter allows to easily track number of elements
type Counter struct {
	data map[interface{}]int64
}

// NewCounter returns a pointer to the Counter
func NewCounter() *Counter {
	h := &Counter{
		make(map[interface{}]int64),
	}
	return h
}

// Put adds a key with count occurrences
func (h *Counter) Put(key interface{}, count int64) {
	h.data[key] = count
}

// Inc increase key occurrences by 1
func (h *Counter) Inc(key interface{}) {
	value, ok := h.data[key]
	if !ok {
		value = 0
	}
	h.data[key] = value + 1
}

// Dec decrease key occurrences by 1
func (h *Counter) Dec(key interface{}) {
	value, ok := h.data[key]
	if !ok || value == 0 {
		value = 1
	}
	h.data[key] = value - 1
}

// Get returns number of occurrences for the key
func (h *Counter) Get(key interface{}) int64 {
	value, ok := h.data[key]
	if ok {
		return value
	}
	return -1
}

// Has returns true if key is presented
func (h *Counter) Has(key interface{}) bool {
	_, ok := h.data[key]
	return ok
}

// Del removes a key from the collection
func (h *Counter) Del(key interface{}) {
	delete(h.data, key)
}

// Keys returns keys stored in Counter
func (h *Counter) Keys() []interface{} {
	res := make([]interface{}, len(h.data))
	i := 0
	for k := range h.data {
		res[i] = k
		i++
	}
	return res
}
