package parser

import (
	"github.com/mehddii/swag/tokenizer"
)

type Parser struct {
	Tokenizer *tokenizer.Tokenizer
	Current   *tokenizer.Token
	Next      *tokenizer.Token
}

func NewParser(tokenizer *tokenizer.Tokenizer) *Parser {
	p := &Parser{
		Tokenizer: tokenizer,
	}

	p.Advance()
	p.Advance()

	return p
}

func (p *Parser) Advance() {
	p.Current = p.Next
	p.Next = p.Tokenizer.NextToken()
}

func (p *Parser) Parse() *Program {
	return nil
}
