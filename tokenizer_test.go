package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "1 + 2;"
	tokenizer := NewTokenizer(input)

	tests := []struct {
		Type    TokenType
		Literal string
	}{
		{INTEGER, "1"},
		{PLUS, "+"},
		{INTEGER, "2"},
	}

	for _, test := range tests {
		next := tokenizer.NextToken()

		if next.Type != test.Type {
			t.Fatalf("expected %v; got %v", test.Type, next.Type)
		}

		if next.Literal != test.Literal {
			t.Fatalf("expected %v; got %v", test.Literal, next.Literal)
		}
	}
}
