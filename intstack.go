package stack

import (
	"errors"
)

type IntStack struct {
	contents []int
}

func (s *IntStack) pop() (int, error) {
	l := len(s.contents)
	if l == 0 {
		return 0, errors.New("Cannot pop from empty stack")
	}
	top := s.contents[l-1]
	s.contents = s.contents[:l-1]
	return top, nil
}

func (s *IntStack) push(val int) {
	s.contents = append(s.contents, val)
	return
}
