package json

import (
	"reflect"
	"testing"
)

func TestNewScanner(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want *Scanner
	}{
		{
			str: "go-json",
			want: &Scanner{
				Pos: &Pos{
					line: 1,
					col:  1,
				},
				offset: 0,
				str:    []rune("go-json"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScanner(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isEOF(t *testing.T) {
	type fields struct {
		Pos    *Pos
		offset int
		str    []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				offset: 0,
				str:    []rune("a"),
			},
			want: false,
		},
		{
			fields: fields{
				offset: 1,
				str:    []rune("a"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				Pos:    tt.fields.Pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			if got := s.isEOF(); got != tt.want {
				t.Errorf("Scanner.isEOF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_get(t *testing.T) {
	type fields struct {
		Pos    *Pos
		offset int
		str    []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{
			fields: fields{
				offset: 0,
				str:    []rune("a"),
			},
			want: 'a',
		},
		{
			fields: fields{
				offset: 1,
				str:    []rune("a"),
			},
			want: EOF,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				Pos:    tt.fields.Pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			if got := s.get(); got != tt.want {
				t.Errorf("Scanner.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_peek(t *testing.T) {
	type fields struct {
		Pos    *Pos
		offset int
		str    []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{
			fields: fields{
				offset: 0,
				str:    []rune("abc"),
			},
			want: 'b',
		},
		{
			fields: fields{
				offset: 1,
				str:    []rune("abc"),
			},
			want: 'c',
		},
		{
			fields: fields{
				offset: 2,
				str:    []rune("abc"),
			},
			want: EOF,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				Pos:    tt.fields.Pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			if got := s.peek(); got != tt.want {
				t.Errorf("Scanner.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_consume(t *testing.T) {
	s := NewScanner("abc")
	type fields struct {
		Pos    *Pos
		offset int
	}
	tests := []struct {
		name  string
		want  rune
		after fields
	}{
		{
			want: 'a',
			after: fields{
				Pos: &Pos{
					line: 1,
					col:  2,
				},
				offset: 1,
			},
		},
		{
			want: 'b',
			after: fields{
				Pos: &Pos{
					line: 1,
					col:  3,
				},
				offset: 2,
			},
		},
		{
			want: 'c',
			after: fields{
				Pos: &Pos{
					line: 1,
					col:  4,
				},
				offset: 3,
			},
		},
		{
			want: EOF,
			after: fields{
				Pos: &Pos{
					line: 1,
					col:  4,
				},
				offset: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.consume(); got != tt.want {
				t.Errorf("Scanner.consume() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(s.Pos, tt.after.Pos) {
				t.Errorf("after Pos = %v, want %v", s.Pos, tt.after.Pos)
			}
			if s.offset != tt.after.offset {
				t.Errorf("after offset = %v, want %v", s.offset, tt.after.offset)
			}
		})
	}
}

func TestScanner_skipSpaces(t *testing.T) {
	type fields struct {
		Pos    *Pos
		offset int
	}
	tests := []struct {
		name  string
		str   string
		after fields
	}{
		{
			str: `
		a`,
			after: fields{
				Pos: &Pos{
					line: 2,
					col:  3,
				},
				offset: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewScanner(tt.str)
			s.skipSpaces()
			if !reflect.DeepEqual(s.Pos, tt.after.Pos) {
				t.Errorf("after Pos = %v, want %v", s.Pos, tt.after.Pos)
			}
			if s.offset != tt.after.offset {
				t.Errorf("after offset = %v, want %v", s.offset, tt.after.offset)
			}
		})
	}
}
