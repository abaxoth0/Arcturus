package token

type Type uint8

const (
	Plain = iota
	Value
)

type Token interface {
	Type() Type
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

	INT
	STRING

	WHITESPACE
)

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

type IDENTIFIER valueToken[string]

func NewIdentifier(name string) IDENTIFIER {
	return IDENTIFIER{
		Value: name,
		valueKind: identifier,
	}
}
