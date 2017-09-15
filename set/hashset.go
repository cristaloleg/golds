package set

// HashSet allows to store values
type HashSet struct {
	data map[interface{}]struct{}
}

// NewHashSet returns a pointer to the HashSet
func NewHashSet() *HashSet {
	s := &HashSet{
		data: make(map[interface{}]struct{}),
	}
	return s
}

// Size return amount of keys in HashMap
func (h *HashSet) Size() int {
	return len(h.data)
}

// IsEmpty returns true if set is empty
func (h *HashSet) IsEmpty() bool {
	return len(h.data) == 0
}

// Clear removes all elements from the set
func (h *HashSet) Clear() {
	h.data = make(map[interface{}]struct{})
}

// Put put value in a HashMap
func (h *HashSet) Put(value interface{}) {
	h.data[value] = struct{}{}
}

// Has return true if given key is inside a HashSet
func (h *HashSet) Has(value interface{}) bool {
	_, ok := h.data[value]
	return ok
}

// Del removes value from a HashSet
func (h *HashSet) Del(value interface{}) {
	delete(h.data, value)
}

// Values returns values stored in HashSet
func (h *HashSet) Values() []interface{} {
	res := make([]interface{}, len(h.data))
	i := 0
	for v := range h.data {
		res[i] = v
		i++
	}
	return res
}
