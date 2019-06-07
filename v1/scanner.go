package json

const EOF = 0

type Scanner struct {
	*Pos
	offset int
	str    []rune
}

func NewScanner(str string) *Scanner {
	return &Scanner{
		Pos:    NewPos(),
		offset: 0,
		str:    []rune(str),
	}
}

func (s *Scanner) isEOF() bool {
	return s.offset >= len(s.str)
}

// get rune of current offset position
func (s *Scanner) get() rune {
	if s.isEOF() {
		return EOF
	}
	return s.str[s.offset]
}

// read next rune without step forward
func (s *Scanner) peek() rune {
	hasNext := (s.offset + 1) < len(s.str)
	if !hasNext {
		return EOF
	}
	return s.str[s.offset+1]
}

// get current rune and step forward if it was not EOF
func (s *Scanner) consume() rune {
	if s.isEOF() {
		return EOF
	}
	c := s.str[s.offset]
	if isNL(c) {
		s.line++
		s.col = 1
	} else {
		s.col++
	}
	s.offset++
	return c
}

// step forward until enconter EOF or not space
func (s *Scanner) skipSpaces() {
	for {
		if s.isEOF() {
			break
		}
		c := s.str[s.offset]
		if isNL(c) {
			s.line++
			s.col = 1
			s.offset++
		} else if isSpace(c) {
			s.col++
			s.offset++
		} else {
			break
		}
	}
}
