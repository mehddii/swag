package main

import (
	"fmt"
	"github.com/mehddii/swag/tokenizer"
	"io"
)

func main() {
	input := "1 + 2;"

	tokenizer := tokenizer.NewTokenizer(input)
	for t, err := tokenizer.NextToken(); err != io.EOF; t, err = tokenizer.NextToken() {
		fmt.Println(t)
	}
}
