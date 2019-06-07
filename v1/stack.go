package json

import _stack "github.com/golang-collections/collections/stack"

type Stack struct {
	inner *_stack.Stack
}

func NewStack() *Stack {
	return &Stack{
		inner: _stack.New(),
	}
}

func (s *Stack) pushArray() *Array {
	arr := NewArray()
	s.inner.Push(arr)
	return arr
}

func (s *Stack) pushObject() *Object {
	obj := NewObject()
	s.inner.Push(obj)
	return obj
}

func (s *Stack) pop() NestValue {
	return s.inner.Pop().(NestValue)
}

func (s *Stack) peek() NestValue {
	if v, ok := s.inner.Peek().(NestValue); ok {
		return v
	}
	return nil
}

func (s *Stack) isObj() bool {
	_, ok := s.peek().(*Object)
	return ok
}

func (s *Stack) isArr() bool {
	_, ok := s.peek().(*Array)
	return ok
}

func (s *Stack) len() int {
	return s.inner.Len()
}
