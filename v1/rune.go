package json

func isNL(r rune) bool {
	return r == '\n'
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

func isDigit(r rune) bool {
	return 0x30 <= r && r <= 0x39
}

func isDigit9(r rune) bool {
	return 0x31 <= r && r <= 0x39
}

func isSign(r rune) bool {
	return r == '+' || r == '-'
}
