package set

type SetN struct {
	data *BitSet
}

func NewSetN(size int) *SetN {
	m := &SetN{
		data: NewBitSet(size),
	}
	return m
}

func (s *SetN) Size() int {
	return s.data.Count()
}

func (s *SetN) IsEmpty() bool {
	return s.data.None()
}

func (s *SetN) Put(key int) {
	s.data.Set(key)
}

func (s *SetN) Has(key int) bool {
	return s.data.Get(key)
}

func (s *SetN) Del(key int) {
	s.data.Unset(key)
}

func (s *SetN) Values() []int {
	return nil
}
