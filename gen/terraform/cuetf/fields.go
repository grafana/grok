package cuetf

import (
	"fmt"

	"cuelang.org/go/cue"
)

type cueField struct {
	Name       string
	Value      cue.Value
	IsOptional bool
}

func schemaToCueFields(schema cue.Value) ([]cueField, error) {
	if !schema.IsConcrete() {
		return nil, nil
	}

	fields := []cueField{}
	iter, err := schema.Fields(
		cue.Definitions(false),
		cue.Optional(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving value fields: %w", err)
	}
	for iter.Next() {
		fields = append(fields, cueField{
			Name:       iter.Selector().String(),
			Value:      iter.Value(),
			IsOptional: iter.IsOptional(),
		})
	}
	return fields, nil
}
