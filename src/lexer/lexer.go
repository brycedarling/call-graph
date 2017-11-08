package lexer

import (
	"token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = token.New(token.LPAREN, "(")
	case ')':
		tok = token.New(token.RPAREN, ")")
	case '[':
		tok = token.New(token.LBRACKET, "[")
	case ']':
		tok = token.New(token.RBRACKET, "]")
	case ',':
		tok = token.New(token.COMMA, ",")
	case '+':
		tok = token.New(token.PLUS, "+")
	case 0:
		tok = token.New(token.EOF, "")
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := token.GetTokenType(literal)
			tok = token.New(tokenType, literal)
			return tok
		} else if isDigit(l.ch) {
			tok = token.New(token.INT, l.readNumber())
			return tok
		} else if isDoubleQuote(l.ch) {
			tok = token.New(token.STRING, l.readString())
			return tok
		} else if isSingleQuote(l.ch) {
			tok = token.New(token.CHAR, l.readCharacter())
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) read(start int, end int) string {
	return l.input[start:end]
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isUnderscore(l.ch) || isPeriod(l.ch) {
		l.readChar()
	}
	return l.read(start, l.position)
}

func (l *Lexer) readString() string {
	l.readChar() // consume opening "
	start := l.position

	for !isDoubleQuote(l.ch) {
		l.readChar()
	}

	end := l.position

	l.readChar() // consume closing "

	return l.read(start, end)
}

func (l *Lexer) readCharacter() string {
	l.readChar() // consume opening '
	start := l.position

	for !isSingleQuote(l.ch) {
		l.readChar()
	}

	end := l.position

	l.readChar() // consume closing '

	return l.read(start, end)
}

func isDoubleQuote(ch byte) bool {
	return ch == '"'
}

func isSingleQuote(ch byte) bool {
	return ch == '\''
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isUnderscore(ch byte) bool {
	return ch == '_'
}

func isPeriod(ch byte) bool {
	return ch == '.'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
