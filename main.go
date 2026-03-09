package main

import (
	"fmt"
	"io"
)

func main() {
	input := "1 + 2;"

	tokenizer := NewTokenizer(input)
	for t, err := tokenizer.NextToken(); err != io.EOF; t, err = tokenizer.NextToken() {
		fmt.Println(t)
	}
}
