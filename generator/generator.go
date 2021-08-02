package generator

import (
	"bytes"
	"graphql-to-go/entity"
	"log"
	"text/template"
)

const StructTemplate = `
type {{exportName .Name}} struct {
    {{range .Fields}} {{exportName  .Name}} {{fieldType .Type}} {{jsonTag .Name}}
    {{end}}
}
`

var tStruct = template.Must(template.New("struct").Funcs(functions).Parse(StructTemplate))

// ParseSchema Create Go code from Grapql schema
func ParseSchema(data entity.ResponseDataSchema) (map[string]string, error) {
	structs := make(map[string]string)

	for _, t := range data.Schema.Types {
		result, err := generateGoStruct(t)
		if err != nil {
			log.Fatal(err)
		}
		structs[t.Name] = *result
	}

	// Inject fake type
	result, err := generateGoStruct(entity.Type{Name: "FAKENotFoundType", Kind: "OBJECT"})
	if err != nil {
		log.Fatal(err)
	}
	structs[`FAKENotFoundType`] = *result

	return structs, nil
}

// ParseType Create Go code from Grapql type schema
func ParseType(data entity.ResponseDataType) (result *string, err error) {
	return generateGoStruct(data.Type)
}

func generateGoStruct(t entity.Type) (*string, error) {
	var buffer = new(bytes.Buffer)

	if err := tStruct.Execute(buffer, t); err != nil {
		return nil, err
	}
	result := buffer.String()
	return &result, nil
}
