package stack

import (
	"strconv"
	"testing"
)

func TestStringStackPush(t *testing.T) {
	s := new(StringStack)
	for i := 0; i < 100; i++ {
		s.push(strconv.Itoa(i))
	}
	for i := 99; i >= 0; i-- {
		v, err := s.pop()
		if err != nil {
			t.Error(err.Error())
			t.Error("Expected Node{", i, ", []}")
		} else if v != strconv.Itoa(i) {
			t.Error("Popped", v, "Expected", i)
		}
	}
}

func TestStringStackEmptyPop(t *testing.T) {
	s := new(StringStack)
	v, err := s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}

	s.push("0")
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
