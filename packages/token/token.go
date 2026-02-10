package token

import (
	"errors"
	"fmt"
	"strings"
)

type Type uint8

const (
	Plain = iota
	Value
)

type Kind uint

const (
	None Kind = iota
	DataType = iota
	Separator
	Declaration

	Identifier
)

type Token interface {
	Kind() Kind
	Type() Type
	String() string
}

type PlainToken uint

func (_ PlainToken) Type() Type {
	return Plain
}

const (
	NONE PlainToken = iota

	SEMICOLON
	LBRACE
	RBRACE

	SCHEMA

	INT
	STRING
)

func (t PlainToken) String() string {
	switch t {
	case SEMICOLON:
		return ";"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case SCHEMA:
		return "schema"
	case INT:
		return "int"
	case STRING:
		return "string"
	default:
		return "none"
	}
}

func (t PlainToken) Kind() Kind {
	switch t {
	case SEMICOLON, LBRACE, RBRACE:
		return Separator
	case INT, STRING:
		return DataType
	case SCHEMA:
		return Declaration
	default:
		return None
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

func (t valueToken[T]) String() string {
	switch v := any(t).(type) {
	case valueToken[string]:
		return v.Value
	}
	if str, ok := any(t).(fmt.Stringer); ok {
		return str.String()
	}
	return "VALUE_TOKEN"
}

func (t valueToken[T]) Kind() Kind {
	switch t.valueKind {
	case identifier:
		return Identifier
	default:
		return None
	}
}

type IDENTIFIER = valueToken[string]

var ErrEmptyIdentifierName = errors.New("Identifier name is empty")

func NewIdentifier(name string) (IDENTIFIER, error) {
	if strings.ReplaceAll(name, " ", "") == "" {
		var zero IDENTIFIER
		return zero, ErrEmptyIdentifierName
	}
	return IDENTIFIER{
		Value:     name,
		valueKind: identifier,
	}, nil
}
