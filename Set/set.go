package Set

import (
	"errors"
)

type set map[interface{}]bool

func NewSet() *set {
	s := make(set)
	return &s
}

func (s *set) Len() int {
	return len((*s))
}

func (s *set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *set) Find(item interface{}) bool {
	_, ok := (*s)[item]
	return ok
}

func (s *set) Push(item interface{}) error {
	if _, ok := (*s)[item]; ok {
		return errors.New("It's already in the set.")
	}
	(*s)[item] = true
	return nil
}

func (s *set) Pop(item interface{}) error {
	if _, ok := (*s)[item]; !ok {
		return errors.New("This item is not in the set.")
	}
	delete((*s), item)
	return nil
}