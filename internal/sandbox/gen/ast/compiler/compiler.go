package compiler

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type Pass interface {
	Process(file *ast.File) (*ast.File, error)
}
