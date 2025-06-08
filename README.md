# goapidoc

This is a simple and fast library for writing OpenAPI specification in Go.

## Import

```go
import "github.com/CyberTea0X/goapidoc"
```

or

```bash
go get github.com/CyberTea0X/goapidoc
```

## Main features

 - Typesafety
 - Struct to component conversion
 - Make refs to schemas
 - Write examples as golang structs
 - Supports openapi v 3.1.0
 - Write about 2 times less code than in json
 - Write spec in golang, convert to openapi.yaml or openapi.json

## Examples

### Petstore

```go
package main

import (
	oapi "github.com/CyberTea0X/goapidoc"
)

type Pet struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Tag  string `json:"tag"`
}

type Error struct {
	Code    int32  `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

var DefaultResponse = oapi.ResponseWithJson("unexpected error", oapi.SchemaFrom(Error{
	Code:    400,
	Message: "bad request",
}))

func main() {
	doc := oapi.Document{
		OpenApiVersion: "3.1.0",
		Info: oapi.Info{
			Version: "1.0.0",
			Title:   "Swagger Petstore",
			License: &oapi.License{
				Name: "MIT",
				Url:  "https://opensource.org/licenses/MIT",
			},
		},
		Servers: []oapi.Server{
			{Url: "http://petstore.swagger.io/v1"},
		},
		Paths: map[string]oapi.Path{
			"/pets": {
				Get: &oapi.Method{
					Summary:     "List all pets",
					OperationId: "listPets",
					Tags:        []string{"pets"},
					Parameters: []oapi.Parameter{
						{
							Name:        "limit",
							In:          "query",
							Description: "How many items to return at one time (max 100)",
							Required:    false,
							Schema:      oapi.SchemaInt32,
						},
					},
					Responses: map[string]oapi.Response{
						"200": oapi.ResponseWithJson("A paged array of pets", oapi.ArrayOf(oapi.Ref(Pet{}))).WithHeaders(
							map[string]oapi.Header{"x-next": {
								Description: "A link to the next page of responses",
								Schema: oapi.Schema{
									Type: oapi.String,
								},
							},
							},
						),
						"default": DefaultResponse,
					},
				},
				Post: &oapi.Method{
					Summary:     "Create a pet",
					OperationId: "createPets",
					Tags:        []string{"pets"},
					Responses: map[string]oapi.Response{
						"201":     oapi.ResponseWithoutContent("Null response"),
						"default": DefaultResponse,
					},
				},
			},
			"/pets/{petId}": {
				Get: &oapi.Method{
					Summary:     "Info for a specific pet",
					OperationId: "showPetById",
					Tags:        []string{"pets"},
					Parameters: []oapi.Parameter{
						{
							Name:        "petId",
							In:          "path",
							Required:    true,
							Description: "The id of the pet to retrieve",
							Schema:      oapi.Schema{Type: oapi.String},
						},
					},
					Responses: map[string]oapi.Response{
						"200":     oapi.ResponseWithJson("Expected response to a valid request", oapi.Ref(Pet{})),
						"default": DefaultResponse,
					},
				},
			},
		},
		Components: &oapi.Components{
			Schemas: oapi.SchemasOf(
				Pet{
					Id:   1,
					Name: "Dog",
					Tag:  "dogs",
				},
			),
		},
	}
	doc.SaveAsJson("petstore.json")
	doc.SaveAsYaml("petstore.yaml")
}
```

Spec:

```yaml
openapi: 3.1.0
servers:
  - url: http://petstore.swagger.io/v1
info:
  title: Swagger Petstore
  version: 1.0.0
  license:
    name: MIT
    url: https://opensource.org/licenses/MIT
paths:
  /pets:
    get:
      summary: List all pets
      operationId: listPets
      tags:
        - pets
      parameters:
        - name: limit
          in: query
          description: How many items to return at one time (max 100)
          required: false
          schema:
            type: integer
            format: int32
      responses:
        "200":
          description: A paged array of pets
          headers:
            x-next:
              description: A link to the next page of responses
              schema:
                type: string
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Pet'
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                type: object
                required:
                  - code
                  - message
                properties:
                  code:
                    type: integer
                    format: int32
                    example: 400
                  message:
                    type: string
                    example: bad request
    post:
      summary: Create a pet
      operationId: createPets
      tags:
        - pets
      responses:
        "201":
          description: Null response
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                type: object
                required:
                  - code
                  - message
                properties:
                  code:
                    type: integer
                    format: int32
                    example: 400
                  message:
                    type: string
                    example: bad request
  /pets/{petId}:
    get:
      summary: Info for a specific pet
      operationId: showPetById
      tags:
        - pets
      parameters:
        - name: petId
          in: path
          description: The id of the pet to retrieve
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Expected response to a valid request
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                type: object
                required:
                  - code
                  - message
                properties:
                  code:
                    type: integer
                    format: int32
                    example: 400
                  message:
                    type: string
                    example: bad request
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
          example: 1
        name:
          type: string
          example: Dog
        tag:
          type: string
          example: dogs
```

check out petstore example
