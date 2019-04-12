package json

import (
	"errors"
	"fmt"
)

type state int

const (
	stateValue state = iota
	stateValueEnd
	stateKey
)

func Parse(str string) error {
	p := newParser(str)

	var c rune
	var err error

	// check first char array or object
	if c, err = p.skipSpaces(); err != nil {
		// empty
		return errors.New("empty contents")
	}
	if c != '{' && c != '[' {
		return newError(p.pos, "expected { or [")
	}

	st := stateValue

loop:
	for {
		if c, err = p.skipSpaces(); err != nil {
			return err
		}
		switch st {
		case stateKey:
			keyStr, err := p.parseString()
			if err != nil {
				return err
			}
			fmt.Printf("%q", keyStr)
			if c, err = p.skipSpaces(); err != nil {
				return err
			}
			if err := p.must(':'); err != nil {
				return err
			}
			fmt.Print(":")
			st = stateValue
		case stateValue:
			switch c {
			case '{': // object start
				p.consume()
				p.push(nestObj)
				st = stateKey
				fmt.Println("{")

			case '[': // array start
				p.consume()
				p.push(nestArr)
				st = stateValue
				fmt.Print("[")

			case 'n': //null
				if err = p.musts("null"); err != nil {
					return err
				}
				st = stateValueEnd
				fmt.Print("null")

			case 'f': //false
				if err = p.musts("false"); err != nil {
					return err
				}
				st = stateValueEnd
				fmt.Print("false")

			case 't':
				if err = p.musts("true"); err != nil {
					return err
				}
				st = stateValueEnd
				fmt.Print("true")

			case '"':
				str, err := p.parseString()
				if err != nil {
					return err
				}
				st = stateValueEnd
				fmt.Printf("%q", str)

			default:
				if isNum(c) || c == '-' { // number
					num, err := p.parseNumber()
					if err != nil {
						return err
					}
					st = stateValueEnd
					fmt.Print(num)
				} else {
					return p.unexpectError("unexpected value")
				}
			}
		case stateValueEnd:
			// after value ',' or '}' or ']'
			last := p.nestStack.peek()
			isArr := last == nestArr
			isObj := last == nestObj
			if c == '}' {
				p.consume()
				if p.pop() != nestObj {
					panic("closing } but stack last is not nestObj")
				}

				st = stateValueEnd
				fmt.Print("}")
				if p.len() == 0 {
					break loop
				}
			} else if c == ']' {
				p.consume()
				if p.pop() != nestArr {
					panic("closing ] but stack last is not nestArr")
				}
				st = stateValueEnd
				fmt.Print("]")
				if p.len() == 0 {
					break loop
				}
			} else if c == ',' {
				p.consume()
				if isArr {
					st = stateValue
					fmt.Print(",")
				} else if isObj {
					st = stateKey
					fmt.Println(",")
				} else {
					return p.unexpectError("unexpected error")
				}
			} else {
				return p.unexpectError("unexpected error")
			}
		}
	}
	return nil
}
