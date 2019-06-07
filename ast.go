package json

import (
	"fmt"
)

type (
	// interface for all ast values
	Value interface {
		Value() interface{}
	}

	// Object and Array
	NestValue interface {
		NestValue()
	}

	Object struct {
		inner []*Pair
	}

	Pair struct {
		Key   string
		Value Value
	}

	Array struct {
		inner []Value
	}

	// Values

	Bool struct {
		inner bool
	}

	Int struct {
		inner int64
	}

	Float struct {
		inner float64
	}

	String struct {
		inner string
	}

	Null struct{}
)

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Value() interface{} {
	return o
}

func (o *Object) AddValue(pair *Pair) {
	o.inner = append(o.inner, pair)
}

func NewPair(k string, v Value) *Pair {
	return &Pair{
		Key:   k,
		Value: v,
	}
}

func NewArray() *Array {
	return &Array{}
}

func (a *Array) Value() interface{} {
	return a
}

func (a *Array) AddValue(v Value) {
	a.inner = append(a.inner, v)
}

func (Object) NestValue() {}
func (Array) NestValue()  {}

func NewBool(b bool) *Bool {
	return &Bool{
		inner: b,
	}
}

func (b *Bool) Value() interface{} {
	return b.inner
}

func (b *Bool) String() string {
	if b.inner {
		return "true"
	} else {
		return "false"
	}
}

func NewInt(i int64) *Int {
	return &Int{
		inner: i,
	}
}

func (i *Int) Value() interface{} {
	return i.inner
}

func (i *Int) String() string {
	return fmt.Sprint(i.inner)
}

func NewFloat(f float64) *Float {
	return &Float{
		inner: f,
	}
}

func (f *Float) Value() interface{} {
	return f.inner
}

func (f *Float) String() string {
	return fmt.Sprint(f.inner)
}

func NewString(s string) *String {
	return &String{
		inner: s,
	}
}

func (s *String) Value() interface{} {
	return s.inner
}

func (s *String) String() string {
	return fmt.Sprintf("%q", s.inner)
}

func NewNull() *Null {
	return &Null{}
}

func (n *Null) Value() interface{} {
	return nil
}

func (n *Null) String() string {
	return "null"
}
