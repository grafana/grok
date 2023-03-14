package cuetf

import (
	"embed"
	"strings"
	"text/template"
)

// All the parsed templates in the tmpl subdirectory
var tmpls *template.Template

func init() {
	base := template.New("cuetf").Funcs(template.FuncMap{
		"lowerCase": strings.ToLower,
	})
	tmpls = template.Must(base.ParseFS(tmplFS, "templates/*.tmpl"))
}

//go:embed templates/*.tmpl
var tmplFS embed.FS

// The following group of types, beginning with tvars_*, all contain the set
// of variables expected by the corresponding named template file under templates/
type (
	TVarsDataSource struct {
		Name              string
		StructName        string
		Description       string
		ModelFields       string
		JSONModelFields   string
		TFModelToJSONFunc string
		SchemaAttributes  string
		Defaults          string
	}

	TVarsSchemaAttribute struct {
		Name                   string
		Description            string
		AttributeType          string
		Computed               bool
		Optional               bool
		Required               bool
		ElementType            string // Used for simple lists
		NestedAttributes       string // Used for objects
		NestedObjectAttributes string // Used for complex lists
	}

	TVarsDefault struct {
		Name               string
		NullFieldCondition string
		Type               string
		Default            string
	}
)
