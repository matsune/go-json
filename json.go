package json

import (
	"errors"
)

type state int

const (
	stateValue state = iota
	stateValueEnd
	stateKey
)

func Parse(str string) (Value, error) {
	p := newParser(str)

	var c rune
	var err error

	// check first char array or object
	if c, err = p.skipSpaces(); err != nil {
		// empty
		return nil, errors.New("empty contents")
	}
	if c != '{' && c != '[' {
		return nil, newError(p.pos, "expected { or [")
	}

	var value Value
	st := stateValue

loop:
	for {
		if c, err = p.skipSpaces(); err != nil {
			return nil, err
		}
		switch st {
		case stateKey:
			keyStr, err := p.parseString()
			if err != nil {
				return nil, err
			}
			if c, err = p.skipSpaces(); err != nil {
				return nil, err
			}
			if err := p.must(':'); err != nil {
				return nil, err
			}
			p.nestStack.peek().(*ObjectValue).AddKey(keyStr)
			st = stateValue
		case stateValue:
			switch c {
			case '{': // object start
				p.consume()
				last := p.nestStack.peek()
				if last != nil {
					last.addValue(p.pushObject())
				} else {
					p.pushObject()
				}

				st = stateKey

			case '[': // array start
				p.consume()
				last := p.nestStack.peek()
				if last != nil {
					last.addValue(p.pushArray())
				} else {
					p.pushObject()
				}

				st = stateValue

			case 'n': //null
				if err = p.musts("null"); err != nil {
					return nil, err
				}

				p.nestStack.peek().addValue(newNull())
				st = stateValueEnd

			case 'f': //false
				if err = p.musts("false"); err != nil {
					return nil, err
				}
				p.nestStack.peek().addValue(newBool(false))
				st = stateValueEnd

			case 't':
				if err = p.musts("true"); err != nil {
					return nil, err
				}
				p.nestStack.peek().addValue(newBool(true))
				st = stateValueEnd

			case '"':
				str, err := p.parseString()
				if err != nil {
					return nil, err
				}
				p.nestStack.peek().addValue(newString(str))
				st = stateValueEnd

			default:
				if isNum(c) || c == '-' { // number
					num, err := p.parseNumber()
					if err != nil {
						return nil, err
					}
					switch v := num.(type) {
					case int64:
						p.nestStack.peek().addValue(newInt(v))
					case float64:
						p.nestStack.peek().addValue(newFloat(v))
					default:
						return nil, p.unexpectError("unknown value ")
					}
					st = stateValueEnd
				} else {
					return nil, p.unexpectError("unexpected value")
				}
			}
		case stateValueEnd:
			// after value ',' or '}' or ']'
			isArr := p.isArr()
			isObj := p.isObj()
			if c == '}' {
				p.consume()
				var ok bool
				if value, ok = p.pop().(*ObjectValue); !ok {
					panic("closing } but stack last is not nestObj")
				}

				st = stateValueEnd
				if p.len() == 0 {
					break loop
				}
			} else if c == ']' {
				p.consume()
				if _, ok := p.pop().(*ArrayValue); !ok {
					panic("closing ] but stack last is not nestArr")
				}
				st = stateValueEnd
				if p.len() == 0 {
					break loop
				}
			} else if c == ',' {
				p.consume()
				if isArr {
					st = stateValue
				} else if isObj {
					st = stateKey
				} else {
					return nil, p.unexpectError("unexpected error")
				}
			} else {
				return nil, p.unexpectError("unexpected error")
			}
		}
	}
	return value, nil
}
