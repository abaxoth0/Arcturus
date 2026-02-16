package types

import (
	"arcturus/packages/token"
	"errors"
)

type Type uint8

const (
	None Type = iota
	Schema
	String
	Int32
)

func (t Type) String() string {
	switch t {
	case Schema:
		return "schema"
	case String:
		return "string"
	case Int32:
		return "int32"
	default:
		return ""
	}
}

func ConvertToken(tk token.Token) (Type, error) {
	if tk.Kind() != token.DataType {
		return None, errors.New("Not a data type")
	}
	switch tk {
	case token.SCHEMA:
		return Schema, nil
	case token.STRING:
		return String, nil
	case token.INT:
		return Int32, nil
	default:
		return None, errors.New("Invalid token: "+tk.String())
	}
}
