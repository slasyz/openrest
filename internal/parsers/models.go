package parsers

type Field struct {
	Name string
	Type string
}

type Struct struct {
	Name   string
	Fields []Field
}

type Method struct {
	File  string
	Name  string
	Path  string
	InDTO string
}
