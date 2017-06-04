package maps

// MapN maps [0,n) to interface{}
type MapN struct {
	keys  []int
	items []mapNEntry
}

type mapNEntry struct {
	id    int
	value interface{}
}

func NewMapN(size int) *MapN {
	m := &MapN{
		keys: make([]int, size),
	}
	return m
}

func (m *MapN) Size() int {
	return len(m.items)
}

func (m *MapN) IsEmpty() bool {
	return len(m.items) == 0
}

func (m *MapN) Clear() {
	m.items = nil
}

func (m *MapN) Put(key int, value interface{}) {
	i, size := m.keys[key], len(m.items)
	if i < size {
		m.items[i].id = i
		m.items[i].value = value
		return
	}
	m.items = append(m.items, mapNEntry{i, value})
	m.keys[key] = size
}

func (m *MapN) Get(key int) (value interface{}, ok bool) {
	i, size := m.keys[key], len(m.items)
	if i < size && m.items[i].value != nil {
		return m.items[i].value, true
	}
	return nil, false
}

func (m *MapN) Has(key int) bool {
	i, size := m.keys[key], len(m.items)
	return i < size && m.items[i].value != nil
}

func (m *MapN) Del(key int) {
	i, size := m.keys[key], len(m.items)
	if i < size && m.items[i].value != nil {
		tmp := m.items[size-1]
		m.items[i] = tmp
		m.keys[tmp.id] = i
		m.items = m.items[:size-1]
	}
}

func (m *MapN) Keys() []int {
	return m.keys[:]
}

func (m *MapN) Values() []interface{} {
	res := make([]interface{}, len(m.items))
	for i, k := range m.items {
		res[i] = k
	}
	return res
}
