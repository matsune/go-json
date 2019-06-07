package json

import "fmt"

type Pos struct {
	line int
	col  int
}

func NewPos() *Pos {
	return &Pos{
		line: 1,
		col:  1,
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("line %d column %d", p.line, p.col)
}
