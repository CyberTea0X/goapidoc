package goapidoc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func addr[T any](val T) *T { return &val }

// Converts types to schemas (can be referenced)
func SchemasOf(schemas ...any) Schemas {
	s := make(Schemas)
	for _, value := range schemas {
		err := s.addSchema(value)
		if err != nil {
			panic(err)
		}
	}
	return s
}

func ParametersFromStruct(value any, in string) []Parameter {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Struct {
		panic("expected struct")
	}
	rValue := reflect.ValueOf(value)
	parameters := make([]Parameter, t.NumField())
	for i := range parameters {
		field := t.Field(i)
		fValue := rValue.Field(i)
		if fValue.Type().Kind() == reflect.Pointer {
			fValue = fValue.Elem()
		}
		propertyName := strings.Split(field.Tag.Get("json"), ",")[0]
		if propertyName == "" {
			propertyName = strings.Split(field.Tag.Get("query"), ",")[0]
		}
		if propertyName == "" {
			propertyName = strings.Split(field.Tag.Get("form"), ",")[0]
		}
		if propertyName == "" {
			panic("failed to infer property name in struct " + t.Name())
		}
		fieldSchema, err := schemaFrom(fValue.Interface())
		if err != nil {
			panic(err)
		}
		parameters[i] = Parameter{
			In:          in,
			Name:        propertyName,
			Description: "",
			Required:    isRequired(field),
			Schema:      fieldSchema,
		}
	}
	return parameters
}

func toSnake(camel string) (snake string) {
	var b strings.Builder
	diff := 'a' - 'A'
	l := len(camel)
	for i, v := range camel {
		// A is 65, a is 97
		if v >= 'a' {
			b.WriteRune(v)
			continue
		}
		// v is capital letter here
		// irregard first letter
		// add underscore if last letter is capital letter
		// add underscore when previous letter is lowercase
		// add underscore when next letter is lowercase
		if (i != 0 || i == l-1) && (          // head and tail
		(i > 0 && rune(camel[i-1]) >= 'a') || // pre
			(i < l-1 && rune(camel[i+1]) >= 'a')) { //next
			b.WriteRune('_')
		}
		b.WriteRune(v + diff)
	}
	return b.String()
}

func SchemaFromStruct(value any) (Schema, error) {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Struct {
		return Schema{}, errors.New("expected struct")
	}

	properties := make(map[string]any)
	required := make([]string, 0)
	rValue := reflect.ValueOf(value)

	for i := range rValue.NumField() {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		propertyName := getPropertyName(field)
		fValue := getFieldValue(rValue.Field(i))
		if fValue.IsValid() == false { // Skip nil pointers
			continue
		}

		fieldSchema, err := schemaFrom(fValue.Interface())
		if err != nil {
			return Schema{}, err
		}

		if isRequired(field) {
			required = append(required, propertyName)
		}
		properties[propertyName] = fieldSchema
	}

	return Schema{
		Type:       "object",
		Required:   required,
		Properties: properties,
	}, nil
}

func getPropertyName(field reflect.StructField) string {
	propertyName := strings.Split(field.Tag.Get("json"), ",")[0]
	if propertyName == "" {
		propertyName = toSnake(field.Name)
	}
	return propertyName
}

func getFieldValue(fValue reflect.Value) reflect.Value {
	if fValue.Type().Kind() == reflect.Pointer {
		if fValue.IsNil() {
			return reflect.Value{} // Return invalid value for nil pointers
		}
		return fValue.Elem()
	}
	return fValue
}

func isRequired(field reflect.StructField) bool {
	return field.Tag.Get("binding") != "" || field.Tag.Get("validate") != ""
}

func SchemaFromSlice(value any) (Schema, error) {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Slice {
		return Schema{}, errors.New("expected slice")
	}
	items, err := schemaFrom(reflect.Zero(t.Elem()).Interface())
	if err != nil {
		return Schema{}, err
	}

	return Schema{
		Type:    Array,
		Example: value,
		Items:   &items,
	}, nil
}

func SchemaFromPrimitive(value any) (Schema, error) {
	t := reflect.TypeOf(value)
	return SchemaFromPrimitiveType(t, value)
}

// SchemaFromPrimitiveType creates a schema from a primitive type.
//
// If the example is not nil, it will be used as the example for the schema.
func SchemaFromPrimitiveType(t reflect.Type, example any) (Schema, error) {
	kind := t.Kind()
	if kind == reflect.Struct || kind == reflect.Slice || kind == reflect.Pointer {
		return Schema{}, fmt.Errorf("expected primitive type, got %v. %w", t, errors.New("expected primitive"))
	}
	var format string
	switch kind {
	case reflect.Int32:
		format = "int32"
	case reflect.Int64:
		format = "int64"
	}
	return Schema{
		Type:    toOapiType(t),
		Example: example,
		Format:  format,
	}, nil
}

func schemaFrom(value any) (Schema, error) {
	if schema, ok := value.(Schema); ok {
		return schema, nil
	}
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct:
		return SchemaFromStruct(value)
	case reflect.Slice:
		return SchemaFromSlice(value)
	case reflect.Pointer:
		return Schema{}, fmt.Errorf("making schema from pointer type %v is not supported. %w", t, errors.New("unexpected pointer"))
	default:
		return SchemaFromPrimitiveType(t, value)
	}
}

// Panics if can't convert the value to a schema.
func SchemaFrom(value any) Schema {
	schema, err := schemaFrom(value)
	if err != nil {
		panic(err)
	}
	return schema
}

// Returns error if can't convert the value to a schema.
func SchemaFromOrErr(value any) (Schema, error) {
	return schemaFrom(value)
}

func ArrayOf(schema Schema) Schema {
	return Schema{
		Type:  "array",
		Items: &schema,
	}
}

func Ref(value any) Schema {
	return Schema{
		Ref: "#/components/schemas/" + oapiSchemaName(value),
	}
}

func oapiSchemaName(value any) string {
	return reflect.ValueOf(value).Type().Name()
}

func ResponseWithForm(description string, schema Schema) Response {
	return Response{
		Description: description,
		Content: &Content{
			FormData: &ContentSchema{
				Schema: schema,
			},
		},
	}
}

func ResponseWithJson(description string, schema Schema) Response {
	return Response{
		Description: description,
		Content: &Content{
			Json: &ContentSchema{
				Schema: schema,
			},
		},
	}
}

func RequestWithJson(description string, schema Schema, required bool) *RequestBody {
	return &RequestBody{
		Description: description,
		Content: &Content{
			Json: &ContentSchema{
				Schema: schema,
			},
		},
		Required: required,
	}
}

func RequestWithForm(description string, schema Schema, required bool) *RequestBody {
	return &RequestBody{
		Description: description,
		Content: &Content{
			FormData: &ContentSchema{
				Schema: schema,
			},
		},
		Required: required,
	}
}

func toOapiType(t reflect.Type) OapiType {
	switch t.Kind() {
	case reflect.Int:
		return Integer
	case reflect.Int32:
		return Integer
	case reflect.Int64:
		return Integer
	case reflect.Uint:
		return Integer
	case reflect.Uint64:
		return Integer
	case reflect.Uint32:
		return Integer
	case reflect.Bool:
		return Boolean
	case reflect.Float64:
		return Number
	case reflect.Float32:
		return Number
	case reflect.String:
		return String
	case reflect.Slice:
		return Array
	case reflect.Struct:
		if t.Name() == "Time" {
			return String
		}
		return Object
	}
	panic("unhandled type: " + t.String())
}
