package stack

import "errors"

type Float64Stack struct {
	s Stack
}

func (s *Float64Stack) pop() (float64, error) {
	var auto float64
	top, err := s.s.pop()
	if err != nil {
		return auto, err
	}
	topInstance, correct := top.(float64)
	if !correct {
		return auto, errors.New("Popped element has invalid type. Expected float64.")
	}
	return topInstance, nil
}

func (s *Float64Stack) push(val float64) {
	s.s.push(val)
	return
}
