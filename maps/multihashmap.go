package maps

// MultiHashMap XXX
type MultiHashMap struct {
	data map[interface{}][]interface{}
}

// Put XXX
func (m *MultiHashMap) Put(key interface{}, value interface{}) {
	values, ok := m.data[key]
	if ok {
		values = append(values, value)
	} else {
		values = []interface{}{value}
	}
	m.data[key] = values
}

// Get XXX
func (m *MultiHashMap) Get(key interface{}) (values []interface{}, ok bool) {
	values, ok = m.data[key]
	if !ok {
		return nil, false
	}
	return values, true
}

// Count XXX
func (m *MultiHashMap) Count(key interface{}) (count int, ok bool) {
	values, ok := m.data[key]
	if !ok {
		return 0, false
	}
	return len(values), true
}

// Has XXX
func (m *MultiHashMap) Has(key interface{}) bool {
	_, ok := m.data[key]
	return ok
}

// HasKeyValue XXX
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

// Del XXX
func (m *MultiHashMap) Del(key interface{}) {
	delete(m.data, key)
}

// DelKeyValue XXX
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

// Keys XXX
func (m *MultiHashMap) Keys() []interface{} {
	res := make([]interface{}, len(m.data))
	i := 0
	for k := range m.data {
		res[i] = k
		i++
	}
	return res
}
