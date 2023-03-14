package types

import "cuelang.org/go/cue"

type Node struct {
	Name     string
	Kind     cue.Kind
	SubKind  cue.Kind // For list only, kind of its elements
	Optional bool
	Default  string
	Doc      string
	Children []Node
}
