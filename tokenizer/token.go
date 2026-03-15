package tokenizer

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal byte) *Token {
	return &Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func (tokType TokenType) String() string {
	return repr[tokType]
}

const (
	// LITERALS
	INTEGER = iota

	// IDENTIFIERS
	IDENT

	// symbols
	ASSIGN
	PLUS
	MINUS
	ASTERISK
	SLASH
	SEMICOLON
	COMMA
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// keywords
	FUNC
	DEF
	RETURN

	ILLEGAL
	EOF
)

var (
	repr = map[TokenType]string{
		INTEGER:   "INTEGER",
		IDENT:     "IDENT",
		ASSIGN:    "ASSIGN",
		PLUS:      "PLUS",
		MINUS:     "MINUS",
		ASTERISK:  "ASTERISK",
		SLASH:     "SLASH",
		SEMICOLON: "SEMICOLON",
		COMMA:     "COMMA",
		LPAREN:    "LPAREN",
		RPAREN:    "RPAREN",
		LBRACE:    "LBRACE",
		RBRACE:    "RBRACE",
		FUNC:      "FUNC",
		DEF:       "DEF",
		RETURN:    "RETURN",
		ILLEGAL:   "ILLEGAL",
		EOF:       "EOF",
	}
)
