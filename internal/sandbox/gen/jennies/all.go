package jennies

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/ast/compiler"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/golang"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/typescript"
)

type LanguageTarget struct {
	Jennies        *codejen.JennyList[[]*ast.File]
	CompilerPasses []compiler.Pass
}

func All() map[string]LanguageTarget {
	targets := map[string]LanguageTarget{
		// Compiler passes should not have side effects, but they do.
		"go": {
			Jennies: golang.Jennies(),
			CompilerPasses: []compiler.Pass{
				&compiler.AnonymousEnumToExplicitType{},
				&compiler.DisjunctionToType{},
			},
		},
		"typescript": {
			Jennies:        typescript.Jennies(),
			CompilerPasses: nil,
		},
	}

	return targets
}
