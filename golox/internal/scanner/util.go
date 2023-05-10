package scanner

func isIdentAlphaNumeric(c byte) bool {
	return isIdentAlpha(c) || isDigit(c)
}

func isIdentAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
