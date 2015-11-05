package stack

import "errors"

type StringStack struct {
	s Stack
}

func (s *StringStack) pop() (string, error) {
	top, err := s.s.pop()
	if err != nil {
		return "", err
	}
	topInstance, correct := top.(string)
	if !correct {
		return "", errors.New("Popped element has invalid type. Expected string.")
	}
	return topInstance, nil
}

func (s *StringStack) push(val string) {
	s.s.push(val)
	return
}
