package goapidoc

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

type OapiType string

const (
	Integer OapiType = "integer"
	Boolean OapiType = "boolean"
	Number  OapiType = "number"
	String  OapiType = "string"
	Object  OapiType = "object"
	Array   OapiType = "array"
)

type Document struct {
	OpenApiVersion string          `json:"openapi" yaml:"openapi"`
	Servers        []Server        `json:"servers,omitempty" yaml:"servers,omitempty"`
	Info           Info            `json:"info" yaml:"info"`
	Security       []Security      `json:"security,omitempty" yaml:"security,omitempty"`
	Tags           []Tag           `json:"tags,omitempty" yaml:"tags,omitempty"`
	Paths          map[string]Path `json:"paths" yaml:"paths"`
	Components     *Components     `json:"components,omitempty" yaml:"components,omitempty"`
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

// SaveAsYaml сериализует документ в формате YAML и сохраняет в файл.
func (d *Document) SaveAsYaml(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2) // Устанавливаем отступ в 2 пробела

	err = encoder.Encode(d)
	if err != nil {
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

// Document Info
type Info struct {
	Title       string            `json:"title,omitempty" yaml:"title,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Version     string            `json:"version,omitempty" yaml:"version,omitempty"`
	Contact     map[string]string `json:"contact,omitempty" yaml:"contact,omitempty"`
	License     *License          `json:"license,omitempty" yaml:"license,omitempty"`
}

// Document license
type License struct {
	Name string `json:"name" yaml:"name"`
	Url  string `json:"url" yaml:"url"`
}

type Components struct {
	Schemas         Schemas         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	SecuritySchemes SecuritySchemes `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
}

type SecuritySchemes map[string]SecurityScheme
type SecurityScheme map[string]any

type Server struct {
	Url string `json:"url" yaml:"url"`
}

type Tag struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type Path struct {
	Get    *Method `json:"get,omitempty" yaml:"get,omitempty"`
	Post   *Method `json:"post,omitempty" yaml:"post,omitempty"`
	Put    *Method `json:"put,omitempty" yaml:"put,omitempty"`
	Patch  *Method `json:"patch,omitempty" yaml:"patch,omitempty"`
	Delete *Method `json:"delete,omitempty" yaml:"delete,omitempty"`
}

// example: if first path doesn't have Get and second has Get, but doesn't have Post, merged will have both
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
	Name        string `json:"name" yaml:"name"`
	In          string `json:"in" yaml:"in"`
	Description string `json:"description" yaml:"description"`
	Required    bool   `json:"required" yaml:"required"`
	Schema      Schema `json:"schema" yaml:"schema"`
}

type Method struct {
	Summary     string              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Security    []Security          `json:"security,omitempty" yaml:"security,omitempty"`
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	OperationId string              `json:"operationId" yaml:"operationId"`
	Tags        []string            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody *RequestBody        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses" yaml:"responses"`
}

type Security map[string]SecurityScopes

type SecurityScopes []string

type Header struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Schema      Schema `json:"schema" yaml:"schema"`
}

type Response struct {
	Description string            `json:"description" yaml:"description"`
	Headers     map[string]Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     *Content          `json:"content,omitempty" yaml:"content,omitempty"`
}

func (r Response) WithHeaders(headers map[string]Header) Response {
	r.Headers = headers
	return r
}

type RequestBody struct {
	Required    bool     `json:"required,omitempty" yaml:"required,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Content     *Content `json:"content,omitempty" yaml:"content,omitempty"`
}

type ContentSchema struct {
	Schema Schema `json:"schema" yaml:"schema"`
}

type Content struct {
	Json        *ContentSchema `json:"application/json,omitempty" yaml:"application/json,omitempty"`
	OctetStream *ContentSchema `json:"application/octet-stream,omitempty" yaml:"application/octet-stream,omitempty"`
	ImagePng    *ContentSchema `json:"image/png,omitempty" yaml:"image/png,omitempty"`
	ImageJpeg   *ContentSchema `json:"image/jpeg,omitempty" yaml:"image/jpeg,omitempty"`
	ImageGif    *ContentSchema `json:"image/gif,omitempty" yaml:"image/gif,omitempty"`
	TextPlain   *ContentSchema `json:"text/plain,omitempty" yaml:"text/plain,omitempty"`
	TextHtml    *ContentSchema `json:"text/html,omitempty" yaml:"text/html,omitempty"`
	Xml         *ContentSchema `json:"application/xml,omitempty" yaml:"application/xml,omitempty"`
	Pdf         *ContentSchema `json:"application/pdf,omitempty" yaml:"application/pdf,omitempty"`
	Zip         *ContentSchema `json:"application/zip,omitempty" yaml:"application/zip,omitempty"`
	FormData    *ContentSchema `json:"multipart/form-data,omitempty" yaml:"multipart/form-data,omitempty"`
}

// Some binary file like image.png
var BinaryFile = ContentSchema{
	Schema: Schema{
		Type:   String,
		Format: "binary",
	},
}

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

type Schema struct {
	Type       OapiType       `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string         `json:"format,omitempty" yaml:"format,omitempty"`
	Items      *Schema        `json:"items,omitempty" yaml:"items,omitempty"`
	Required   []string       `json:"required,omitempty" yaml:"required,omitempty"`
	Properties map[string]any `json:"properties,omitempty" yaml:"properties,omitempty"`
	Ref        string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Example    any            `json:"example,omitempty" yaml:"example,omitempty"`
}

func (s Schema) WithExample(example any) Schema {
	s.Example = example
	return s
}

func (s Schema) WithFormat(format string) Schema {
	s.Format = format
	return s
}
