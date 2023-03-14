package cuetf

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToCamelCase(str string) string {
	words := strings.Split(str, "_")
	camelCase := ""
	for _, s := range words {
		camelCase += strings.Title(s)
	}
	return camelCase
}

func GetKindName(rawName string) string {
	name := rawName
	if strings.HasSuffix(name, "PanelCfg") {
		name = "Panel" + strings.TrimSuffix(name, "PanelCfg")
	} else if strings.HasSuffix(name, "DataQuery") {
		name = "Query" + strings.TrimSuffix(name, "DataQuery")
	} else {
		switch name {
		case "dashboard", "playlist", "preferences", "team":
			name = strings.ToUpper(name[:1]) + name[1:]
		case "publicdashboard":
			name = "PublicDashboard"
		case "librarypanel":
			name = "LibraryPanel"
		case "serviceaccount":
			name = "ServiceAccount"
		}
		name = "Core" + name
	}

	return name
}

func GetStructName(rawName string) string {
	return strings.Title(GetKindName(rawName)) + "DataSource"
}

func GetResourceName(rawName string) string {
	return ToSnakeCase(GetKindName(rawName))
}
