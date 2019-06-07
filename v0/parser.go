package json

import (
	"fmt"
	"strconv"
)

type parser struct {
	*scanner
	*nestStack
}

func newParser(str string) *parser {
	return &parser{
		scanner:   newScanner(str),
		nestStack: newStack(),
	}
}

func (p *parser) expectError(expect, got rune) error {
	return newError(p.pos, fmt.Sprintf("expected char `%c` but got `%c`", expect, got))
}

func (p *parser) expectsError(expects string) error {
	return newError(p.pos, fmt.Sprintf("expected `%s`", expects))
}

func (p *parser) unexpectError(msg string) error {
	return newError(p.pos, msg)
}

func (p *parser) must(b rune) error {
	c, err := p.consume()
	if err != nil {
		return err
	}
	if c != b {
		return p.expectError(b, c)
	}
	return nil
}

func (p *parser) musts(str string) error {
	for _, r := range []rune(str) {
		c, err := p.consume()
		if err != nil {
			return err
		}
		if c != r {
			return p.expectsError(str)
		}
	}
	return nil
}

func (p *parser) parseString() (string, error) {
	var err error
	if err = p.must('"'); err != nil {
		return "", err
	}
	var runes []rune
	for {
		r, err := p.consume()
		if err != nil {
			return "", err
		}
		if r == '"' {
			break
		} else {
			runes = append(runes, r)
		}
	}
	return string(runes), nil
}

// int or float
func (p *parser) parseNumber() (interface{}, error) {
	var str []rune
	hasDot := false
	for {
		r, err := p.get()
		if err != nil {
			return nil, err
		}
		if r == '.' {
			if hasDot {
				return nil, p.unexpectError("invalid number format")
			}
			hasDot = true
		} else if r == '-' {
			if len(str) != 0 {
				return nil, p.unexpectError("invalid number format")
			}
		} else if !isNum(r) {
			break
		}
		str = append(str, r)
		p.consume()
	}
	if hasDot {
		f, err := strconv.ParseFloat(string(str), 64)
		if err != nil {
			return nil, p.unexpectError(fmt.Sprintf("could not parse as float `%s`", string(str)))
		}
		return f, nil
	}
	i, err := strconv.ParseInt(string(str), 10, 64)
	if err != nil {
		return nil, p.unexpectError(fmt.Sprintf("could not parse as int `%s`", string(str)))
	}
	return i, nil
}
