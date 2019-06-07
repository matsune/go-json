package json

import (
	"reflect"
	"testing"

	_stack "github.com/golang-collections/collections/stack"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_pushArray(t *testing.T) {
	type fields struct {
		inner *_stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   *Array
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				inner: tt.fields.inner,
			}
			if got := s.pushArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.pushArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_pushObject(t *testing.T) {
	type fields struct {
		inner *_stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   *Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				inner: tt.fields.inner,
			}
			if got := s.pushObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.pushObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_pop(t *testing.T) {
	type fields struct {
		inner *_stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   NestValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				inner: tt.fields.inner,
			}
			if got := s.pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_peek(t *testing.T) {
	type fields struct {
		inner *_stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   NestValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				inner: tt.fields.inner,
			}
			if got := s.peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_isObj(t *testing.T) {
	var s *Stack
	tests := []struct {
		name string
		f    func()
		want bool
	}{
		{
			f: func() {
				s.pushArray()
			},
			want: false,
		},
		{
			f: func() {
				s.pushObject()
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = NewStack()
			tt.f()
			if got := s.isObj(); got != tt.want {
				t.Errorf("Stack.isObj() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_isArr(t *testing.T) {
	var s *Stack
	tests := []struct {
		name string
		f    func()
		want bool
	}{
		{
			f: func() {
				s.pushArray()
			},
			want: true,
		},
		{
			f: func() {
				s.pushObject()
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = NewStack()
			tt.f()
			if got := s.isArr(); got != tt.want {
				t.Errorf("Stack.isArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_len(t *testing.T) {
	type fields struct {
		inner *_stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				inner: tt.fields.inner,
			}
			if got := s.len(); got != tt.want {
				t.Errorf("Stack.len() = %v, want %v", got, tt.want)
			}
		})
	}
}
