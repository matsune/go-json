package json

import (
	"errors"
	"fmt"
	"strconv"
)

func Parse(str string) (Value, error) {
	p := NewParser(str)
	p.skipSpaces()
	if p.get() == EOF {
		return nil, errors.New("json is empty")
	}
	return p.parseJson()
}

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

func (p *Parser) unexpected(r rune) *ParseError {
	if r == 0 {
		return p.error("Unexpected EOF")
	} else {
		return p.error(fmt.Sprintf("Unexpected char %c", r))
	}
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

func (p *Parser) musts(str string) error {
	for _, r := range []rune(str) {
		c := p.consume()
		if c != r {
			return p.error(fmt.Sprintf("Expected %s", str))
		}
	}
	return nil
}

// Json := <Object> | <Array>
func (p *Parser) parseJson() (Value, error) {
	c := p.get()
	switch c {
	case '{':
		return p.parseObject()
	case '[':
		return p.parseArray()
	default:
		return nil, p.unexpected(c)
	}
}

// Object := '{' <s>* '}' | '{' <s>* <Members> <s>* '}'
// Members := <Pair> | <Pair> <s>* ',' <s>* <Members>
func (p *Parser) parseObject() (*Object, error) {
	err := p.must('{')
	if err != nil {
		return nil, err
	}
	obj := p.pushObject()

	p.skipSpaces()

	// key string or }
	c := p.get()
	if c == '"' {
		for {
			// parse pairs and add to this obj
			pair, err := p.parsePair()
			if err != nil {
				return nil, err
			}
			obj.AddValue(pair)
			p.skipSpaces()

			c = p.get()
			if c == ',' {
				p.consume()
				p.skipSpaces()
			} else {
				break
			}
		}
	}

	if c == '}' {
		p.consume()
		p.pop()
	}
	return obj, nil
}

// Pair := <String> <s>* ':' <s>* <Value>
func (p *Parser) parsePair() (*Pair, error) {
	s, err := p.parseString()
	if err != nil {
		return nil, err
	}
	p.skipSpaces()
	if err = p.must(':'); err != nil {
		return nil, err
	}
	p.skipSpaces()
	v, err := p.parseValue()
	if err != nil {
		return nil, err
	}
	return NewPair(s.inner, v), nil
}

// Value := <String> | <Number> | <Object> | <Array> | <Bool> | null
func (p *Parser) parseValue() (Value, error) {
	c := p.get()
	if c == '"' {
		return p.parseString()
	} else if c == '-' || isDigit(c) {
		return p.parseNumber()
	} else if c == '{' {
		return p.parseObject()
	} else if c == '[' {
		return p.parseArray()
	} else if c == 't' || c == 'f' {
		return p.parseBool()
	} else if c == 'n' {
		return p.parseNull()
	} else {
		return nil, p.unexpected(c)
	}
}

// Array := '[' <s>* ']' | '[' <s>* <Elements> <s>* ']'
// Elements := <Value> | <Value> <s>* ',' <s>* <Elements>
func (p *Parser) parseArray() (*Array, error) {
	err := p.must('[')
	if err != nil {
		return nil, err
	}
	arr := p.pushArray()

	p.skipSpaces()

	// Elements or ]
	c := p.get()
	if c == ']' {
		p.consume()
		// ']'
		p.pop()
		return arr, nil
	}
	// Elements
	for {
		// parse Value and add to this arr
		v, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		arr.AddValue(v)
		p.skipSpaces()

		c = p.get()
		if c == ',' {
			p.consume()
			p.skipSpaces()
		} else {
			break
		}
	}
	if err = p.must(']'); err != nil {
		return nil, err
	}

	return arr, nil
}

// Bool := true | false
func (p *Parser) parseBool() (*Bool, error) {
	c := p.get()
	if c == 't' {
		if err := p.musts("true"); err != nil {
			return nil, err
		}
		return NewBool(true), nil
	} else if c == 'f' {
		if err := p.musts("false"); err != nil {
			return nil, err
		}
		return NewBool(false), nil
	} else {
		return nil, p.unexpected(c)
	}
}

// Null := null
func (p *Parser) parseNull() (*Null, error) {
	if err := p.musts("null"); err != nil {
		return nil, err
	}
	return NewNull(), nil
}

// 4hex := <Hex>{4}
func (p *Parser) parse4hex() (rune, error) {
	hexStr := make([]rune, 4)
	for i := 0; i < 4; i++ {
		r := p.consume()
		if r == EOF || !isHex(r) {
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

// String := '"' ({Unescaped}|'\'(["\/bfnrt]|'u'<4hex>))* '"'
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
			return nil, p.unexpected(r)
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

// Number := '-'?('0'|<Digit9><Digit>*)('.'<Digit>+)?([Ee][+-]?<Digit>+)?
func (p *Parser) parseNumber() (Value, error) {
	c := p.get()
	if c == EOF {
		return nil, p.unexpected(c)
	}
	r := []rune{}

	isFloat := false
	// '-'?
	if c == '-' {
		r = append(r, c)
		c = p.next()
	}
	// 0|[1-9]<digit>*
	if c == '0' {
		r = append(r, c)
		c = p.next()
		if c == EOF {
			return NewInt(0), nil
		}
	} else if isDigit9(c) {
		r = append(r, c)
		c = p.next()

		for {
			if isDigit(c) {
				r = append(r, c)
				c = p.next()
			} else {
				break
			}
		}
	} else {
		return nil, p.unexpected(c)
	}

	if c == '.' {
		isFloat = true
		r = append(r, c)
		// '.'<digit>+
		c = p.next()
		if c == EOF {
			return nil, p.unexpected(c)
		}
		// <digit>+
		if !isDigit(c) {
			return nil, p.unexpected(c)
		}
		for {
			r = append(r, c)
			c = p.next()
			if !isDigit(c) {
				break
			}
		}
	}
	if c == 'e' || c == 'E' {
		isFloat = true
		r = append(r, c)
		// [eE][+-]?<digit>+
		c = p.next()
		if c == EOF {
			return nil, p.unexpected(c)
		}
		if isSign(c) {
			r = append(r, c)
			c = p.next()
			if c == EOF {
				return nil, p.unexpected(c)
			}
		}
		// <digit>+
		if !isDigit(c) {
			return nil, p.unexpected(c)
		}
		for {
			r = append(r, c)
			c = p.next()
			if !isDigit(c) {
				break
			}
		}
	}

	if isFloat {
		f, err := strconv.ParseFloat(string(r), 64)
		if err != nil {
			return nil, p.error(err.Error())
		}
		return NewFloat(f), nil
	} else {
		n, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			return nil, p.error(err.Error())
		}
		return NewInt(n), nil
	}
}
