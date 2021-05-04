package swagger

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/slasyz/openrest/internal/parsers/schema"
)

type Document struct {
	Paths Paths `yaml:"paths"`
}

type Paths struct {
	List []Path
}

func (p *Paths) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := make(map[string]*PathItem)
	var ms yaml.MapSlice

	// The only way I came up to unmarshal map with saving keys order and value type information.

	err := unmarshal(&m)
	if err != nil {
		return fmt.Errorf("cannot unmarshal to standard map: %w", err)
	}

	err = unmarshal(&ms)
	if err != nil {
		return fmt.Errorf("cannot unmarshal to yaml.MapSlice: %w", err)
	}

	for _, msEl := range ms {
		key, ok := msEl.Key.(string)
		if !ok {
			return fmt.Errorf("all keys must be strings")
		}

		val, ok := m[key]
		if !ok {
			return fmt.Errorf("%s not found in standard map", key)
		}

		p.List = append(p.List, Path{
			Key:   key,
			Value: val,
		})
	}

	return nil
}

type Path struct {
	Key   string
	Value *PathItem
}

type PathItem struct {
	Summary string     `yaml:"summary"`
	Get     *Operation `yaml:"get"`
	Put     *Operation `yaml:"put"`
	Post    *Operation `yaml:"post"`
	Delete  *Operation `yaml:"delete"`
	Options *Operation `yaml:"options"`
	Head    *Operation `yaml:"head"`
	Patch   *Operation `yaml:"patch"`
	Trace   *Operation `yaml:"trace"`
	//Parameters []*Parameter `yaml:"parameters"`
}

type Operation struct {
	Summary string `yaml:"summary"`
	//Parameters  []*Parameter         `yaml:"parameters"`
	RequestBody *RequestBody         `yaml:"requestBody"`
	Responses   map[string]*Response `yaml:"responses"` // key is "default" or HTTP code
}

//type ParameterIn string
//
//const (
//	ParameterInQuery  ParameterIn = "query"
//	ParameterInHeader ParameterIn = "header"
//	ParameterInPath   ParameterIn = "path"
//	ParameterInCookie ParameterIn = "cookie"
//)
//
//type Parameter struct {
//	Name     string      `yaml:"name"`
//	In       ParameterIn `yaml:"in"`
//	Required bool        `yaml:"required"`
//}

type MediaType struct {
	Schema *schema.Schema `yaml:"schema"`
}

type Content map[string]*MediaType

type RequestBody struct {
	Required bool    `yaml:"required"`
	Content  Content `yaml:"content"`
}

type Response struct {
	Content Content `yaml:"content"`
}
