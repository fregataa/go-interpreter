package token

type TokenType string

type Token struct {
	Type TokenType	// To disitinguish whether it is a 'number' or a ']'
	Literal string	// The original value of token. e.a - 5 is INT Type and '5' Literal.
}

// Since there are a limited number of TokenType,
// define it with constants.
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifier, Literal
	IDENT = "IDENT"		// x, y, foobar, ...
	INT = "INT"			// 123456

	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)