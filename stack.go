package stack

import (
	"errors"
)

type Stack struct {
	contents []interface{}
}

func (s *Stack) pop() (interface{}, error) {
	l := len(s.contents)
	if l == 0 {
		return nil, errors.New("Cannot pop from empty stack")
	}
	top := s.contents[l-1]
	s.contents = s.contents[:l-1]
	return top, nil
}

func (s *Stack) push(val interface{}) {
	s.contents = append(s.contents, val)
	return
}
