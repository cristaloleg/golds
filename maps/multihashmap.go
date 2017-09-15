package maps

// MultiHashMap hashmap with multiple values per key
type MultiHashMap struct {
	data map[interface{}][]interface{}
}

// NewMultiHashMap returns a pointer to the MultiHashMap
func NewMultiHashMap() *MultiHashMap {
	m := &MultiHashMap{
		data: make(map[interface{}][]interface{}),
	}
	return m
}

// Put ...
func (m *MultiHashMap) Put(key interface{}, value interface{}) {
	values, ok := m.data[key]
	if ok {
		values = append(values, value)
	} else {
		values = []interface{}{value}
	}
	m.data[key] = values
}

// Get ...
func (m *MultiHashMap) Get(key interface{}) (values []interface{}, ok bool) {
	values, ok = m.data[key]
	if ok {
		return values, true
	}
	return nil, false
}

// Count ...
func (m *MultiHashMap) Count(key interface{}) (count int, ok bool) {
	values, ok := m.data[key]
	if ok {
		return len(values), true
	}
	return 0, false
}

// Has ...
func (m *MultiHashMap) Has(key interface{}) bool {
	_, ok := m.data[key]
	return ok
}

// HasKeyValue ...
func (m *MultiHashMap) HasKeyValue(key interface{}, value interface{}) bool {
	values, ok := m.data[key]
	if !ok {
		return false
	}
	for val := range values {
		if val == value {
			return true
		}
	}
	return false
}

// Del ...
func (m *MultiHashMap) Del(key interface{}) {
	delete(m.data, key)
}

// DelKeyValue ...
func (m *MultiHashMap) DelKeyValue(key interface{}, value interface{}) {
	values, ok := m.data[key]
	if !ok {
		return
	}
	for i, val := range values {
		if val == value {
			values = append(values[:i], values[i+1:]...)
		}
	}
	m.data[key] = values
}

// Keys ...
func (m *MultiHashMap) Keys() []interface{} {
	res := make([]interface{}, len(m.data))
	i := 0
	for k := range m.data {
		res[i] = k
		i++
	}
	return res
}
