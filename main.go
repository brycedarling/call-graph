package main

import (
	"fmt"
	"lexer"
	"os"
	"strings"
	"token"
)

func main() {
	args := os.Args[1:]
	input := strings.Join(args, " ")
	l := lexer.New(input)
	tok := l.NextToken()
	for tok.Type != token.EOF {
		fmt.Println(tok)
		tok = l.NextToken()
	}
}
