package tokenizer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	def a = 4;
	def f = func(x) {
		return 2 * x;
	};

	f(a);

	def add = func(x, y) {
		return x + y;
	};

	add(3, 5);
	`
	tokenizer := NewTokenizer(input)

	tests := []struct {
		Type    TokenType
		Literal string
	}{
		{DEF, "def"},
		{IDENT, "a"},
		{ASSIGN, "="},
		{INTEGER, "4"},
		{SEMICOLON, ";"},

		{DEF, "def"},
		{IDENT, "f"},
		{ASSIGN, "="},
		{FUNC, "func"},
		{LPAREN, "("},
		{IDENT, "x"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{INTEGER, "2"},
		{ASTERISK, "*"},
		{IDENT, "x"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},

		{IDENT, "f"},
		{LPAREN, "("},
		{IDENT, "a"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{DEF, "def"},
		{IDENT, "add"},
		{ASSIGN, "="},
		{FUNC, "func"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},

		{IDENT, "add"},
		{LPAREN, "("},
		{INTEGER, "3"},
		{COMMA, ","},
		{INTEGER, "5"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{EOF, ""},
	}

	for _, test := range tests {
		tok := tokenizer.NextToken()

		if tok.Type != test.Type {
			t.Fatalf("expected %v; got %v", test.Type, tok.Type)
		}

		if tok.Literal != test.Literal {
			t.Fatalf("expected %q; got %q", test.Literal, tok.Literal)
		}
	}
}
