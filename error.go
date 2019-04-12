package json

import "fmt"

type ParseError struct {
	pos
	msg string
}

func newError(p pos, msg string) error {
	return &ParseError{
		pos: p,
		msg: msg,
	}
}

func (e ParseError) Error() string {
	return fmt.Sprintf("error at %s: %s", e.pos, e.msg)
}
