package lexer

import (
	"testing"
	"token"
)

func TestNextToken(t *testing.T) {
	input := `(ns main.core) (def five 5) (def ten 10) (defn add [x, y] (+ x y)) (defn main [] (let [result (add five ten)] (print result)))`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.NS, "ns"},
		{token.IDENTIFIER, "main.core"},
		{token.RPAREN, ")"},

		{token.LPAREN, "("},
		{token.DEF, "def"},
		{token.IDENTIFIER, "five"},
		{token.INT, "5"},
		{token.RPAREN, ")"},

		{token.LPAREN, "("},
		{token.DEF, "def"},
		{token.IDENTIFIER, "ten"},
		{token.INT, "10"},
		{token.RPAREN, ")"},

		{token.LPAREN, "("},
		{token.DEFN, "defn"},
		{token.IDENTIFIER, "add"},
		{token.LBRACKET, "["},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RBRACKET, "]"},
		{token.LPAREN, "("},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "x"},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},

		{token.LPAREN, "("},
		{token.DEFN, "defn"},
		{token.IDENTIFIER, "main"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LPAREN, "("},
		{token.LET, "let"},
		{token.LBRACKET, "["},
		{token.IDENTIFIER, "result"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "add"},
		{token.IDENTIFIER, "five"},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.RBRACKET, "]"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "print"},
		{token.IDENTIFIER, "result"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
