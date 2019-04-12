package json

import _stack "github.com/golang-collections/collections/stack"

type nestType int

const (
	_ nestType = iota
	nestObj
	nestArr
)

type nestStack struct {
	inner *_stack.Stack
}

func newStack() *nestStack {
	return &nestStack{
		inner: _stack.New(),
	}
}

func (s *nestStack) push(n nestType) {
	s.inner.Push(n)
}

func (s *nestStack) pop() nestType {
	p := s.inner.Pop()
	if p != nil {
		return p.(nestType)
	}
	return 0
}

func (s *nestStack) peek() nestType {
	p := s.inner.Peek()
	if p != nil {
		return p.(nestType)
	}
	return 0
}

func (s *nestStack) len() int {
	return s.inner.Len()
}
