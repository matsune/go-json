package json

import (
	"reflect"
	"testing"
)

func TestNewPos(t *testing.T) {
	tests := []struct {
		name string
		want *Pos
	}{
		{
			want: &Pos{
				line: 1,
				col:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPos_String(t *testing.T) {
	type fields struct {
		line int
		col  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				line: 10,
				col:  12,
			},
			want: "line 10 column 12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pos{
				line: tt.fields.line,
				col:  tt.fields.col,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Pos.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
