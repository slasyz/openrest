package swagger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/slasyz/openrest/internal/parsers/schema"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name string

		document   string
		outPackage string

		check func(parsed *Parsed, err error)
	}{
		{
			name: "single path example",

			document: `openapi: 3.0.0

info:
  title: Test File
  version: test

paths:
  /auth/register:
    post:
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                login:
                  type: string
                password:
                  type: string
              required:
                - login
                - password
      responses:
        200:
          description: Success.
          content:
            application/json:
              schema:
                type: object
                properties:
                  token:
                    type: string
`,
			outPackage: "out/package",

			check: func(parsed *Parsed, err error) {
				assert.NoError(t, err)
				// TODO
				return
			},
		},
	}

	schm := schema.New()
	swgr := New(schm)

	for _, tt := range tests {
		var document Document
		err := yaml.Unmarshal([]byte(tt.document), &document)
		assert.NoError(t, err, tt.name)

		parsed, err := swgr.Parse(&document)
		tt.check(parsed, err)
	}
}
