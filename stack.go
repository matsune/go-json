package json

import _stack "github.com/golang-collections/collections/stack"

type nestStack struct {
	inner *_stack.Stack
}

func newStack() *nestStack {
	return &nestStack{
		inner: _stack.New(),
	}
}

func (s *nestStack) pushArray() *ArrayValue {
	arr := newArray()
	s.inner.Push(arr)
	return arr
}

func (s *nestStack) pushObject() *ObjectValue {
	obj := newObject()
	s.inner.Push(obj)
	return obj
}

func (s *nestStack) pop() nestValue {
	return s.inner.Pop().(nestValue)
}

func (s *nestStack) peek() nestValue {
	if v, ok := s.inner.Peek().(nestValue); ok {
		return v
	}
	return nil
}

func (s *nestStack) isObj() bool {
	switch s.peek().(type) {
	case *ObjectValue:
		return true
	default:
		return false
	}
}

func (s *nestStack) isArr() bool {
	switch s.peek().(type) {
	case *ArrayValue:
		return true
	default:
		return false
	}
}

func (s *nestStack) len() int {
	return s.inner.Len()
}
