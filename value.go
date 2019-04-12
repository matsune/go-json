package json

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
		keyValues []*KeyValue
	}

	ArrayValue struct {
		inner []Value
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

func newInt(i int) *IntValue {
	return &IntValue{
		inner: i,
	}
}

func (i *IntValue) Value() interface{} {
	return i.inner
}

func newFloat(f float64) *FloatValue {
	return &FloatValue{
		inner: f,
	}
}

func (f *FloatValue) Value() interface{} {
	return f.inner
}

func newString(s string) *StringValue {
	return &StringValue{
		inner: s,
	}
}

func (s *StringValue) Value() interface{} {
	return s.inner
}

func newNull() *NullValue {
	return &NullValue{}
}

func (n *NullValue) Value() interface{} {
	return nil
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
	o.keyValues = append(o.keyValues, &KeyValue{
		Key: k,
	})
}

func (o *ObjectValue) SetValue(v Value) {
	o.keyValues[len(o.keyValues)-1].Value = v
}

func newArray() *ArrayValue {
	return &ArrayValue{}
}

func (a *ArrayValue) Value() interface{} {
	return a
}

func (a *ArrayValue) push(v Value) {
	a.inner = append(a.inner, v)
}

func (a *ArrayValue) addValue(v Value) {
	a.push(v)
}
