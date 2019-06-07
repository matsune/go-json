package json

func isNL(r rune) bool {
	return r == '\n'
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

// [0-9]
func isDigit(r rune) bool {
	return 0x30 <= r && r <= 0x39
}

// [1-9]
func isDigit9(r rune) bool {
	return 0x31 <= r && r <= 0x39
}

// [+-]
func isSign(r rune) bool {
	return r == '+' || r == '-'
}

// [0-9a-fA-F]
func isHex(r rune) bool {
	return isDigit(r) || (0x41 <= r && r <= 0x46) || (0x61 <= r && r <= 0x66)
}
