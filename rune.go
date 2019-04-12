package json

func isNL(c rune) bool {
	return c == '\n'
}

func isSpace(c rune) bool {
	return c == ' '
}

func isNum(r rune) bool {
	return 0x30 <= r && r <= 0x39
}
