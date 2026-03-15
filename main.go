package main

import (
	"fmt"
	"github.com/mehddii/swag/tokenizer"
)

func main() {
	input := "+()-=;"

	t := tokenizer.NewTokenizer(input)
	for tok := t.NextToken(); tok.Type != tokenizer.EOF; tok = t.NextToken() {
		fmt.Printf("Token Type: %q, Token Literal: %q\n", tok.Type, tok.Literal)
	}
}
