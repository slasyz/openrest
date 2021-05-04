package schema

import (
	"errors"
	"fmt"
	"strings"

	"github.com/slasyz/openrest/internal/parsers"
)

var (
	errEmptyPrefix = errors.New("prefix cannot be empty")
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

// Parse reads defined schema and returns its golang type name with list of structs that need to be defined.
func (p *Parser) Parse(schema *Schema, prefix string) (string, []parsers.Struct, error) {
	if prefix == "" {
		return "", nil, errEmptyPrefix
	}

	switch schema.Type {
	case "string":
		return "string", nil, nil
	case "integer":
		return "int64", nil, nil
	case "object":
		var fields []parsers.Field
		var structs []parsers.Struct

		for _, prop := range schema.Properties.List {
			nameCapitalized := strings.ToUpper(string(prop.Key[0])) + prop.Key[1:]

			propType, propStructs, err := p.Parse(prop.Value, prefix+nameCapitalized)
			if err != nil {
				return "", nil, fmt.Errorf("error parsing %s%s: %w", prefix, nameCapitalized, err)
			}

			fields = append(fields, parsers.Field{
				Name: nameCapitalized,
				Type: propType,
			})

			structs = append(structs, propStructs...)
		}

		structs = append([]parsers.Struct{
			{
				Name:   prefix,
				Fields: fields,
			},
		}, structs...)

		return prefix, structs, nil
	case "array":
		itemType, itemStructs, err := p.Parse(schema.Items, prefix+"Item")
		if err != nil {
			return "", nil, fmt.Errorf("error parsing %s item: %w", prefix, err)
		}

		return "[]" + itemType, itemStructs, nil
	default:
		return "", nil, fmt.Errorf("unexpected type: %s", schema.Type)
	}
}
