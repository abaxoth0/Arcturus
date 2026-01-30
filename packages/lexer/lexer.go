package lexer

import (
	"arcturus/packages/token"
)

func Parse(input []byte) []token.Token {
	tokens := []token.Token{}
	reading := false
	lexeme := make([]byte, 0, 256)

	for i, char := range input {
		switch char {
		case '{', '}', ';', ' ':
			// TODO: FIX
			// Currently it works fine, but i still think that this i+1 is kinda problematic,
			// since there are no any bounds checking before it.
			if char == ' ' && char == input[i+1] {
				continue
			}
			if reading {
				reading = false
				tokens = append(tokens, tokenize(string(lexeme)))
				lexeme = lexeme[:0] // set length to 0 keeping capacity to avoid buffer reallocation
			}
			tokens = append(tokens, tokenize(string(char)))
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
	case token.WHITESPACE.Raw():
		return token.WHITESPACE
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
