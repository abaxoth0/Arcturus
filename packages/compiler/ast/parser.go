package ast

import (
	"arcturus/packages/compiler/types"
	"arcturus/packages/token"
	"errors"
	"fmt"
)

var keywords map[string]token.PlainToken = map[string]token.PlainToken{
	token.SCHEMA.String(): token.SCHEMA,
	token.INT.String():    token.INT,
	token.STRING.String(): token.STRING,
}

func ParseModule(name string, tokens []token.Token) (*Module, error) {
	if modules[name] != nil {
		return nil, errors.New("Module \""+name+"\" already exists")
	}

	const ( // states
		none = iota
		expectId
		expectSemicolon
	)
	var (
		schema  *Schema // TODO replace with stack to track current scope
		prev    token.Token
		module  = &Module{Name: name}
		inBlock = false // TODO replace with stack to track current scope
		state   = none
		ids     = map[string]bool{}
	)

	for i, tk := range tokens {
		if state == expectSemicolon && tk != token.SEMICOLON {
			return nil, errors.New("Semicolon expected, but got "+tk.String())
		}

		switch tk {
		case token.SEMICOLON:
			if state == expectSemicolon {
				astLogger.Trace("State change: none", nil)
				state = none
				continue
			}
			astLogger.Trace(prev.String(), nil)
			astLogger.Trace(tk.String(), nil)
			return nil, errors.New("Unexpected ';'")
		case token.LBRACE:
			if inBlock {
				return nil, errors.New("Can't nest blocks")
			}
			inBlock = true
		case token.RBRACE:
			if !inBlock {
				return nil, errors.New("Unexpected '}': Block wasn't open")
			}
			inBlock = false
		case token.SCHEMA:
		case token.INT, token.STRING:
			if !inBlock {
				return nil, fmt.Errorf("Unexpected '%s': Out of schema property declaration", tk.String())
			}
		}

		if id, ok := tk.(token.IDENTIFIER); ok {
			astLogger.Debug("New identifier \""+id.Value+"\": "+prev.String(), nil)
			if tk.Kind() == token.Identifier && prev != token.SCHEMA {
				state = expectSemicolon
				astLogger.Trace("State change: expectSemicolon", nil)
			}
			if keywords[id.Value] != token.NONE {
				return nil, errors.New("Can't use keyword as an identifier name")
			}
			if ids[id.Value] {
				return nil, errors.New("Identifier duplication")
			}
			switch prev {
			case token.SCHEMA:
				schema = NewSchema(id.Value)
				module.Schemas = append(module.Schemas, schema)
				ids[id.Value] = true
			case token.STRING, token.INT:
				idType, err := types.ConvertToken(prev)
				if err != nil {
					return nil, err
				}
				schema.props[id.Value] = idType
			default:
				astLogger.Trace(prev.String(), nil)
				astLogger.Trace(tk.String(), nil)
				return nil, errors.New("Invalid identifier declaration")
			}
		}

		if len(tokens)-1 == i && inBlock && tk != token.RBRACE {
			return nil, errors.New("Expecting '}', but got '" + tk.String() + "'")
		}

		prev = tk
	}

	return module, nil
}
