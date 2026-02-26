// Copyright 2021 by Leipzig University Library, http://ub.uni-leipzig.de
//                   The Finc Authors, http://finc.info
//                   Martin Czygan, <martin.czygan@uni-leipzig.de>
//
// This file is part of some open source application.
//
// Some open source application is free software: you can redistribute
// it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either
// version 3 of the License, or (at your option) any later version.
//
// Some open source application is distributed in the hope that it will
// be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//
// @license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>

package zek

// Stack is a simple generic stack.
type Stack[T any] struct {
	v []T
}

// Len returns number of items on the stack.
func (s *Stack[T]) Len() int {
	return len(s.v)
}

// Put item onto stack.
func (s *Stack[T]) Put(item T) {
	s.v = append(s.v, item)
}

// Peek returns the top element without removing it. Panics if stack is empty.
func (s *Stack[T]) Peek() T {
	if len(s.v) == 0 {
		panic("peek at empty stack")
	}
	return s.v[len(s.v)-1]
}

// Pop item from stack. Panics if stack is empty.
func (s *Stack[T]) Pop() T {
	if len(s.v) == 0 {
		panic("pop from empty stack")
	}
	last := len(s.v) - 1
	v := s.v[last]
	s.v = s.v[0:last]
	return v
}
