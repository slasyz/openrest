package schema

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Schema struct {
	Type       string     `yaml:"type"`
	Properties Properties `yaml:"properties"`
	Items      *Schema    `yaml:"items"`
}

// Properties is ordered map[string]*Schema
type Properties struct {
	List []Property
}

func (p *Properties) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := make(map[string]*Schema)
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

		p.List = append(p.List, Property{
			Key:   key,
			Value: val,
		})
	}

	return nil
}

type Property struct {
	Key   string
	Value *Schema
}
