package main

import (
	"strings"
)

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	INTEGER = iota
	PLUS
	SEMICOLON
)

type Tokenizer struct {
	r *strings.Reader
}

func NewTokenizer(input string) Tokenizer {
	return Tokenizer{
		r: strings.NewReader(input),
	}
}

func (t Tokenizer) NextToken() (string, error) {
	b := make([]byte, 1)
	_, err := t.r.Read(b)

	return string(b[0]), err
}
