package stack

import (
	"testing"
)

func TestRuneStackPush(t *testing.T) {
	s := new(RuneStack)
	for i := 0; i < 100; i++ {
		s.push(rune(i))
	}
	for i := 99; i >= 0; i-- {
		v, err := s.pop()
		if err != nil {
			t.Error(err.Error())
			t.Error("Expected", i)
		} else if v != rune(i) {
			t.Error("Popped", v, "Expected", string(rune(i)))
		}
	}
}

func TestRuneStackEmptyPop(t *testing.T) {
	s := new(RuneStack)
	v, err := s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}

	s.push('a')
	v, err = s.pop()
	if err != nil {
		t.Error(err.Error())
		t.Error("Expected a")
	}
	v, err = s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}
}
