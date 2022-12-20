package _go

import (
	"strings"

	cueopenapi "cuelang.org/go/encoding/openapi"
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/thema/encoding/gocode"
	"github.com/grafana/thema/encoding/openapi"
)

// GoTypesJenny creates a [OneToOne] that produces Go types for the provided
// [thema.Schema].
//
// Copied/hacked out of grafana/pkg/codegen for now to deal with ExpandReferences situation
type GoTypesJenny struct{}

func (j GoTypesJenny) JennyName() string {
	return "GoTypesJenny"
}

func (j GoTypesJenny) Generate(sfg codegen.SchemaForGen) (*codejen.File, error) {
	// TODO allow using name instead of machine name in thema generator
	b, err := gocode.GenerateTypesOpenAPI(sfg.Schema, &gocode.TypeConfigOpenAPI{
		PackageName: strings.ToLower(sfg.Schema.Lineage().Name()),
		Config: &openapi.Config{
			Config:   &cueopenapi.Config{
				// ExpandReferences: true,
			},
			Group:    sfg.IsGroup,
		},
	})
	if err != nil {
		return nil, err
	}

	return codejen.NewFile(sfg.Schema.Lineage().Name()+"_types_gen.go", b, j), nil
}
