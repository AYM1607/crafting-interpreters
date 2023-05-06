package runner

import "fmt"

type Token struct {
	Type    TokenType
	Lexme   string
	Literal interface{}
	Line    int
}

func NewToken(
	typ TokenType,
	lexme string,
	lit interface{},
	line int,
) Token {
	return Token{
		Type:    typ,
		Lexme:   lexme,
		Literal: lit,
		Line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf(
		"%s %s %v",
		t.Type,
		t.Lexme,
		t.Literal,
	)
}
