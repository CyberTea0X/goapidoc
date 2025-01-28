package goapidoc

import (
	"encoding/json"
	"os"
	"reflect"
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
	s[schemaName], err = SchemaFrom(value)
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

func NewResponse(description string, schema Schema) *Response {
	return &Response{
		Description: description,
		Content: &Content{
			Json: &ContentSchema{
				Schema: schema,
			},
		},
	}
}

type OapiType string

const (
	Integer OapiType = "integer"
	Boolean OapiType = "boolean"
	Number  OapiType = "number"
	String  OapiType = "string"
	Object  OapiType = "object"
	Array   OapiType = "array"
)

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

// Some binary file like image.png
var BinaryFile = ContentSchema{
	Schema: Schema{
		Type:   String,
		Format: "binary",
	},
}

// Document Info
type Info struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Version     string            `json:"version,omitempty"`
	Contact     map[string]string `json:"contact,omitempty"`
	License     *License          `json:"license,omitempty"`
}

// Document license
type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Components struct {
	Schemas Schemas `json:"schemas,omitempty"`
}

type Server struct {
	Url string `json:"url"`
}

type Document struct {
	OpenApiVersion string          `json:"openapi"`
	Info           Info            `json:"info"`
	Tags           []Tag           `json:"tags,omitempty"`
	Paths          map[string]Path `json:"paths"`
	Components     *Components     `json:"components,omitempty"`
	Servers        []Server        `json:"servers,omitempty"`
}

func (d *Document) SaveAsJson(filename string) error {
	bytes, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(bytes); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil

}

func (d *Document) AddPath(route string, newpath Path) {
	path, exists := d.Paths[route]
	if exists {
		newpath.Merge(path)
	}
	d.Paths[route] = newpath
}

type Tag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Path struct {
	Get    *Method `json:"get,omitempty"`
	Post   *Method `json:"post,omitempty"`
	Put    *Method `json:"put,omitempty"`
	Patch  *Method `json:"patch,omitempty"`
	Delete *Method `json:"delete,omitempty"`
}

func (p *Path) Merge(p2 Path) {
	if p.Post == nil {
		p.Post = p2.Post
	}
	if p.Get == nil {
		p.Get = p2.Get
	}
	if p.Put == nil {
		p.Put = p2.Put
	}
	if p.Patch == nil {
		p.Patch = p2.Patch
	}
	if p.Delete == nil {
		p.Delete = p2.Delete
	}
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Schema      Schema `json:"schema"`
}

type Method struct {
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	OperationId string              `json:"operationId"`
	Tags        []string            `json:"tags,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
	RequestBody *RequestBody        `json:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses"`
}

type Header struct {
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty"`
	Schema      Schema `json:"schema"`
}

type Response struct {
	Description string            `json:"description"`
	Headers     map[string]Header `json:"headers,omitempty"`
	Content     *Content          `json:"content,omitempty"`
}

type RequestBody struct {
	Description string   `json:"description,omitempty"`
	Content     *Content `json:"content,omitempty"`
}

type ContentSchema struct {
	Schema Schema `json:"schema"`
}

type Content struct {
	Json        *ContentSchema `json:"application/json,omitempty"`
	OctetStream *ContentSchema `json:"application/octet-stream,omitempty"`
	ImagePng    *ContentSchema `json:"image/png,omitempty"`
}

type Schema struct {
	Type       OapiType       `json:"type,omitempty"`
	Format     string         `json:"format,omitempty"`
	Items      *Schema        `json:"items,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
	Ref        string         `json:"$ref,omitempty"`
	Example    any            `json:"example,omitempty"`
}
