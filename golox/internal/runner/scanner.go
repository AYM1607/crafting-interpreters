package runner

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
	}
}

func (s *Scanner) advance() byte {
	idx := s.current
	s.current += 1
	return s.source[idx]
}

func (s *Scanner) addToken(typ TokenType) {
	s.addTokenWithLiteral(typ, nil)
}

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
