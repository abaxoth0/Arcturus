package lexer

import (
	"arcturus/packages/token"
)

func Parse(input []byte) []token.Token {
	tokens := []token.Token{}
	reading := false
	lexeme := make([]byte, 0, 256)

	for _, char := range input {
		switch char {
		case '{', '}', ';', ' ':
			if reading {
				reading = false
				tokens = append(tokens, tokenize(string(lexeme)))
				lexeme = lexeme[:0] // set length to 0 keeping capacity to avoid buffer reallocation
			}
			if char != ' ' {
				tokens = append(tokens, tokenize(string(char)))
			}
			continue
		case '\n':
			continue
		}

		if !reading {
			reading = true
		}
		lexeme = append(lexeme, char)
		// print(string(char))
	}
	return tokens
}

func tokenize(lexeme string) token.Token {
	switch lexeme {
	case token.SEMICOLON.Raw():
		return token.SEMICOLON
	case token.LBRACE.Raw():
		return token.LBRACE
	case token.RBRACE.Raw():
		return token.RBRACE
	case token.MESSAGE.Raw():
		return token.MESSAGE
	case token.INT.Raw():
		return token.INT
	case token.STRING.Raw():
		return token.STRING
	default:
		return token.NewIdentifier(lexeme)
	}
}
