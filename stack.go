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

// Put item onto stack.
func (s *Stack) Put(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.v = append(s.v, item)
}

// Peek returns the top element without removing it. Panic it stack is empty.
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
