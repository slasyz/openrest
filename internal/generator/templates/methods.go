package templates

import "text/template"

var TplMethods = template.Must(template.New("").Parse(`package methods

type Methods struct {
	// TODO: you can define all dependencies if you need
}

func New() *Methods {
	return &Methods{}
}
`))
