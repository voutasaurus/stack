package stack

import "errors"

type IntStack struct {
	s Stack
}

func (s *IntStack) pop() (int, error) {
	var auto int
	top, err := s.s.pop()
	if err != nil {
		return auto, err
	}
	topInstance, correct := top.(int)
	if !correct {
		return auto, errors.New("Popped element has invalid type. Expected int.")
	}
	return topInstance, nil
}

func (s *IntStack) push(val int) {
	s.s.push(val)
	return
}
