package stack

// ArrayStack XXX
type ArrayStack struct {
	data []interface{}
}

// NewArrayStack XXX
func NewArrayStack() *ArrayStack {
	s := &ArrayStack{
		data: make([]interface{}, 0),
	}
	return s
}

// Size XXX
func (s *ArrayStack) Size() int {
	return len(s.data)
}

// IsEmpty XXX
func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear XXX
func (s *ArrayStack) Clear() {
	s.data = make([]interface{}, 0)
}

// Push XXX
func (s *ArrayStack) Push(value interface{}) {
	s.data = append(s.data, value)
}

// Pop XXX
func (s *ArrayStack) Pop() (value interface{}, ok bool) {
	size := len(s.data)
	if size == 0 {
		return nil, false
	}
	value = s.data[size-1]
	s.data = s.data[:size-1]
	return value, true
}

// Top XXX
func (s *ArrayStack) Top() (value interface{}, ok bool) {
	size := len(s.data)
	if size == 0 {
		return nil, false
	}
	return s.data[size-1], true
}

// Values XXX
func (s *ArrayStack) Values() []interface{} {
	return s.data[:]
}
