package goapidoc

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

func addr[T any](val T) *T { return &val }

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
		required := false
		if field.Tag.Get("binding") != "" {
			required = true
		}
		parameters[i] = Parameter{
			In:          in,
			Name:        propertyName,
			Description: "",
			Required:    required,
			Schema: Schema{
				Type:    toOapiType(fValue.Type()),
				Example: fValue.Interface(),
			},
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
	properties := make(map[string]any)
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Struct {
		return Schema{}, errors.New("expected struct")
	}
	rValue := reflect.ValueOf(value)
	numberOfFields := rValue.NumField()
	for i := 0; i < numberOfFields; i++ {
		field := t.Field(i)
		if field.IsExported() == false {
			continue
		}
		fValue := rValue.Field(i)
		propertyName := strings.Split(field.Tag.Get("json"), ",")[0]
		if propertyName == "" {
			propertyName = toSnake(field.Name)
		}
		if fValue.Type().Kind() == reflect.Pointer {
			if fValue.IsNil() {
				continue
			}
			fValue = fValue.Elem()
		}
		var fieldSchema Schema
		if fValue.Type().Kind() == reflect.Struct && fValue.Type().Name() != "Time" {
			var err error
			example, err := SchemaFrom(fValue.Interface())
			if err != nil {
				return Schema{}, err
			}
			fieldSchema = Schema{Type: "object", Properties: example.Properties}
		} else {
			var err error
			fieldSchema, err = SchemaFrom(fValue.Interface())
			if err != nil {
				log.Println("failed to generate schema from " + fValue.Type().String())
				continue
			}
		}
		properties[propertyName] = fieldSchema
		if fValue.Type().Kind() == reflect.Slice {
			property := properties[propertyName].(Schema)
			elemType := fValue.Type().Elem()
			property.Items = &Schema{
				Type: toOapiType(elemType),
			}
			properties[propertyName] = property
		}
	}
	return Schema{
		Type:       "object",
		Properties: properties,
	}, nil
}

func MustBuildSchemaFrom(value any) Schema {
	schema, err := SchemaFrom(value)
	if err != nil {
		panic(err)
	}
	return schema
}

func SchemaFromSlice(value any) (Schema, error) {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Slice {
		return Schema{}, errors.New("expected slice")
	}
	return Schema{
		Type:    Array,
		Example: value,
		Items: &Schema{
			Type: toOapiType(t.Elem()),
		},
	}, nil
}

func SchemaFromPrimitive(value any) (Schema, error) {
	t := reflect.TypeOf(value)
	kind := t.Kind()
	if kind == reflect.Struct || kind == reflect.Slice || kind == reflect.Pointer {
		return Schema{}, errors.New("expected primitive")
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
		Example: value,
		Format:  format,
	}, nil
}

func SchemaFrom(value any) (Schema, error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct:
		return SchemaFromStruct(value)
	case reflect.Slice:
		return SchemaFromSlice(value)
	default:
		return SchemaFromPrimitive(value)
	}
}

func ArrayOf(schema Schema) Schema {
	return Schema{
		Type:  "array",
		Items: &schema,
	}
}

func getSchemaRef(value any) Schema {
	return Schema{
		Ref: "#/components/schemas/" + oapiSchemaName(value),
	}
}

func ContentJsonSchemaRef(value any) *Content {
	return &Content{
		Json: &ContentSchema{
			Schema: getSchemaRef(value),
		},
	}
}

func oapiSchemaName(value any) string {
	return reflect.ValueOf(value).Type().Name()
}

// content is an empty struct witch works as a referense to a schema
func ResponseFromRef(description string, content any) Response {
	return Response{
		Description: description,
		Content:     ContentJsonSchemaRef(content),
	}
}
