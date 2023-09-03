package compiler

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type Pass interface {
	Process(files []*ast.File) ([]*ast.File, error)
}
