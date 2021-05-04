package swagger

import (
	"fmt"

	"github.com/slasyz/openrest/internal/names"
	"github.com/slasyz/openrest/internal/parsers"
	"github.com/slasyz/openrest/internal/parsers/schema"
)

type Parser struct {
	schemaParser *schema.Parser
}

func New(schemaParser *schema.Parser) *Parser {
	return &Parser{
		schemaParser: schemaParser,
	}
}

type Parsed struct {
	Structs []parsers.Struct
	Methods []parsers.Method
}

func (p *Parser) Parse(document *Document) (*Parsed, error) {
	var structs []parsers.Struct
	var methods []parsers.Method

	for _, path := range document.Paths.List {
		if path.Value == nil {
			return nil, fmt.Errorf("no value in path")
		}
		if path.Value.Post == nil {
			return nil, fmt.Errorf("expecting post request description in %s", path.Key)
		}
		if path.Value.Post.RequestBody == nil {
			return nil, fmt.Errorf("expecting requestBody in %s", path.Key)
		}
		if len(path.Value.Post.RequestBody.Content) != 1 {
			return nil, fmt.Errorf("expecting one request body in %s", path.Key)
		}

		method := parsers.Method{
			File: names.PathToSnakeCase(path.Key) + ".go",
			Name: names.PathToPascalCase(path.Key),
			Path: path.Key,
		}

		for ct, val := range path.Value.Post.RequestBody.Content {
			if ct != "application/json" {
				return nil, fmt.Errorf("unexpected Content-Type: %s", ct)
			}

			var inStructs []parsers.Struct
			var err error
			method.InDTO, inStructs, err = p.schemaParser.Parse(val.Schema, method.Name+"In")
			if err != nil {
				return nil, fmt.Errorf("cannot parse %s %s request body: %w", path.Key, ct, err)
			}

			structs = append(structs, inStructs...)
		}

		for code, val := range path.Value.Post.Responses {
			if len(val.Content) != 1 {
				return nil, fmt.Errorf("expecting one response body in %s %s", path.Key, code)
			}

			for ct, val := range val.Content {
				if ct != "application/json" {
					return nil, fmt.Errorf("unexpected Content-Type: %s", ct)
				}

				_, outStructs, err := p.schemaParser.Parse(val.Schema, method.Name+names.Capitalize(code))
				if err != nil {
					return nil, fmt.Errorf("cannot parse %s %s %s response: %w", path.Key, code, ct, err)
				}

				structs = append(structs, outStructs...)
			}
		}

		methods = append(methods, method)
	}

	return &Parsed{
		Structs: structs,
		Methods: methods,
	}, nil
}
