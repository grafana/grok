package internal

import (
	"errors"
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/grok/gen/terraform/cuetf/types"
)

func GetAllNodes(val cue.Value) ([]types.Node, error) {
	if err := val.Validate(); err != nil {
		return nil, fmt.Errorf("error validating value: %w", err)
	}

	iter, err := val.Fields(
		cue.Definitions(false), // Should we do something with those?
		cue.Optional(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving value fields: %w", err)
	}

	nodes := make([]types.Node, 0)
	for iter.Next() {
		node, err := GetSingleNode(iter.Selector().String(), iter.Value(), iter.IsOptional())
		if err != nil {
			return nil, err
		}

		if node != nil {
			nodes = append(nodes, *node)
		}
	}

	return nodes, nil
}

func GetSingleNode(name string, val cue.Value, optional bool) (*types.Node, error) {
	// TODO: fixme
	if name == "mappings" || name == "points" || name == "bucketAggs" || name == "metrics" {
		return nil, nil
	}

	node := types.Node{
		Name:     name,
		Kind:     val.IncompleteKind(),
		Optional: optional,
		Default:  GetDefault(val),
	}

	for _, comment := range val.Doc() {
		node.Doc += comment.Text()
	}
	node.Doc = strings.ReplaceAll(strings.Trim(node.Doc, "\n "), "`", "")

	switch node.Kind {
	case cue.ListKind:
		// From cuetsy:
		// If the default (all lists have a default, usually self, ugh) differs from the
		// input list, peel it off. Otherwise our AnyIndex lookup may end up getting
		// sent on the wrong path.
		defv, _ := val.Default()
		if !defv.Equals(val) {
			_, v := val.Expr()
			val = v[0]
		}

		e := val.LookupPath(cue.MakePath(cue.AnyIndex))
		if e.Exists() {
			node.SubKind = e.IncompleteKind()
			// TODO: fixme
			// Using a string type to allow composition of panel datasources
			// Doesn't seem possible to have an arbitrary map type here
			if name == "panels" {
				node.SubKind = cue.StringKind
			}

			if node.SubKind == cue.StructKind || node.SubKind == cue.ListKind {
				children, err := GetAllNodes(e)
				if err != nil {
					return nil, err
				}

				node.Children = children
			}
		} else {
			return nil, errors.New("unreachable - open list must have a type")
		}
	case cue.StructKind:
		children, err := GetAllNodes(val.Value())
		if err != nil {
			return nil, err
		}

		node.Children = children
	}

	return &node, nil
}

func GetDefault(v cue.Value) string {
	_, ok := v.Default()
	if !ok {
		return ""
	}

	switch v.IncompleteKind() {
	case cue.StringKind:
		s, err := v.String()
		if err != nil {
			return ""
		}
		return fmt.Sprintf("`%s`", s)
	case cue.FloatKind:
		f, err := v.Float64()
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%f", f)
	case cue.IntKind:
		i, err := v.Int64()
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%d", i)
	case cue.NumberKind:
		i, err := v.Float64()
		if err != nil {
			return ""
		}
		return fmt.Sprintf("new(big.Float).SetFloat64(%f)", i)
	case cue.BoolKind:
		b, err := v.Bool()
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%t", b)
	default:
		return ""
	}
}
