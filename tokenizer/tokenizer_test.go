package tokenizer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+-*/({});"
	tokenizer := NewTokenizer(input)

	tests := []struct {
		Type    TokenType
		Literal string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{MINUS, "-"},
		{ASTERISK, "*"},
		{SLASH, "/"},
		{LPAREN, "("},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
	}

	for _, test := range tests {
		tok := tokenizer.NextToken()

		if tok.Type != test.Type {
			t.Fatalf("expected %v; got %v", test.Type, tok.Type)
		}

		if tok.Literal != test.Literal {
			t.Fatalf("expected %v; got %v", test.Literal, tok.Literal)
		}
	}
}
