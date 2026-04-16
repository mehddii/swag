package parser

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expression()
}

type Statement interface {
	Node
	statement()
}
