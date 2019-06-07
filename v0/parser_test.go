package json

import (
	"reflect"
	"testing"
)

func Test_parser_parseNumber(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    interface{}
		wantErr bool
	}{
		{
			name: "u int",
			str:  "100,",
			want: int64(100),
		},
		{
			name: "s int",
			str:  "-100,",
			want: int64(-100),
		},
		{
			name: "u float",
			str:  "0.8329,",
			want: float64(0.8329),
		},
		{
			name: "s float",
			str:  "-0.8329,",
			want: float64(-0.8329),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				scanner: newScanner(tt.str),
			}
			got, err := p.parseNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.parseNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.parseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
