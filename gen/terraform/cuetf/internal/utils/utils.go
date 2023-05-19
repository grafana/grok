package utils

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

func CapitalizeFirstLetter(str string) string {
	sep := " "
	parts := strings.SplitN(str, sep, 2)
	if len(parts) != 2 {
		return strings.Title(str)
	}
	return strings.Title(parts[0]) + sep + parts[1]
}
