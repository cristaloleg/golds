package stack

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Top() (value interface{}, ok bool)
	Values() []interface{}
}
