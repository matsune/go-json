package json

import "fmt"

type ParseError struct {
	*Pos
	Msg string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("error at %s: %s", e.Pos, e.Msg)
}
