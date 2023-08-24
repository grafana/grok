package jennies

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/golang"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/typescript"
)

func All(outputPrefix string) map[string]*codejen.JennyList[*ast.File] {
	targets := map[string]*codejen.JennyList[*ast.File]{
		"go":         golang.Jennies(outputPrefix),
		"typescript": typescript.Jennies(outputPrefix),
	}

	return targets
}
