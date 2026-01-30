package token

import (
	"fmt"
	"strings"
)

type Type uint8

const (
	Plain = iota
	Value
)

type Token interface {
	Type() Type
	Raw()  string
}

type PlainToken uint

func (_ PlainToken) Type() Type {
	return Plain
}

const (
	none PlainToken = iota

	SEMICOLON
	LBRACE
	RBRACE

	MESSAGE

	INT
	STRING

	WHITESPACE
)

func (t PlainToken) Raw() string {
	switch t {
	case SEMICOLON:
		return ";"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case WHITESPACE:
		return " "
	case MESSAGE:
		return "message"
	case INT:
		return "int"
	case STRING:
		return "string"
	default:
		return "none"
	}
}

type valueKind uint

const (
	identifier valueKind = iota
)

type valueToken[T any] struct {
	Value T
	valueKind

	PlainToken
}

func (_ valueToken[T]) Type() Type {
	return Value
}

func (t valueToken[T]) Raw() string {
	switch v := any(t).(type) {
	case valueToken[string]:
		return v.Value
	}
	if str, ok := any(t).(fmt.Stringer); ok {
		return str.String()
	}
	return "VALUE_TOKEN"
}

type IDENTIFIER = valueToken[string]

func NewIdentifier(name string) IDENTIFIER {
	if strings.ReplaceAll(name, " ", "") == "" {
		panic("Empty IDENTIFIER")
	}
	return IDENTIFIER{
		Value: name,
		valueKind: identifier,
	}
}
