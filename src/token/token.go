package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"

	// Delimiters
	COMMA = "COMMA"

	// Identifiers & literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Keywords
	NS   = "NS"
	DEF  = "DEF"
	DEFN = "DEFN"
	LET  = "LET"

	// Operators
	PLUS = "+"

	// Characters and Strings
	CHAR   = "CHAR"
	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"ns":   NS,
	"def":  DEF,
	"defn": DEFN,
	"let":  LET,
}

func GetTokenType(value string) TokenType {
	if tokenType, ok := keywords[value]; ok {
		return tokenType
	}
	return IDENTIFIER
}

func New(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}
