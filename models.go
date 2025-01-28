package main

import (
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

var BinaryFile = ContentSchema{
	Schema: Schema{
		Type:   String,
		Format: "binary",
	},
}

type Info struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Version     string            `json:"version,omitempty"`
	Contact     map[string]string `json:"contact,omitempty"`
	License     *License          `json:"license,omitempty"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Components struct {
	Schemas Schemas `json:"schemas,omitempty"`
}

type Document struct {
	OpenApiVersion string          `json:"openapi"`
	Info           Info            `json:"info"`
	Tags           []Tag           `json:"tags,omitempty"`
	Paths          map[string]Path `json:"paths"`
	Components     *Components     `json:"components,omitempty"`
}

type Tag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Path struct {
	Post   *Method `json:"post,omitempty"`
	Get    *Method `json:"get,omitempty"`
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
	In          string `json:"in"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Schema      Schema `json:"schema"`
}

type Method struct {
	Tags        []string         `json:"tags,omitempty"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	OperationId string           `json:"operationId"`
	Parameters  []Parameter      `json:"parameters,omitempty"`
	RequestBody *RequestBody     `json:"requestBody,omitempty"`
	Responses   map[int]Response `json:"responses"`
}

type Response struct {
	Description string   `json:"description"`
	Content     *Content `json:"content,omitempty"`
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
	Properties map[string]any `json:"properties,omitempty"`
	Ref        string         `json:"$ref,omitempty"`
	Example    any            `json:"example,omitempty"`
	Items      *Schema        `json:"items,omitempty"`
	Format     string         `json:"format,omitempty"`
}
