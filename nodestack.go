package stack

import "errors"

type Node struct {
	value    string
	adjacent []Node
}

type NodeStack struct {
	contents []Node
}

func (s *NodeStack) pop() (Node, error) {
	l := len(s.contents)
	if l == 0 {
		return Node{}, errors.New("Cannot pop from empty stack")
	}
	top := s.contents[l-1]
	s.contents = s.contents[:l-1]
	return top, nil
}

func (s *NodeStack) push(val Node) {
	s.contents = append(s.contents, val)
	return
}
