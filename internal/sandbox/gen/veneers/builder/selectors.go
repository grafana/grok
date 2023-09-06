package builder

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type Selector func(builder ast.Builder) bool

func ExactBuilder(objectName string) Selector {
	return func(builder ast.Builder) bool {
		return builder.For.Name == objectName
	}
}

func EveryBuilder() Selector {
	return func(builder ast.Builder) bool {
		return true
	}
}
