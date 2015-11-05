package stack

import (
	"testing"
)

func TestFloat64StackPush(t *testing.T) {
	s := new(Float64Stack)
	for i := 0; i < 100; i++ {
		s.push(float64(i))
	}
	for i := 99; i >= 0; i-- {
		v, err := s.pop()
		if err != nil {
			t.Error(err.Error())
			t.Error("Expected", i)
		} else if v != float64(i) {
			t.Error("Popped", v, "Expected", float64(i))
		}
	}
}

func TestFloat64StackEmptyPop(t *testing.T) {
	s := new(Float64Stack)
	v, err := s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}

	s.push(1.0)
	v, err = s.pop()
	if err != nil {
		t.Error(err.Error())
		t.Error("Expected 1.0")
	}
	v, err = s.pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, instead got", v)
	}
}
