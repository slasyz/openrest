package templates

import "text/template"

var TplStub = template.Must(template.New("").Parse(`package methods

import (
	"errors"
	"net/http"
	
	"{{.OutputPackage}}/generated"
)

{{with .Method}}
func (m *Methods) {{.Name}}(r *http.Request, in *generated.{{.Name}}In) (int, interface{}, error) {
    return 0, nil, errors.New("implement me")
}
{{end}}`))
