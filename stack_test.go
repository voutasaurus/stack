package stack

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	s := new(Stack)
	for i := 0; i < 100; i++ {
		s.push(i)
	}
	for i := 99; i >= 0; i-- {
		v, err := s.pop()
		if err != nil {
			t.Error(err.Error())
			t.Error("Expected Node{", i, ", []}")
		} else if v.(int) != i {
			t.Error("Popped", v, "Expected", i)
		}
	}
}

func TestStackEmptyPop(t *testing.T) {
	s := new(Stack)
	v, err := s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}

	s.push(0)
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
