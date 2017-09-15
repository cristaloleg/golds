package list

// ArrayStack ...
type ArrayStack struct {
	data []interface{}
}

// NewArrayStack returns a pointer to the ArrayStack
func NewArrayStack() *ArrayStack {
	s := &ArrayStack{
		data: make([]interface{}, 0),
	}
	return s
}

// Size return amount of keys in the stack
func (s *ArrayStack) Size() int {
	return len(s.data)
}

// IsEmpty returns true if stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the stack
func (s *ArrayStack) Clear() {
	s.data = make([]interface{}, 0)
}

// Push adds element to the top of the stack
func (s *ArrayStack) Push(value interface{}) {
	s.data = append(s.data, value)
}

// PushMany adds elements to the top of the stack
func (s *ArrayStack) PushMany(values ...interface{}) {
	for _, v := range values {
		s.Push(v)
	}
}

// Pop removes and returns top element of the stack
func (s *ArrayStack) Pop() (value interface{}, ok bool) {
	size := len(s.data)
	if size == 0 {
		return nil, false
	}
	value = s.data[size-1]
	s.data = s.data[:size-1]
	return value, true
}

// PopMany removes and returns top element of the stack
func (s *ArrayStack) PopMany(k int) (values []interface{}, ok bool) {
	if s.Size() == 0 {
		return nil, false
	}
	k = min(k, s.Size())
	values = make([]interface{}, k)
	for i := 0; i < k; i++ {
		value, _ := s.Pop()
		values[i] = value
	}
	return values, true
}

// Top returns top element of the stack
func (s *ArrayStack) Top() (value interface{}, ok bool) {
	size := len(s.data)
	if size == 0 {
		return nil, false
	}
	return s.data[size-1], true
}

// Values returns values presented in stack
func (s *ArrayStack) Values() []interface{} {
	return s.data[:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
