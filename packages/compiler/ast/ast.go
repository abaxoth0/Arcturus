package ast

import (
	"arcturus/packages/compiler/types"
	"arcturus/packages/shared/logger"
	"maps"

	log "github.com/abaxoth0/Ain/logger"
)

var astLogger = log.NewSource("AST", logger.Default)

type Module struct {
	Name 	string
	Schemas []*Schema
}

var modules map[string]*Module = map[string]*Module{}

type Node interface {
	Name() string
	Type() types.Type
}

type Schema struct {
	name string
	props map[string]types.Type
}

func NewSchema(name string) *Schema {
	return &Schema{
		name: name,
		props: make(map[string]types.Type),
	}
}

func (s *Schema) Name() string {
	return s.name
}

func (s *Schema) Type() types.Type {
	return types.Schema
}

func (s *Schema) Properties() map[string]types.Type {
	return maps.Clone(s.props)
}
