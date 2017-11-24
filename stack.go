package zek

import "sync"

// Stack is a simple stack for arbitrary types.
type Stack struct {
	sync.Mutex
	v []interface{}
}

// Len returns number of items on the stack.
func (s *Stack) Len() int {
	s.Lock()
	defer s.Unlock()
	return len(s.v)
}

// Push item onto stack.
func (s *Stack) Put(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.v = append(s.v, item)
}

func (s *Stack) Peek() interface{} {
	s.Lock()
	defer s.Unlock()
	if len(s.v) == 0 {
		panic("peek at empty stack")
	}
	return s.v[len(s.v)-1]
}

// Pop item from stack. It's a panic if stack is empty.
func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	if len(s.v) == 0 {
		panic("pop from empty stack")
	}
	last := len(s.v) - 1
	v := s.v[last]
	s.v = s.v[0:last]
	return v
}
