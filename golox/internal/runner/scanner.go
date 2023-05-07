package runner

import (
	"strconv"

	lerrors "github.com/AYM1607/crafting-interpreters/golox/internal/errors"
)

type Scanner struct {
	source string

	// State.
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		tokens: []Token{},

		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LPAREN)
	case ')':
		s.addToken(RPAREN)
	case '{':
		s.addToken(LBRACE)
	case '}':
		s.addToken(RBRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMI)
	case '*':
		s.addToken(STAR)
	case '!':
		tok := BANG
		if s.match('=') {
			tok = BANG_EQUAL
		}
		s.addToken(tok)
	case '=':
		tok := EQUAL
		if s.match('=') {
			tok = EQUAL_EQUAL
		}
		s.addToken(tok)
	case '<':
		tok := LT
		if s.match('=') {
			tok = LTE
		}
		s.addToken(tok)
	case '>':
		tok := GT
		if s.match('=') {
			tok = GTE
		}
		s.addToken(tok)
	case '/':
		if s.match('/') {
			// Consume all characters in a line comment.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case '"':
		s.scanString()
	// Ignore whitespace.
	case ' ':
	case '\t':
	case '\r':
	// Handle new lines.
	case '\n':
		s.line += 1
	default:
		// NOTE: adding this here to avoid listing all digits in a case.
		if isDigit(c) {
			s.scanNumber()
			return
		}
		if isIdentAlpha(c) {
			s.scanIdentifier()
			return
		}
		lerrors.EmitError(s.line, "Unexpected character.")
	}
}

// advance consumes a single character from the source.
func (s *Scanner) advance() byte {
	idx := s.current
	s.current += 1
	return s.source[idx]
}

// match returns true if the given byte is equal to the next one in source,
// it consumes the character if so.
func (s *Scanner) match(c byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != c {
		return false
	}

	// Next character in the source matches.
	s.current += 1
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNex() byte {
	idx := s.current + 1
	if idx >= len(s.source) {
		return 0
	}
	return s.source[idx]
}

func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		// Lox allows multi-line strings.
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		lerrors.EmitError(s.line, "Unterminated string.")
		return
	}

	// Consume the closing "
	s.advance()

	// Trim enclosing quotes
	val := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(STRING, val)
}

func (s *Scanner) scanNumber() {
	// Consume all digits preceding a dot (if any)
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for a decimal part.
	// Only literals in the form 123 and 123.123 are allowed.
	if s.peek() == '.' && isDigit(s.peekNex()) {
		// Only consume the dot if we're sure the format is valid.
		s.advance()

		// Consume the rest of the digis.
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	// NOTE: Ignoring error because we're sure the string follows the float
	// format. This should probably still report it but will leave as-is
	// for now.
	val, _ := strconv.ParseFloat(
		s.source[s.start:s.current],
		64,
	)
	s.addTokenWithLiteral(
		NUMBER,
		val,
	)
}

func (s *Scanner) scanIdentifier() {
	for isIdentAlphaNumeric(s.peek()) {
		s.advance()
	}
	l := s.source[s.start:s.current]
	typ := IDENT
	if kTyp, ok := KeywordTypes[l]; ok {
		typ = kTyp
	}
	s.addToken(typ)
}

// addToken produces a single token without a literal value.
func (s *Scanner) addToken(typ TokenType) {
	s.addTokenWithLiteral(typ, nil)
}

// addTokenWithLiteral produces a single token with the given literal value.
func (s *Scanner) addTokenWithLiteral(typ TokenType, literal interface{}) {
	lexme := s.source[s.start:s.current]
	s.tokens = append(
		s.tokens,
		NewToken(typ, lexme, literal, s.line),
	)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
