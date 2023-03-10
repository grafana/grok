package terraform

import (
	"fmt"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/gen/terraform/cuetf"
)

type TerraformDataSourceJenny struct{}

func (j TerraformDataSourceJenny) JennyName() string {
	return "TerraformDataSourceJenny"
}

func (j TerraformDataSourceJenny) Generate(sfg codegen.SchemaForGen) (*codejen.File, error) {
	// TODO allow using name instead of machine name in thema generator
	bytes, err := cuetf.GenerateDataSource(sfg.Schema)
	if err != nil {
		return nil, err
	}

	name := sfg.Schema.Lineage().Name()
	fmt.Println(sfg.Schema.Lineage().Name())

	return codejen.NewFile("datasource_"+name+".go", bytes, j), nil
}
