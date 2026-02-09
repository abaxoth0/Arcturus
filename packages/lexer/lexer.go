package lexer

import (
	"arcturus/packages/common/logger"
	"arcturus/packages/token"

	log "github.com/abaxoth0/Ain/logger"
)

var lexerLogger = log.NewSource("LEXER", logger.Default)

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
	lexerLogger.Trace("lexeme: "+lexeme, nil)
	switch lexeme {
	case token.SEMICOLON.String():
		return token.SEMICOLON
	case token.LBRACE.String():
		return token.LBRACE
	case token.RBRACE.String():
		return token.RBRACE
	case token.SCHEMA.String():
		return token.SCHEMA
	case token.INT.String():
		return token.INT
	case token.STRING.String():
		return token.STRING
	default:
		return token.NewIdentifier(lexeme)
	}
}
