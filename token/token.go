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
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	// Two-char
	EQ = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {	// Check the map. map returns (value, exists).
		return tok
	}
	return IDENT
}