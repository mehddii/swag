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
	t.ConsumeWhitespace()

	ch := t.Char
	tok := NewToken(ILLEGAL, ch)

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
		tok.Literal = ""
	default:
		if IsLetter(ch) {
			tok.Literal = t.ReadIdent()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if IsDigit(ch) {
			tok.Literal = t.ReadNum()
			tok.Type = INTEGER
			return tok
		}
	}

	t.Advance()
	return tok
}

func (t *Tokenizer) ReadIdent() string {
	start := t.Current

	for IsLetter(t.Char) || IsDigit(t.Char) {
		t.Advance()
	}

	return t.Input[start:t.Current]
}

func (t *Tokenizer) ReadNum() string {
	start := t.Current

	for IsDigit(t.Char) {
		t.Advance()
	}

	return t.Input[start:t.Current]

}

func (t *Tokenizer) ConsumeWhitespace() {
	for IsSpace(t.Char) {
		t.Advance()
	}
}

// we need to handle integers and floats
