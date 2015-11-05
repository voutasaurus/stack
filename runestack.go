package stack

import "errors"

type RuneStack struct {
	s Stack
}

func (s *RuneStack) pop() (rune, error) {
	var auto rune
	top, err := s.s.pop()
	if err != nil {
		return auto, err
	}
	topInstance, correct := top.(rune)
	if !correct {
		return auto, errors.New("Popped element has invalid type. Expected rune.")
	}
	return topInstance, nil
}

func (s *RuneStack) push(val rune) {
	s.s.push(val)
	return
}
