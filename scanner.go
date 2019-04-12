package json

type scanner struct {
	pos
	offset int
	str    []rune
}

func newScanner(str string) *scanner {
	return &scanner{
		pos:    newPos(),
		offset: 0,
		str:    []rune(str),
	}
}

func (s *scanner) eof() error {
	return newError(s.pos, "unexpected EOF")
}

func (s *scanner) isEOF() bool {
	return s.offset >= len(s.str)
}

func (s *scanner) get() (rune, error) {
	if s.isEOF() {
		return 0, s.eof()
	}
	return s.str[s.offset], nil
}

func (s *scanner) peek() (rune, error) {
	isEof := (s.offset + 1) >= len(s.str)
	if isEof {
		return 0, s.eof()
	}
	c := s.str[s.offset+1]
	return c, nil
}

func (s *scanner) consume() (rune, error) {
	if s.isEOF() {
		return 0, s.eof()
	}
	c := s.str[s.offset]
	if isNL(c) {
		s.line++
		s.col = 1
	}
	s.offset++
	return c, nil
}

func (s *scanner) skipSpaces() (rune, error) {
	var c rune
	var err error
	for {
		c, err = s.get()
		if err != nil {
			return 0, err
		}
		if isNL(c) {
			s.line++
			s.col = 1
		} else if !isSpace(c) {
			break
		}
		s.offset++
	}
	return c, nil
}
