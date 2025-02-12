package goapidoc

import (
	"encoding/json"
	"os"
)

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

func NewRequestBody(description string, content *Content) *RequestBody {
	return &RequestBody{Description: description, Content: content}
}

type ContentSchema struct {
	Schema Schema `json:"schema"`
}

type Content struct {
	Json        *ContentSchema `json:"application/json,omitempty"`
	OctetStream *ContentSchema `json:"application/octet-stream,omitempty"`
	ImagePng    *ContentSchema `json:"image/png,omitempty"`
	ImageJpeg   *ContentSchema `json:"image/jpeg,omitempty"`
	ImageGif    *ContentSchema `json:"image/gif,omitempty"`
	TextPlain   *ContentSchema `json:"text/plain,omitempty"`
	TextHtml    *ContentSchema `json:"text/html,omitempty"`
	Xml         *ContentSchema `json:"application/xml,omitempty"`
	Pdf         *ContentSchema `json:"application/pdf,omitempty"`
	Zip         *ContentSchema `json:"application/zip,omitempty"`
	FormData    *ContentSchema `json:"multipart/form-data,omitempty"`
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

// Some binary file like image.png
var BinaryFile = ContentSchema{
	Schema: Schema{
		Type:   String,
		Format: "binary",
	},
}
