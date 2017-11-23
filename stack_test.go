package zek

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	if s.Len() != 0 {
		t.Errorf("empty stack contains %d items", s.Len())
	}
	s.Put(1)
	s.Put(2)
	s.Put(3)
	if s.Len() != 3 {
		t.Errorf("empty stack contains %d items", s.Len())
	}
	if v := s.Pop(); v != 3 {
		t.Errorf("pop want 3, got %v", v)
	}
	s.Put(4)
	if v := s.Pop(); v != 4 {
		t.Errorf("pop want 4, got %v", v)
	}
	if v := s.Pop(); v != 2 {
		t.Errorf("pop want 2, got %v", v)
	}
	if v := s.Pop(); v != 1 {
		t.Errorf("pop want 1, got %v", v)
	}
	if s.Len() != 0 {
		t.Errorf("empty stack contains %d items", s.Len())
	}
}
