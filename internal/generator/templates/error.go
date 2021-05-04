package templates

import "text/template"

var TplError = template.Must(template.New("").Parse(`package methods

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

type UserError error

func (m *Methods) Error(err error, w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())
	logger.Err(err)

	errText := "Internal server error"
	if userError, ok := err.(UserError); ok {
		errText = userError.Error()
	}

	errData, err := json.Marshal(struct {
		Error string ` + "`" + `json:"error"` + "`" + `
	}{
		Error: errText,
	})
	if err != nil {
		logger.Err(fmt.Errorf("cannot create error response: %w", err))
		return
	}

	w.Write(errData)
}
`))
