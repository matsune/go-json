package json

import (
	"testing"
)

func Test_scanner_eof(t *testing.T) {
	type fields struct {
		pos    pos
		offset int
		str    []rune
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "empty str",
			fields: fields{
				pos:    newPos(),
				offset: 0,
				str:    []rune{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanner{
				pos:    tt.fields.pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			if err := s.eof(); (err != nil) != tt.wantErr {
				t.Errorf("scanner.eof() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanner_isEOF(t *testing.T) {
	type fields struct {
		pos    pos
		offset int
		str    []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanner{
				pos:    tt.fields.pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			if got := s.isEOF(); got != tt.want {
				t.Errorf("scanner.isEOF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanner_get(t *testing.T) {
	type fields struct {
		pos    pos
		offset int
		str    []rune
	}
	tests := []struct {
		name    string
		fields  fields
		want    rune
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanner{
				pos:    tt.fields.pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			got, err := s.get()
			if (err != nil) != tt.wantErr {
				t.Errorf("scanner.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("scanner.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanner_peek(t *testing.T) {
	type fields struct {
		pos    pos
		offset int
		str    []rune
	}
	tests := []struct {
		name    string
		fields  fields
		want    rune
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanner{
				pos:    tt.fields.pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			got, err := s.peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("scanner.peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("scanner.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanner_consume(t *testing.T) {
	type fields struct {
		pos    pos
		offset int
		str    []rune
	}
	tests := []struct {
		name    string
		fields  fields
		want    rune
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanner{
				pos:    tt.fields.pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			got, err := s.consume()
			if (err != nil) != tt.wantErr {
				t.Errorf("scanner.consume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("scanner.consume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanner_skipSpaces(t *testing.T) {
	type fields struct {
		pos    pos
		offset int
		str    []rune
	}
	tests := []struct {
		name    string
		fields  fields
		want    rune
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanner{
				pos:    tt.fields.pos,
				offset: tt.fields.offset,
				str:    tt.fields.str,
			}
			got, err := s.skipSpaces()
			if (err != nil) != tt.wantErr {
				t.Errorf("scanner.skipSpaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("scanner.skipSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
