package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/slasyz/openrest/internal/parsers"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name string

		prefix string
		schema *Schema

		expectedType    string
		expectedStructs []parsers.Struct
	}{
		{
			name: "string example",

			prefix: "TestBla",
			schema: &Schema{
				Type: "string",
			},

			expectedType:    "string",
			expectedStructs: nil,
		},
		{
			name: "object example",

			prefix: "TestBla",
			schema: &Schema{
				Type: "object",
				Properties: Properties{List: []Property{
					{
						Key: "prop1",
						Value: &Schema{
							Type: "string",
						},
					},
					{
						Key: "prop2",
						Value: &Schema{
							Type: "integer",
						},
					},
				}},
			},

			expectedType: "TestBla",
			expectedStructs: []parsers.Struct{
				{
					Name: "TestBla",
					Fields: []parsers.Field{
						{
							Name: "Prop1",
							Type: "string",
						},
						{
							Name: "Prop2",
							Type: "int64",
						},
					},
				},
			},
		},
		{
			name: "nested object example",

			prefix: "TestBla",
			schema: &Schema{
				Type: "object",
				Properties: Properties{List: []Property{
					{
						Key: "prop1",
						Value: &Schema{
							Type: "object",
							Properties: Properties{List: []Property{
								{
									Key: "prop1a",
									Value: &Schema{
										Type: "object",
										Properties: Properties{List: []Property{
											{
												Key: "prop1a1",
												Value: &Schema{
													Type: "integer",
												},
											},
											{
												Key: "prop1a2",
												Value: &Schema{
													Type: "object",
													Properties: Properties{List: []Property{
														{
															Key: "prop1a2i",
															Value: &Schema{
																Type: "integer",
															},
														},
													}},
												},
											},
										}},
									},
								},
								{
									Key: "prop1b",
									Value: &Schema{
										Type: "object",
										Properties: Properties{List: []Property{
											{
												Key: "prop1b1",
												Value: &Schema{
													Type: "string",
												},
											},
											{
												Key: "prop1b2",
												Value: &Schema{
													Type: "string",
												},
											},
										}},
									},
								},
							}},
						},
					},
					{
						Key: "prop2",
						Value: &Schema{
							Type: "integer",
						},
					},
				}},
			},

			expectedType: "TestBla",
			expectedStructs: []parsers.Struct{
				{
					Name: "TestBla",
					Fields: []parsers.Field{
						{
							Name: "Prop1",
							Type: "TestBlaProp1",
						},
						{
							Name: "Prop2",
							Type: "int64",
						},
					},
				},
				{
					Name: "TestBlaProp1",
					Fields: []parsers.Field{
						{
							Name: "Prop1a",
							Type: "TestBlaProp1Prop1a",
						},
						{
							Name: "Prop1b",
							Type: "TestBlaProp1Prop1b",
						},
					},
				},
				{
					Name: "TestBlaProp1Prop1a",
					Fields: []parsers.Field{
						{
							Name: "Prop1a1",
							Type: "int64",
						},
						{
							Name: "Prop1a2",
							Type: "TestBlaProp1Prop1aProp1a2",
						},
					},
				},
				{
					Name: "TestBlaProp1Prop1aProp1a2",
					Fields: []parsers.Field{
						{
							Name: "Prop1a2i",
							Type: "int64",
						},
					},
				},
				{
					Name: "TestBlaProp1Prop1b",
					Fields: []parsers.Field{
						{
							Name: "Prop1b1",
							Type: "string",
						},
						{
							Name: "Prop1b2",
							Type: "string",
						},
					},
				},
			},
		},
		{
			name: "array example",

			prefix: "TestBla",
			schema: &Schema{
				Type: "array",
				Items: &Schema{
					Type: "object",
					Properties: Properties{List: []Property{
						{
							Key: "prop1",
							Value: &Schema{
								Type: "string",
							},
						},
					}},
				},
			},

			expectedType: "[]TestBlaItem",
			expectedStructs: []parsers.Struct{
				{
					Name: "TestBlaItem",
					Fields: []parsers.Field{
						{
							Name: "Prop1",
							Type: "string",
						},
					},
				},
			},
		},
	}

	schema := New()

	for _, tt := range tests {
		typeName, structs, err := schema.Parse(tt.schema, tt.prefix)
		assert.NoError(t, err, tt.name)

		assert.Equal(t, tt.expectedType, typeName, tt.name)
		assert.ElementsMatch(t, tt.expectedStructs, structs, tt.name)
	}
}

func TestParser_Parse_FromYaml(t *testing.T) {
	tests := []struct {
		name string

		prefix string
		schema string

		expectedType    string
		expectedStructs []parsers.Struct
	}{
		{
			name: "string example",

			prefix: "TestBla",
			schema: `type: string`,

			expectedType:    "string",
			expectedStructs: nil,
		},
		{
			name: "object example",

			prefix: "TestBla",
			schema: `
              type: object
              properties:
                prop1:
                  type: string
                prop2:
                  type: integer
`,

			expectedType: "TestBla",
			expectedStructs: []parsers.Struct{
				{
					Name: "TestBla",
					Fields: []parsers.Field{
						{
							Name: "Prop1",
							Type: "string",
						},
						{
							Name: "Prop2",
							Type: "int64",
						},
					},
				},
			},
		},
	}

	schema := New()

	for _, tt := range tests {
		var s Schema
		err := yaml.Unmarshal([]byte(tt.schema), &s)
		assert.NoError(t, err, tt.name)

		typeName, structs, err := schema.Parse(&s, tt.prefix)
		assert.NoError(t, err, tt.name)

		assert.Equal(t, tt.expectedType, typeName, tt.name)
		assert.ElementsMatch(t, tt.expectedStructs, structs, tt.name)
	}
}
