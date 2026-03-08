package main

import (
	"fmt"
	"io"
	"strings"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal int
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

func main() {
	input := "1 + 2;"

	tokenizer := NewTokenizer(input)
	for t, err := tokenizer.NextToken(); err != io.EOF; t, err = tokenizer.NextToken() {
		fmt.Println(t)
	}
}
