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

	def fibo = func(n) {
		if n == 0 or n == 1 {
			return 1;
		}

		return fibo(n - 2) + fibo(n - 1);
	};

	def factorial = func(n) {
		def result = 1;

		while n != 0 {
			result = result * n;
			n = n - 1;
		}

		return result;
	};

	if add(1, 2) == 3 and factorial(1) == 1 {
		print(1);
	} else {
		print(0);
	}
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

		{DEF, "def"},
		{IDENT, "fibo"},
		{ASSIGN, "="},
		{FUNC, "func"},
		{LPAREN, "("},
		{IDENT, "n"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{IF, "if"},
		{IDENT, "n"},
		{EQ_EQ, "=="},
		{INTEGER, "0"},
		{OR, "or"},
		{IDENT, "n"},
		{EQ_EQ, "=="},
		{INTEGER, "1"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{INTEGER, "1"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{RETURN, "return"},
		{IDENT, "fibo"},
		{LPAREN, "("},
		{IDENT, "n"},
		{MINUS, "-"},
		{INTEGER, "2"},
		{RPAREN, ")"},
		{PLUS, "+"},
		{IDENT, "fibo"},
		{LPAREN, "("},
		{IDENT, "n"},
		{MINUS, "-"},
		{INTEGER, "1"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},

		{DEF, "def"},
		{IDENT, "factorial"},
		{ASSIGN, "="},
		{FUNC, "func"},
		{LPAREN, "("},
		{IDENT, "n"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{DEF, "def"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{INTEGER, "1"},
		{SEMICOLON, ";"},
		{WHILE, "while"},
		{IDENT, "n"},
		{BANG_EQ, "!="},
		{INTEGER, "0"},
		{LBRACE, "{"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "result"},
		{ASTERISK, "*"},
		{IDENT, "n"},
		{SEMICOLON, ";"},
		{IDENT, "n"},
		{ASSIGN, "="},
		{IDENT, "n"},
		{MINUS, "-"},
		{INTEGER, "1"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{RETURN, "return"},
		{IDENT, "result"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},

		{IF, "if"},
		{IDENT, "add"},
		{LPAREN, "("},
		{INTEGER, "1"},
		{COMMA, ","},
		{INTEGER, "2"},
		{RPAREN, ")"},
		{EQ_EQ, "=="},
		{INTEGER, "3"},
		{AND, "and"},
		{IDENT, "factorial"},
		{LPAREN, "("},
		{INTEGER, "1"},
		{RPAREN, ")"},
		{EQ_EQ, "=="},
		{INTEGER, "1"},
		{LBRACE, "{"},
		{PRINT, "print"},
		{LPAREN, "("},
		{INTEGER, "1"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{ELSE, "else"},
		{LBRACE, "{"},
		{PRINT, "print"},
		{LPAREN, "("},
		{INTEGER, "0"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},

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
