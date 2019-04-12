package json

import "fmt"

type pos struct {
	line int
	col  int
}

func newPos() pos {
	return pos{
		line: 1,
		col:  1,
	}
}

func (p pos) String() string {
	return fmt.Sprintf("line %d column %d", p.line, p.col)
}
