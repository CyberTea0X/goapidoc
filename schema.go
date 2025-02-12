package goapidoc

var (
	SchemaInt     = Schema{Type: Integer}
	SchemaInt64   = Schema{Type: Integer, Format: "int64"}
	SchemaInt32   = Schema{Type: Integer, Format: "int32"}
	SchemaFloat   = Schema{Type: Number, Format: "float"}
	SchemaDouble  = Schema{Type: Number, Format: "double"}
	SchemaString  = Schema{Type: String}
	SchemaBoolean = Schema{Type: Boolean}
	SchemaNumber  = Schema{Type: Number}
)

type Schemas map[string]Schema

func (s Schemas) Merge(s2 Schemas) {
	for key := range s2 {
		value, exists := s2[key]
		if exists {
			s[key] = value
		}
	}
}

func (s Schemas) addSchema(value any) error {
	var err error
	schemaName := oapiSchemaName(value)
	_, exists := s[schemaName]
	if exists {
		return nil
	}
	s[schemaName], err = schemaFrom(value)
	if err != nil {
		return err
	}
	return nil
}

func (s Schemas) addSchemas(schemas ...any) []error {
	errs := make([]error, len(schemas))
	for i, value := range schemas {
		errs[i] = s.addSchema(value)
	}
	return errs
}

type Schema struct {
	Type       OapiType       `json:"type,omitempty"`
	Format     string         `json:"format,omitempty"`
	Items      *Schema        `json:"items,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
	Ref        string         `json:"$ref,omitempty"`
	Example    any            `json:"example,omitempty"`
}

func (s Schema) WithExample(example any) Schema {
	s.Example = example
	return s
}

func (s Schema) WithFormat(format string) Schema {
	s.Format = format
	return s
}
