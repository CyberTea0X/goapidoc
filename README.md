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

## Examples

### Write doc in golang

```golang
type Pet struct {
	Id   int64  `json:"id" validate:"required"`
  	Name string `json:"name" validate:"required"`
  	Tag  string `json:"tag"`
}

type Pets []Pet

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
						Schema:      oapi.MustBuildSchemaFrom(int32(1)),
					},
				},
				Responses: map[string]oapi.Response{
					"200": {
						Description: "A paged array of pets",
						Headers: map[string]oapi.Header{
							"x-next": {
								Description: "A link to the next page of responses",
								Schema: oapi.Schema{
									Type: oapi.String,
								},
							},
						},
						Content: oapi.ContentJsonSchemaRef(Pets{}),
					},
					"default": DefaultResponse,
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
			Pets{{Id: 1, Name: "Dog", Tag: "dogs"}, {Id: 2, Name: "Cat", Tag: "cats"}},
			Error{
				Code:    500,
				Message: "server crushed",
			},
		),
	},
}
doc.SaveAsJson("petstore.yaml")
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
            example: 1
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
                $ref: '#/components/schemas/Pets'
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
components:
  schemas:
    Error:
      type: object
      required:
        - code
        - message
      properties:
        code:
          type: integer
          format: int32
          example: 500
        message:
          type: string
          example: server crushed
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
    Pets:
      type: array
      items:
        type: object
        required:
          - id
          - name
        properties:
          id:
            type: integer
            format: int64
            example: 0
          name:
            type: string
            example: ""
          tag:
            type: string
            example: ""
      example:
        - id: 1
          name: Dog
          tag: dogs
        - id: 2
          name: Cat
          tag: cats
```

check out petstore example
