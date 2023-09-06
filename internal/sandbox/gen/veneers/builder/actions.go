package builder

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type RewriteAction func(builder ast.Builder) ast.Builder

func OmitAction() RewriteAction {
	return func(_ ast.Builder) ast.Builder {
		return ast.Builder{}
	}
}
