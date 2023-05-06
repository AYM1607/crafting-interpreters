package runner

import (
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
	// Ignore whitespace.
	case ' ':
	case '\t':
	case '\r':
	// Handle new lines.
	case '\n':
		s.line += 1
	default:
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
