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

	// SYMBOLS
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
	DOT
	BANG
	EQ_EQ
	BANG_EQ
	LESS
	LESS_EQ
	GREATER
	GREATER_EQ

	// KEYWORDS
	FUNC
	DEF
	PRINT
	RETURN
	TRUE
	FALSE
	IF
	ELSE
	FOR
	WHILE
	AND
	OR
	NIL

	ILLEGAL
	EOF
)

var (
	repr = map[TokenType]string{
		INTEGER:    "INTEGER",
		IDENT:      "IDENT",
		ASSIGN:     "ASSIGN",
		PLUS:       "PLUS",
		MINUS:      "MINUS",
		ASTERISK:   "ASTERISK",
		SLASH:      "SLASH",
		SEMICOLON:  "SEMICOLON",
		COMMA:      "COMMA",
		LPAREN:     "LPAREN",
		RPAREN:     "RPAREN",
		LBRACE:     "LBRACE",
		RBRACE:     "RBRACE",
		FUNC:       "FUNC",
		DEF:        "DEF",
		PRINT:      "PRINT",
		RETURN:     "RETURN",
		TRUE:       "TRUE",
		FALSE:      "FALSE",
		IF:         "IF",
		ELSE:       "ELSE",
		FOR:        "FOR",
		WHILE:      "WHILE",
		AND:        "AND",
		OR:         "OR",
		NIL:        "NIL",
		ILLEGAL:    "ILLEGAL",
		DOT:        "DOT",
		BANG:       "BANG",
		EQ_EQ:      "EQ_EQ",
		BANG_EQ:    "BANG_EQ",
		LESS:       "LESS",
		LESS_EQ:    "LESS_EQ",
		GREATER:    "GREATER",
		GREATER_EQ: "GREATER_EQ",
		EOF:        "EOF",
	}

	keywords = map[string]TokenType{
		"def":    DEF,
		"func":   FUNC,
		"print":  PRINT,
		"return": RETURN,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"for":    FOR,
		"while":  WHILE,
		"and":    AND,
		"or":     OR,
		"nil":    NIL,
	}
)

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}
