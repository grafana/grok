package golang

import (
	"embed"
	"html/template"
)

var templates *template.Template

//go:embed veneers/*.tmpl
var veneersFS embed.FS

func init() {
	base := template.New("golang")
	base.Funcs(map[string]any{
		"formatIdentifier": formatIdentifier,
	})
	templates = template.Must(base.ParseFS(veneersFS, "veneers/*.tmpl"))
}
