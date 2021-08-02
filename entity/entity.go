package entity

type (
	Kind int

	Type struct {
		Name   string  `json:"name"`
		Kind   string  `json:"kind"`
		Fields []Field `json:"fields"`
		OfType OfType  `json:"ofType"`
	}

	OfType struct {
		Name   string  `json:"name"`
		Kind   string  `json:"kind"`
		Fields []Field `json:"fields"`
	}

	Field struct {
		Name string `json:"name"`
		Type Type   `json:"type"`
		Args []Arg  `json:"args"`
	}

	Arg struct {
		Name         string `json:"name"`
		DefaultValue string `json:"defaultValue"`
		Type         Type   `json:"type"`
	}

	Schema struct {
		Types []Type `json:"types"`
	}

	ResponseDataSchema struct {
		Schema Schema `json:"__schema"`
	}

	ResponseDataType struct {
		Type Type `json:"__type"`
	}
)

var ScalarToGo = map[string]string{
	`String`:  `string`,
	`Int`:     `int`,
	`Float`:   `float32`,
	`Boolean`: `bool`,
	`ID`:      `string`,
	`Time`:    `time.Time`,
}

// Kind enumeration
const (
	ENUM Kind = iota + 1
	INTERFACE
	LIST
	NON_NULL
	OBJECT
	SCALAR
	UNION
)

var kindText = map[Kind]string{
	ENUM:      "ENUM",
	INTERFACE: "INTERFACE",
	LIST:      "LIST",
	NON_NULL:  "NON_NULL",
	OBJECT:    "OBJECT",
	SCALAR:    "SCALAR",
	UNION:     "UNION",
}

func KindText(kind Kind) string {
	return kindText[kind]
}

func KindCode(kind string) Kind {
	for k, v := range kindText {
		if v == kind {
			return k
		}
	}
	return 0
}
