package stack

import "errors"

type NodeStack struct {
	s Stack
}

func (s *NodeStack) pop() (Node, error) {
	var auto Node
	top, err := s.s.pop()
	if err != nil {
		return auto, err
	}
	topInstance, correct := top.(Node)
	if !correct {
		return auto, errors.New("Popped element has invalid type. Expected Node.")
	}
	return topInstance, nil
}

func (s *NodeStack) push(val Node) {
	s.s.push(val)
	return
}
