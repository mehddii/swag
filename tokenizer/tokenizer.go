package tokenizer

type Tokenizer struct {
	Input   string
	Char    byte
	Current int
	Next    int
}

func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{
		Input: input,
	}
	t.Advance()

	return t
}

func (t *Tokenizer) Advance() {
	if t.Next >= len(t.Input) {
		t.Char = 0
		return
	}

	t.Current = t.Next
	t.Char = t.Input[t.Current]
	t.Next++
}

func (t *Tokenizer) NextToken() *Token {
	tok := NewToken(ILLEGAL, 0)
	ch := t.Char

	switch ch {
	case '=':
		tok = NewToken(ASSIGN, ch)
	case '+':
		tok = NewToken(PLUS, ch)
	case '-':
		tok = NewToken(MINUS, ch)
	case '*':
		tok = NewToken(ASTERISK, ch)
	case '/':
		tok = NewToken(SLASH, ch)
	case ';':
		tok = NewToken(SEMICOLON, ch)
	case ',':
		tok = NewToken(COMMA, ch)
	case '(':
		tok = NewToken(LPAREN, ch)
	case ')':
		tok = NewToken(RPAREN, ch)
	case '{':
		tok = NewToken(LBRACE, ch)
	case '}':
		tok = NewToken(RBRACE, ch)
	case 0:
		tok.Type = EOF
	}

	t.Advance()
	return tok
}
