package stack

import "errors"

type stack []interface{}

func NewStack() *stack {
	s := make(stack, 0)
	return &s
}

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return (*s)[s.Len()-1], nil
}

func (s *stack) Push(item interface{}) {
	(*s) = append((*s), item)
}

func (s *stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	top := (*s)[s.Len() - 1]
	(*s) = (*s)[:s.Len() - 1]
	return top, nil
}