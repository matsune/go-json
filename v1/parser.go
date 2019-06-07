package json

import (
	"fmt"
	"strconv"
)

type Parser struct {
	*Scanner
	*Stack
}

func NewParser(str string) *Parser {
	return &Parser{
		Scanner: NewScanner(str),
		Stack:   NewStack(),
	}
}

func (p *Parser) error(msg string) *ParseError {
	return &ParseError{
		Pos: p.Pos,
		Msg: msg,
	}
}

func (p *Parser) unexpectedEOF() *ParseError {
	return p.error("Unexpected EOF")
}

func (p *Parser) unexpected(r rune) *ParseError {
	return p.error(fmt.Sprintf("Unexpected char %c", r))
}

func (p *Parser) must(c rune) error {
	r := p.consume()
	if r == EOF {
		return p.error(fmt.Sprintf("Expected %c but got EOF", c))
	}
	if r != c {
		return p.error(fmt.Sprintf("Expected %c but got %c", c, r))
	}
	return nil
}

// isHex{4}
func (p *Parser) parse4hex() (rune, error) {
	hexStr := make([]rune, 4)
	for i := 0; i < 4; i++ {
		r := p.consume()
		if r == EOF {
			return 0, p.unexpectedEOF()
		}
		if !isHex(r) {
			return 0, p.unexpected(r)
		}
		hexStr[i] = r
	}
	h, err := strconv.ParseUint(string(hexStr), 16, 16)
	if err != nil {
		return 0, err
	}
	return rune(h), nil
}

func (p *Parser) parseString() (*String, error) {
	var err error
	if err = p.must('"'); err != nil {
		return nil, err
	}
	var runes []rune
	isEscaping := false
	for {
		r := p.consume()
		if r == EOF {
			return nil, p.unexpectedEOF()
		}
		if r == '"' && !isEscaping {
			break
		}
		if isEscaping {
			switch r {
			case '"', '\\', '/':
				break
			case 'b':
				r = '\b'
			case 'f':
				r = '\f'
			case 'n':
				r = '\n'
			case 'r':
				r = '\r'
			case 't':
				r = '\t'
			case 'u':
				r, err = p.parse4hex()
				if err != nil {
					return nil, err
				}
			default:
				return nil, p.unexpected(r)
			}
			isEscaping = false
			runes = append(runes, r)
		} else {
			if r == '"' {
				break
			}
			if r == '\\' {
				isEscaping = true
				continue
			}

			runes = append(runes, r)
		}
	}
	return NewString(string(runes)), nil
}
