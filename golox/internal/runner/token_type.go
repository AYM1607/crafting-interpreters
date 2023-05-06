package runner

type TokenType string

const (
	// Single character tokens.
	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	LBRACE TokenType = "LBRACE"
	RBRACE TokenType = "RBRACE"
	COMMA  TokenType = "COMMA"
	DOT    TokenType = "DOT"
	MINUS  TokenType = "MINUS"
	PLUS   TokenType = "PLUS"
	SEMI   TokenType = "SEMI"
	SLASH  TokenType = "SLASH"
	STAR   TokenType = "STAR"

	// One or two character tokens.
	BANG        TokenType = "BANG"
	BANG_EQUAL  TokenType = "BANG_EQUAL"
	EQUAL       TokenType = "EQUAL"
	EQUAL_EQUAL TokenType = "EQUAL_EQUAL"
	GT          TokenType = "GT"
	GTE         TokenType = "GTE"
	LT          TokenType = "LT"
	LTE         TokenType = "LTE"

	// Literals.
	IDENT  TokenType = "IDENT"
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"

	// Keywords
	AND    TokenType = "AND"
	CLASS  TokenType = "CLASS"
	ELSE   TokenType = "ELSE"
	FALSE  TokenType = "FALSE"
	FUN    TokenType = "FUN"
	FOR    TokenType = "FOR"
	IF     TokenType = "IF"
	NIL    TokenType = "NIL"
	OR     TokenType = "OR"
	PRINT  TokenType = "PRINT"
	RETURN TokenType = "RETURN"
	SUPER  TokenType = "SUPER"
	THIS   TokenType = "THIS"
	TRUE   TokenType = "TRUE"
	VAR    TokenType = "VAR"
	WHILE  TokenType = "WHILE"

	EOF TokenType = "EOF"
)
