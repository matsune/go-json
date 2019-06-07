package json

import (
	"testing"
)

func TestParser_must(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		c       rune
		wantErr bool
	}{
		{
			str:     "a",
			c:       'a',
			wantErr: false,
		},
		{
			str:     "a",
			c:       'b',
			wantErr: true,
		},
		{
			str:     "",
			c:       'a',
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.str)
			if err := p.must(tt.c); (err != nil) != tt.wantErr {
				t.Errorf("Parser.must() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_parseString(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    string
		wantErr bool
	}{
		{
			str:  `"abc"`,
			want: "abc",
		},
		{
			name: "quotation mark",
			str:  `"\"abc\""`,
			want: `"abc"`,
		},
		{
			name: "reverse solidus",
			str:  `"\\"`,
			want: `\`,
		},
		{
			name: "solidus",
			str:  `"\/"`,
			want: `/`,
		},
		{
			name: "backspace",
			str:  `"\b"`,
			want: "\b",
		},
		{
			name: "formfeed",
			str:  `"\f"`,
			want: "\f",
		},
		{
			name: "newline",
			str:  `"\n"`,
			want: "\n",
		},
		{
			name: "carriage return",
			str:  `"\r"`,
			want: "\r",
		},
		{
			name: "horizontal tab",
			str:  `"\t"`,
			want: "\t",
		},
		{
			name: "4hexadecimal digits",
			str:  `"\u1234"`,
			want: "\u1234",
		},
		{
			name:    "invalid escaping",
			str:     `"\"`,
			wantErr: true,
		},
		{
			name:    "invalid hexadecimal digits",
			str:     `"\u123"`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.str)
			got, err := p.parseString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.inner != tt.want {
				t.Errorf("Parser.parseString() = %v, want %v", got.inner, tt.want)
			}
		})
	}
}

func TestParser_parse4hex(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    rune
		wantErr bool
	}{
		{
			str:  "0020",
			want: '\u0020',
		},
		{
			str:     "002",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.str)
			got, err := p.parse4hex()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parse4hex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parser.parse4hex() = %v, want %v", got, tt.want)
			}
		})
	}
}
