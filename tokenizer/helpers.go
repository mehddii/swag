package tokenizer

func IsLetter(char byte) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_'
}

func IsDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func IsSpace(char byte) bool {
	return char == ' ' || char == '\t' ||
		char == '\n' || char == '\r'
}
