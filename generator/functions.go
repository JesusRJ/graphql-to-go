package generator

import (
	"graphql-to-go/entity"
	"strings"
	"text/template"
)

var functions = template.FuncMap{
	"exportName": exportName,
	"fieldType":  fieldType,
	"jsonTag":    jsonTag,
}

var exportName = func(name string) string {
	return strings.Title(name)
}

var fieldType = func(t entity.Type) string {
	var typeName = t.Name

	if typeName == "" {
		if t.OfType.Name != "" {
			typeName = t.OfType.Name
		} else {
			return `FAKENotFoundType`
		}
	}

	if v, ok := entity.ScalarToGo[typeName]; ok {
		return v
	}

	typeName = strings.Title(typeName)

	if t.Kind == entity.KindText(entity.LIST) {
		typeName = `[]` + typeName
	}

	return typeName
}

var jsonTag = func(name string) string {
	return "`json:\"" + name + "\"`"
}
