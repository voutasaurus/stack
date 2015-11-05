package stack

import (
	"strconv"
	"testing"
)

func TestNodeStackPush(t *testing.T) {
	s := new(NodeStack)
	for i := 0; i < 100; i++ {
		s.push(Node{strconv.Itoa(i), []Node{}})
	}
	for i := 99; i >= 0; i-- {
		v, err := s.pop()
		if err != nil {
			t.Error(err.Error())
			t.Error("Expected Node{", i, ", []}")
		} else if v.value != strconv.Itoa(i) {
			t.Error("Popped", v.value, "Expected", i)
		}
	}
}

func TestNodeStackEmptyPop(t *testing.T) {
	s := new(NodeStack)
	v, err := s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}

	s.push(Node{"0", []Node{}})
	v, err = s.pop()
	if err != nil {
		t.Error(err.Error())
		t.Error("Expected Node{0, []}")
	}
	v, err = s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}
}
