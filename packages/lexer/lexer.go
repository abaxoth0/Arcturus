package lexer

import (
	"arcturus/packages/shared/logger"
	"arcturus/packages/token"

	log "github.com/abaxoth0/Ain/logger"
)

var lexerLogger = log.NewSource("LEXER", logger.Default)

func Parse(input []byte) ([]token.Token, error) {
	tokens := []token.Token{}
	reading := false
	lexeme := make([]byte, 0, 256)

	for _, char := range input {
		switch char {
		case '{', '}', ';', ' ':
			if reading {
				token, err := tokenize(string(lexeme))
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, token)
				lexeme = lexeme[:0] // set length to 0 keeping capacity to avoid buffer reallocation
				reading = false
			}
			if char != ' ' {
				token, err := tokenize(string(char))
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, token)
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
	return tokens, nil
}

func tokenize(lexeme string) (token.Token, error) {
	lexerLogger.Trace("lexeme: "+lexeme, nil)
	switch lexeme {
	case token.SEMICOLON.String():
		return token.SEMICOLON, nil
	case token.LBRACE.String():
		return token.LBRACE, nil
	case token.RBRACE.String():
		return token.RBRACE, nil
	case token.SCHEMA.String():
		return token.SCHEMA, nil
	case token.INT.String():
		return token.INT, nil
	case token.STRING.String():
		return token.STRING, nil
	default:
		return token.NewIdentifier(lexeme)
	}
}
