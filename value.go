package json

import "fmt"

type (
	Value interface {
		Value() interface{}
	}

	nestValue interface {
		addValue(v Value)
	}

	KeyValue struct {
		Key   string
		Value Value
	}

	BoolValue struct {
		inner bool
	}

	IntValue struct {
		inner int
	}

	FloatValue struct {
		inner float64
	}

	StringValue struct {
		inner string
	}

	NullValue struct{}

	ObjectValue struct {
		KeyValues []*KeyValue
	}

	ArrayValue struct {
		Values []Value
	}
)

func newBool(b bool) *BoolValue {
	return &BoolValue{
		inner: b,
	}
}

func (b *BoolValue) Value() interface{} {
	return b.inner
}

func (b *BoolValue) String() string {
	if b.inner {
		return "true"
	} else {
		return "false"
	}
}

func newInt(i int) *IntValue {
	return &IntValue{
		inner: i,
	}
}

func (i *IntValue) Value() interface{} {
	return i.inner
}

func (i *IntValue) String() string {
	return string(i.inner)
}

func newFloat(f float64) *FloatValue {
	return &FloatValue{
		inner: f,
	}
}

func (f *FloatValue) Value() interface{} {
	return f.inner
}

func (f *FloatValue) String() string {
	return fmt.Sprint(f.inner)
}

func newString(s string) *StringValue {
	return &StringValue{
		inner: s,
	}
}

func (s *StringValue) Value() interface{} {
	return s.inner
}

func (s *StringValue) String() string {
	return fmt.Sprintf("%q", s.inner)
}

func newNull() *NullValue {
	return &NullValue{}
}

func (n *NullValue) Value() interface{} {
	return nil
}

func (n *NullValue) String() string {
	return "null"
}

func newObject() *ObjectValue {
	return &ObjectValue{}
}

func (o *ObjectValue) Value() interface{} {
	return o
}

func (o *ObjectValue) addValue(v Value) {
	o.SetValue(v)
}

func (o *ObjectValue) AddKey(k string) {
	o.KeyValues = append(o.KeyValues, &KeyValue{
		Key: k,
	})
}

func (o *ObjectValue) SetValue(v Value) {
	o.KeyValues[len(o.KeyValues)-1].Value = v
}

func newArray() *ArrayValue {
	return &ArrayValue{}
}

func (a *ArrayValue) Value() interface{} {
	return a
}

func (a *ArrayValue) push(v Value) {
	a.Values = append(a.Values, v)
}

func (a *ArrayValue) addValue(v Value) {
	a.push(v)
}
