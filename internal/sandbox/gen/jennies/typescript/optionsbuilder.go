package typescript

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type TypescriptOptionsBuilder struct {
}

func (jenny TypescriptOptionsBuilder) JennyName() string {
	return "TypescriptOptionsBuilder"
}

func (jenny TypescriptOptionsBuilder) Generate(file *ast.File) (*codejen.File, error) {
	output := jenny.generateFile(file)

	return codejen.NewFile("options_builder_gen.ts", []byte(output), jenny), nil
}

func (jenny TypescriptOptionsBuilder) generateFile(file *ast.File) string {
	return `export interface OptionsBuilder<T> {
  build: () => T;
}
`
}
