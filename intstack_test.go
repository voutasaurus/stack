package stack

import (
	"testing"
)

func TestIntStackPush(t *testing.T) {
	s := new(IntStack)
	for i := 0; i < 100; i++ {
		s.push(i)
	}
	for i := 99; i >= 0; i-- {
		v, err := s.pop()
		if err != nil {
			t.Error(err.Error())
			t.Error("Expected", i)
		} else if v != i {
			t.Error("Popped", v, "Expected", i)
		}
	}
}

func TestIntStackEmptyPop(t *testing.T) {
	s := new(IntStack)
	v, err := s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}

	s.push(0)
	v, err = s.pop()
	if err != nil {
		t.Error(err.Error())
		t.Error("Expected", 0)
	}
	v, err = s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}
}
