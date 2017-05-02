package maps

// BidiMap XXX
type BidiMap struct {
	data map[interface{}]interface{}
	inv  map[interface{}]interface{}
}

// NewBidiMap XXX
func NewBidiMap() *BidiMap {
	m := &BidiMap{
		data: make(map[interface{}]interface{}),
		inv:  make(map[interface{}]interface{}),
	}
	return m
}

// Put XXX
func (m *BidiMap) Put(key interface{}, value interface{}) {
	m.data[key] = value
	m.inv[value] = key
}

// Get XXX
func (m *BidiMap) Get(key interface{}) (value interface{}, ok bool) {
	value, ok = m.data[key]
	return value, ok
}

// GetInv XXX
func (m *BidiMap) GetInv(value interface{}) (key interface{}, ok bool) {
	key, ok = m.inv[value]
	return key, ok
}

// Has XXX
func (m *BidiMap) Has(key interface{}) bool {
	_, ok := m.data[key]
	return ok
}

// HasInv XXX
func (m *BidiMap) HasInv(value interface{}) bool {
	_, ok := m.inv[value]
	return ok
}

// Del XXX
func (m *BidiMap) Del(key interface{}) {
	value, ok := m.data[key]
	if ok {
		delete(m.data, key)
		delete(m.inv, value)
	}
}

// DelInv XXX
func (m *BidiMap) DelInv(value interface{}) {
	key, ok := m.inv[value]
	if ok {
		delete(m.inv, value)
		delete(m.data, key)
	}
}

// Keys XXX
func (m *BidiMap) Keys() []interface{} {
	keys := make([]interface{}, len(m.data))
	i := 0
	for k := range m.data {
		keys[i] = k
		i++
	}
	return keys
}

// Values XXX
func (m *BidiMap) Values() []interface{} {
	values := make([]interface{}, len(m.inv))
	i := 0
	for k := range m.inv {
		values[i] = k
		i++
	}
	return values
}
